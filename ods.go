package ods

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Document represents an ODS (OpenDocument Spreadsheet) file.
type Document struct {
	path    string
	entries []zipEntry
	files   map[string][]byte

	content  Content
	meta     Meta
	manifest Manifest

	contentDirty  bool
	metaDirty     bool
	manifestDirty bool
	err           error

	styles *styleRegistry
}

// zipEntry preserves the original metadata of a zip archive entry to
// preserve zip shape.
type zipEntry struct {
	Name     string
	Method   uint16
	Modified time.Time
	Flags    uint16
}

// Open opens an existing ODS file from disk and parses its XML content.
func Open(path string) (*Document, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open ODS %q: %w", path, err)
	}

	doc, err := readZip(&zr.Reader)
	if err != nil {
		_ = zr.Close()
		return nil, fmt.Errorf("failed to parse archive: %w", err)
	}

	if err := zr.Close(); err != nil {
		return nil, fmt.Errorf("failed to close ODS %q: %w", path, err)
	}

	doc.path = path
	return doc, nil
}

// OpenReader opens an ODS document from memory.
func OpenReader(r io.ReaderAt, size int64) (*Document, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("failed to read ODS: %w", err)
	}

	doc, err := readZip(zr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse archive: %w", err)
	}

	return doc, nil
}

// Create creates an empty ODS document with a default sheet and metadata.
// Optionally provide a path to be used by Save.
func Create(path string) (*Document, error) {
	doc := &Document{
		path: path,
		files: map[string][]byte{
			zipPathMimetype: []byte(odsMimetype),
			zipPathSettings: minimalSettingsXML,
			zipPathStyles:   minimalStylesXML,
			zipPathManifest: nil,
		},
		content:       newContent(),
		meta:          newMeta(),
		manifest:      newManifest(),
		contentDirty:  true,
		metaDirty:     true,
		manifestDirty: true,
		styles:        newStyleRegistry(),
	}
	doc.styles.attach(doc)
	doc.entries = []zipEntry{
		{Name: zipPathMimetype, Method: zip.Store},
		{Name: zipPathContent, Method: zip.Deflate},
		{Name: zipPathMeta, Method: zip.Deflate},
		{Name: zipPathSettings, Method: zip.Deflate},
		{Name: zipPathStyles, Method: zip.Deflate},
		{Name: zipPathManifest, Method: zip.Deflate},
	}
	return doc, nil
}

// Err returns the sticky error value associated with the document.
func (d *Document) Err() error {
	if d == nil {
		return fmt.Errorf("document is nil")
	}
	return d.err
}

// Save wraps SaveAs with the document's path.
func (d *Document) Save() error {
	if d == nil {
		return fmt.Errorf("document is nil")
	}

	return d.SaveAs(d.path)
}

// SaveAs writes the document to the specified path safely via an atomic write.
// Parent directories are created if needed.
func (d *Document) SaveAs(path string) error {
	if d == nil {
		return fmt.Errorf("document is nil")
	} else if path == "" {
		return fmt.Errorf("path not provided")
	} else if d.err != nil {
		return d.err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create directories %q: %w", dir, err)
	}

	tmp, err := os.CreateTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmp.Name())

	if err := d.writeTo(tmp); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("failed to write to temporary file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("failed to replace file: %w", err)
	}

	d.path = path
	return nil
}

// SetTitle sets the document title metadata.
func (d *Document) SetTitle(title string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Title = title
	d.metaDirty = true
	return d
}

// SetAuthor sets the document creator and initial-creator metadata.
// initial-creator is only set if it has no value.
func (d *Document) SetAuthor(author string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Creator = author
	d.meta.Body.InitialCreator = coalesce(d.meta.Body.InitialCreator, author)
	d.metaDirty = true
	return d
}

// SetSubject sets the document subject metadata.
func (d *Document) SetSubject(subject string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Subject = subject
	d.metaDirty = true
	return d
}

// SetDescription sets the document description metadata.
func (d *Document) SetDescription(description string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Description = description
	d.metaDirty = true
	return d
}

// SetKeywords sets the document keywords metadata.
func (d *Document) SetKeywords(keywords string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Keywords = keywords
	d.metaDirty = true
	return d
}

