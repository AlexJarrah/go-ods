package ods

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"io"
	"maps"
	"strconv"
	"strings"
	"time"
)

// writeTo serializes the document's content, metadata, and manifest into the
// given writer as a valid ODS zip archive. It handles dirty-flag checks,
// compression of tables, and preservation of original zip entry metadata.
func (d *Document) writeTo(w io.Writer) error {
	var err error
	content := d.files[zipPathContent]
	if d.contentDirty || len(content) == 0 {
		for i := range d.content.Body.Spreadsheet.Table {
			compressTable(&d.content.Body.Spreadsheet.Table[i])
		}
		content, err = marshalXML(d.content)
		if err != nil {
			return fmt.Errorf("failed to marshal %s: %w", zipPathContent, err)
		}
	}

	meta := d.files[zipPathMeta]
	if d.contentDirty {
		d.updateDocumentStatistics()
		d.metaDirty = true
	}
	if d.metaDirty || len(meta) == 0 {
		meta, err = marshalXML(d.meta)
		if err != nil {
			return fmt.Errorf("failed to marshal %s: %w", zipPathMeta, err)
		}
	}

	manifest := d.files[zipPathManifest]
	if d.manifestDirty || len(manifest) == 0 {
		d.manifest = manifestFor(d)
		manifest, err = marshalXML(d.manifest)
		if err != nil {
			return fmt.Errorf("failed to marshal %s: %w", zipPathManifest, err)
		}
	}

	files := map[string][]byte{}
	maps.Copy(files, d.files)
	files[zipPathMimetype] = []byte(odsMimetype)
	files[zipPathContent] = content
	files[zipPathMeta] = meta
	if settings, ok := files[zipPathSettings]; ok {
		files[zipPathSettings] = syncSettingsSheets(settings, d.sheetNames())
	}
	files[zipPathManifest] = manifest
	if _, ok := files[zipPathSettings]; !ok {
		files[zipPathSettings] = minimalSettingsXML
	}
	if _, ok := files[zipPathStyles]; !ok {
		files[zipPathStyles] = minimalStylesXML
	}

	zw := zip.NewWriter(w)
	if err := writeMimetypeEntry(zw, d.entryFor(zipPathMimetype)); err != nil {
		return fmt.Errorf("failed to write mimetype entry: %w", err)
	}

	written := map[string]bool{zipPathMimetype: true}
	for _, entry := range d.entries {
		if written[entry.Name] {
			continue
		}
		data, ok := files[entry.Name]
		if !ok {
			continue
		}
		if err := writeZipEntry(zw, entry, data); err != nil {
			return fmt.Errorf("failed to write zip entry: %w", err)
		}
		written[entry.Name] = true
	}

	for _, name := range requiredZipPaths {
		if written[name] {
			continue
		}
		if err := writeZipEntry(zw, zipEntry{Name: name, Method: zip.Deflate}, files[name]); err != nil {
			return fmt.Errorf("failed to write zip entry: %w", err)
		}
		written[name] = true
	}

	if err := zw.Close(); err != nil {
		return fmt.Errorf("failed to close ODS archive: %w", err)
	}

	return nil
}

// updateDocumentStatistics recalculates the table and cell counts in the
// document's metadata.
func (d *Document) updateDocumentStatistics() {
	stats := &d.meta.Body.DocumentStatistic
	stats.TableCount = strconv.Itoa(len(d.content.Body.Spreadsheet.Table))
	stats.CellCount = strconv.Itoa(countUsedCells(d.content))
	if stats.ObjectCount == "" {
		stats.ObjectCount = "0"
	}
}

// countUsedCells returns the total number of non-empty cells across all
// tables in the content, accounting for column repetition.
func countUsedCells(content Content) int {
	count := 0
	for _, table := range content.Body.Spreadsheet.Table {
		for _, row := range table.TableRow {
			for _, cell := range row.TableCell {
				if !cellIsEmpty(cell) {
					count += repeatCount(cell.NumberColumnsRepeated)
				}
			}
		}
	}
	return count
}

// sheetNames returns all sheet names in the document.
func (d *Document) sheetNames() []string {
	names := make([]string, 0, len(d.content.Body.Spreadsheet.Table))
	for _, table := range d.content.Body.Spreadsheet.Table {
		names = append(names, table.Name)
	}
	return names
}

// entryFor returns the zipEntry metadata for the given zip path, or a
// default entry with zip.Store method if not found.
func (d *Document) entryFor(name string) zipEntry {
	for _, entry := range d.entries {
		if entry.Name == name {
			return entry
		}
	}
	return zipEntry{Name: name, Method: zip.Store}
}

// syncSettingsSheets updates the settings.xml byte slice to include entries
// for all sheet names in the document. It clones the first table entry as a
// template for any new sheets, escaping XML special characters in names.
func syncSettingsSheets(settings []byte, sheetNames []string) []byte {
	raw := string(settings)
	const tablesStart = `<config:config-item-map-named config:name="Tables">`
	const tablesEnd = `</config:config-item-map-named>`
	start := strings.Index(raw, tablesStart)
	if start < 0 {
		return settings
	}

	bodyStart := start + len(tablesStart)
	end := strings.Index(raw[bodyStart:], tablesEnd)
	if end < 0 {
		return settings
	}
	bodyEnd := bodyStart + end
	body := raw[bodyStart:bodyEnd]

	entryStart := strings.Index(body, `<config:config-item-map-entry config:name="`)
	if entryStart < 0 {
		return settings
	}
	entryEndRel := strings.Index(body[entryStart:], `</config:config-item-map-entry>`)
	if entryEndRel < 0 {
		return settings
	}
	entryEnd := entryStart + entryEndRel + len(`</config:config-item-map-entry>`)
	template := body[entryStart:entryEnd]

	updated := body
	for _, name := range sheetNames {
		escapedName := xmlAttrEscape(name)
		if name == "" || strings.Contains(updated, `config:name="`+escapedName+`"`) {
			continue
		}
		updated += strings.Replace(template, templateSheetName(template), escapedName, 1)
	}
	if updated == body {
		return settings
	}

	return []byte(raw[:bodyStart] + updated + raw[bodyEnd:])
}

// templateSheetName extracts the sheet name from a config-item-map-entry
// template string. Returns an empty string if the template is invalid.
func templateSheetName(template string) string {
	const prefix = `<config:config-item-map-entry config:name="`
	start := strings.Index(template, prefix)
	if start < 0 {
		return ""
	}
	start += len(prefix)
	end := strings.Index(template[start:], `"`)
	if end < 0 {
		return ""
	}
	return template[start : start+end]
}

// xmlAttrEscape escapes XML special characters in a string for use in
// attribute values.
func xmlAttrEscape(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
		"'", "&apos;",
	)
	return replacer.Replace(s)
}

// writeMimetypeEntry writes the mimetype zip entry using raw (uncompressed)
// storage with the exact byte layout required by the ODS specification.
func writeMimetypeEntry(zw *zip.Writer, entry zipEntry) error {
	data := []byte(odsMimetype)
	flags := entry.Flags &^ 0x08
	if flags == 0 {
		flags = 0x0800
	}

	modTime, modDate := zipDOSTime(entry.Modified)
	header := &zip.FileHeader{
		Name:               zipPathMimetype,
		CreatorVersion:     20,
		ReaderVersion:      20,
		Flags:              flags,
		Method:             zip.Store,
		ModifiedTime:       modTime,
		ModifiedDate:       modDate,
		CRC32:              crc32.ChecksumIEEE(data),
		CompressedSize:     uint32(len(data)),
		UncompressedSize:   uint32(len(data)),
		CompressedSize64:   uint64(len(data)),
		UncompressedSize64: uint64(len(data)),
	}

	fw, err := zw.CreateRaw(header)
	if err != nil {
		return fmt.Errorf("failed to create zip member %s: %w", zipPathMimetype, err)
	}

	if _, err := fw.Write(data); err != nil {
		return fmt.Errorf("failed to write zip member %s: %w", zipPathMimetype, err)
	}

	return nil
}