// SetLanguage sets the document language metadata.
func (d *Document) SetLanguage(language string) *Document {
	if d == nil {
		return nil
	}
	d.meta.Body.Language = language
	d.metaDirty = true
	return d
}

// Sheet returns the sheet at the given zero-based index. If the index exceeds
// the current number of sheets, additional sheets are created automatically
// with default names.
func (d *Document) Sheet(index int) *Sheet {
	if d == nil {
		return &Sheet{index: -1}
	}
	if index < 0 {
		d.setErr(fmt.Errorf("sheet index must be non-negative: %d", index))
		return &Sheet{doc: d, index: -1}
	}
	for len(d.content.Body.Spreadsheet.Table) <= index {
		d.content.Body.Spreadsheet.Table = append(d.content.Body.Spreadsheet.Table, d.newSheetTable(fmt.Sprintf(defaultSheetNameFormat, len(d.content.Body.Spreadsheet.Table)+1)))
		d.contentDirty = true
	}
	return &Sheet{doc: d, index: index}
}

// SheetByName returns the sheet with the specified name. If no sheet with the
// name exists, a document error is set and an invalid sheet is returned.
func (d *Document) SheetByName(name string) *Sheet {
	if d == nil {
		return &Sheet{index: -1}
	}
	for i := range d.content.Body.Spreadsheet.Table {
		if d.content.Body.Spreadsheet.Table[i].Name == name {
			return &Sheet{doc: d, index: i}
		}
	}
	d.setErr(fmt.Errorf("sheet %q not found", name))
	return &Sheet{doc: d, index: -1}
}

// AddSheet appends a new sheet with the given name to the document.
// If name is empty, a default name is generated. Returns the new sheet.
func (d *Document) AddSheet(name string) *Sheet {
	if d == nil {
		return &Sheet{index: -1}
	}
	name = coalesce(name, fmt.Sprintf(defaultSheetNameFormat, len(d.content.Body.Spreadsheet.Table)+1))
	d.content.Body.Spreadsheet.Table = append(d.content.Body.Spreadsheet.Table, d.newSheetTable(name))
	d.contentDirty = true
	return &Sheet{doc: d, index: len(d.content.Body.Spreadsheet.Table) - 1}
}

// RemoveSheet removes the sheet at the given zero-based index.
func (d *Document) RemoveSheet(index int) *Document {
	if d == nil {
		return nil
	}
	if index < 0 || index >= len(d.content.Body.Spreadsheet.Table) {
		d.setErr(fmt.Errorf("sheet index out of range: %d", index))
		return d
	}
	d.content.Body.Spreadsheet.Table = append(d.content.Body.Spreadsheet.Table[:index], d.content.Body.Spreadsheet.Table[index+1:]...)
	d.contentDirty = true
	return d
}

// setErr records the first error on the document. It is nil-safe and only
// sets the error once; subsequent errors are silently discarded.
func (d *Document) setErr(err error) {
	if d == nil {
		return
	}
	if d.err == nil {
		d.err = err
	}
}

// readZip reads all files of a zip archive and parses the ODF XML content,
// meta, and manifest parts. Validates that the mimetype matches the ODS format.
func readZip(zr *zip.Reader) (*Document, error) {
	doc := &Document{
		files:  make(map[string][]byte, len(zr.File)),
		styles: newStyleRegistry(),
	}
	doc.styles.attach(doc)

	for _, f := range zr.File {
		doc.entries = append(doc.entries, zipEntry{
			Name:     f.Name,
			Method:   f.Method,
			Modified: f.Modified,
			Flags:    f.Flags,
		})
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("%s: failed to open file: %w", f.Name, err)
		}
		data, readErr := io.ReadAll(rc)
		closeErr := rc.Close()
		if readErr != nil {
			return nil, fmt.Errorf("%s: failed to read file: %w", f.Name, readErr)
		}
		if closeErr != nil {
			return nil, fmt.Errorf("%s: failed to close file: %w", f.Name, closeErr)
		}
		doc.files[f.Name] = data

		switch f.Name {
		case zipPathContent:
			if err := xml.Unmarshal(data, &doc.content); err != nil {
				return nil, fmt.Errorf("%s: failed to parse XML: %w", f.Name, err)
			}
			fillContentNamespaces(&doc.content)

		case zipPathMeta:
			if err := xml.Unmarshal(data, &doc.meta); err != nil {
				return nil, fmt.Errorf("%s: failed to parse XML: %w", f.Name, err)
			}
			fillMetaNamespaces(&doc.meta)

		case zipPathManifest:
			if err := xml.Unmarshal(data, &doc.manifest); err == nil {
				doc.manifest.Xmlns = coalesce(doc.manifest.Xmlns, nsManifest)
			}
		}
	}

	if mimetype := string(doc.files[zipPathMimetype]); mimetype != odsMimetype {
		return nil, fmt.Errorf("invalid ODS mimetype %q", mimetype)
	}

	if len(doc.content.Body.Spreadsheet.Table) == 0 {
		doc.content = newContent()
		doc.contentDirty = true
	}

	if doc.meta.Office == "" && doc.meta.Body.Generator == "" {
		doc.meta = newMeta()
		doc.metaDirty = true
	}

	if len(doc.manifest.Files) == 0 {
		doc.manifest = newManifest()
		doc.manifestDirty = true
	}

	return doc, nil
}

// fillContentNamespaces ensures all ODF namespace attributes are populated on
// the Content root element.
func fillContentNamespaces(c *Content) {
	c.Office = coalesce(c.Office, nsOffice)
	c.Table = coalesce(c.Table, nsTable)
	c.Text = coalesce(c.Text, nsText)
	c.Style = coalesce(c.Style, nsStyle)
	c.FO = coalesce(c.FO, nsFO)
	c.SVG = coalesce(c.SVG, nsSVG)
	c.Number = coalesce(c.Number, nsNumber)
	c.MetaNS = coalesce(c.MetaNS, nsMeta)
	c.DC = coalesce(c.DC, nsDC)
	c.Of = coalesce(c.Of, nsOf)
	c.Calcext = coalesce(c.Calcext, nsCalcext)
	c.Version = coalesce(c.Version, odfVersion)
}

// fillMetaNamespaces ensures all ODF namespace attributes are populated on
// the Meta root element.
func fillMetaNamespaces(m *Meta) {
	m.Office = coalesce(m.Office, nsOffice)
	m.MetaNS = coalesce(m.MetaNS, nsMeta)
	m.DC = coalesce(m.DC, nsDC)
	m.Version = coalesce(m.Version, odfVersion)
}

// newTable creates a default Table with the standard ODS dimensions
// compressed into repeated entries.
func newTable(name string) Table {
	return Table{
		Name:        name,
		TableColumn: []TableColumn{{NumberColumnsRepeated: defaultColumnRepeat}},
		TableRow:    []TableRow{{NumberRowsRepeated: defaultRowRepeat, TableCell: []TableCell{{NumberColumnsRepeated: defaultColumnRepeat}}}},
	}
}

// newSheetTable creates a new table that inherits structural styles from the
// document's first sheet, or falls back to newTable defaults if no sheets exist.
func (d *Document) newSheetTable(name string) Table {
	if len(d.content.Body.Spreadsheet.Table) == 0 {
		return newTable(name)
	}

	template := d.content.Body.Spreadsheet.Table[0]
	table := Table{
		Name:      name,
		StyleName: template.StyleName,
	}

	if len(template.TableColumn) > 0 {
		col := template.TableColumn[0]
		col.NumberColumnsRepeated = ""
		table.TableColumn = []TableColumn{col}
	}

	if len(template.TableRow) > 0 {
		row := template.TableRow[0]
		row.NumberRowsRepeated = ""

		if len(row.TableCell) > 0 {
			cell := row.TableCell[0]
			cell.NumberColumnsRepeated = ""
			row.TableCell = []TableCell{cell}
		} else {
			row.TableCell = []TableCell{{}}
		}

		table.TableRow = []TableRow{row}
	}

	if len(table.TableColumn) == 0 {
		table.TableColumn = []TableColumn{{}}
	}

	if len(table.TableRow) == 0 {
		table.TableRow = []TableRow{{TableCell: []TableCell{{}}}}
	}

	return table
}