// marshalXML encodes an ODF value (Content, Meta, or Manifest) into XML bytes
// with proper namespace prefixes and indented formatting.
func marshalXML(v any) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	switch value := v.(type) {
	case Content:
		fillContentNamespaces(&value)
		v = value

	case Meta:
		fillMetaNamespaces(&value)
		v = value

	case Manifest:
		if value.Xmlns == "" {
			value.Xmlns = nsManifest
		}
		v = value
	}

	if err := enc.Encode(v); err != nil {
		return nil, fmt.Errorf("failed to encode XML: %w", err)
	}

	if err := enc.Flush(); err != nil {
		return nil, fmt.Errorf("failed to flush buffer to writer: %w", err)
	}

	return buf.Bytes(), nil
}

// writeZipEntry writes a single zip entry to the writer. For Store-method
// entries, it creates a raw header to preserve the exact disk layout.
// For Deflate entries, it uses the standard zip.Writer.CreateHeader.
func writeZipEntry(zw *zip.Writer, entry zipEntry, data []byte) error {
	modTime, modDate := zipDOSTime(entry.Modified)
	if entry.Method == zip.Store {
		flags := entry.Flags &^ 0x08
		header := &zip.FileHeader{
			Name:               entry.Name,
			CreatorVersion:     20,
			ReaderVersion:      20,
			Flags:              flags,
			Method:             zip.Store,
			ModifiedTime:       modTime,
			ModifiedDate:       modDate,
			CRC32:              crc32.ChecksumIEEE(data),
			CompressedSize:     uint32(len(data)),
			UncompressedSize:   uint32(len(data)),
			CompressedSize64:   uint64(len(data)),
			UncompressedSize64: uint64(len(data)),
		}

		fw, err := zw.CreateRaw(header)
		if err != nil {
			return fmt.Errorf("failed to create zip member %s: %w", entry.Name, err)
		}

		if _, err := fw.Write(data); err != nil {
			return fmt.Errorf("failed to write zip member %s: %w", entry.Name, err)
		}

		return nil
	}

	header := &zip.FileHeader{
		Name:         entry.Name,
		Flags:        entry.Flags,
		Method:       entry.Method,
		ModifiedTime: modTime,
		ModifiedDate: modDate,
	}

	fw, err := zw.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("failed to create zip member %s: %w", entry.Name, err)
	}

	if _, err := fw.Write(data); err != nil {
		return fmt.Errorf("failed to write zip member %s: %w", entry.Name, err)
	}

	return nil
}

// manifestFor builds a Manifest listing all zip entries in the document,
// preserving original entries that were present in the source archive.
func manifestFor(d *Document) Manifest {
	seen := map[string]bool{zipPathRoot: true}
	names := []string{zipPathRoot}
	for _, name := range manifestPaths {
		seen[name] = true
		names = append(names, name)
	}

	for _, entry := range d.entries {
		if entry.Name == zipPathMimetype || entry.Name == zipPathManifest || seen[entry.Name] {
			continue
		}
		names = append(names, entry.Name)
		seen[entry.Name] = true
	}

	manifest := Manifest{Xmlns: nsManifest, Version: odfVersion}
	wrote := map[string]bool{}
	for _, name := range names {
		if name == "" || wrote[name] {
			continue
		}
		wrote[name] = true
		manifest.Files = append(manifest.Files, ManifestFile{
			FullPath:  name,
			MediaType: mediaTypeFor(name),
		})
	}

	return manifest
}

// mediaTypeFor returns the MIME media type for a standard ODS zip entry.
func mediaTypeFor(name string) string {
	switch name {
	case zipPathRoot:
		return odsMimetype

	case zipPathContent, zipPathMeta, zipPathSettings, zipPathStyles:
		return xmlMediaType

	default:
		return ""
	}
}

// zipDOSTime converts a timestamp to MS-DOS date/time fields for zip local
// headers. Modified is stored as time.Time when reading, writing DOS fields
// directly (without FileHeader.Modified) avoids extended-timestamp extra data
// that would change on-disk zip shape.
func zipDOSTime(t time.Time) (uint16, uint16) {
	if t.IsZero() {
		return 0, 0
	}
	t = t.UTC()
	year := t.Year()
	if year < 1980 {
		return 0, 0
	}
	return uint16(t.Second()/2 + t.Minute()<<5 + t.Hour()<<11), uint16(t.Day() + int(t.Month())<<5 + (year-1980)<<9)
}
