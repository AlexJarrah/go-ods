package ods

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCreateSaveOpenRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "created.ods")

	doc, err := Create(path)
	if err != nil {
		t.Fatal(err)
	}
	doc.SetTitle("Q3 Report").SetAuthor("Alex")
	doc.Sheet(0).SetName("Sales").Cell(0, 0).Set("Hello").
		Style().
		Bold(true).
		Color("#ff0000").
		BackgroundColor("#ffffcc").
		HAlign("center").
		Apply()
	doc.Sheet(0).Cell(1, 0).Set(42)
	doc.Sheet(0).Cell(2, 0).Set(true)
	now := time.Date(2026, 5, 27, 1, 2, 3, 0, time.UTC)
	doc.Sheet(0).Cell(3, 0).Set(now)
	doc.Sheet(0).Cell(4, 0).SetFormula("SUM(A1:A4)")

	if err := doc.Save(); err != nil {
		t.Fatal(err)
	}
	assertSpecZip(t, path)

	opened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := opened.Sheet(0).Name(); got != "Sales" {
		t.Fatalf("sheet name = %q", got)
	}
	if got, err := opened.Sheet(0).Cell(0, 0).String(); err != nil || got != "Hello" {
		t.Fatalf("string cell = %q, %v", got, err)
	}
	if got, err := opened.Sheet(0).Cell(1, 0).Float64(); err != nil || got != 42 {
		t.Fatalf("float cell = %v, %v", got, err)
	}
	if got, err := opened.Sheet(0).Cell(2, 0).Bool(); err != nil || !got {
		t.Fatalf("bool cell = %v, %v", got, err)
	}
	if got, err := opened.Sheet(0).Cell(3, 0).Time(); err != nil || !got.Equal(now) {
		t.Fatalf("time cell = %v, %v", got, err)
	}

	content := zipMember(t, path, zipPathContent)
	if !strings.Contains(content, `table:formula="of:=SUM([.A1:.A4])"`) {
		t.Fatalf("formula was not normalized to ODF A1 syntax:\n%s", content)
	}
	for _, want := range []string{
		`office:value-type="float"`,
		`table:formula="of:=SUM([.A1:.A4])"`,
		`office:value="0"`,
		`calcext:value-type="float"`,
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("formula cell is missing %q:\n%s", want, content)
		}
	}
	if got, err := opened.Sheet(0).Cell(4, 0).Formula(); err != nil || got != "of:=SUM([.A1:.A4])" {
		t.Fatalf("formula cell = %q, %v", got, err)
	}
	manifest := zipMember(t, path, zipPathManifest)
	if strings.Contains(manifest, `manifest:full-path="`+zipPathManifest+`"`) {
		t.Fatalf("manifest.xml should not list itself:\n%s", manifest)
	}
	if strings.Contains(content, `table:number-rows-repeated="104857`) || strings.Contains(content, `table:number-columns-repeated="102`) {
		t.Fatalf("created content.xml contains unnecessary max-sheet empty tails:\n%s", content)
	}
}

func TestOpenBlankSaveAsPreservesPassthroughAndZipShape(t *testing.T) {
	src := filepath.Join("tmp", "Blank.ods")
	if _, err := os.Stat(src); err != nil {
		t.Skip("sample Blank.ods not available")
	}

	doc, err := Open(src)
	if err != nil {
		t.Fatal(err)
	}
	doc.Sheet(0).Cell(100, 20).Set("sparse")
	dst := filepath.Join(t.TempDir(), "copy.ods")
	if err := doc.SaveAs(dst); err != nil {
		t.Fatal(err)
	}
	assertSpecZip(t, dst)

	opened, err := Open(dst)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := opened.Sheet(0).Cell(100, 20).String(); err != nil || got != "sparse" {
		t.Fatalf("sparse cell = %q, %v", got, err)
	}
	if opened.Sheet(0).RowCount() != 101 {
		t.Fatalf("row count = %d", opened.Sheet(0).RowCount())
	}
	if opened.Sheet(0).ColCount() != 21 {
		t.Fatalf("col count = %d", opened.Sheet(0).ColCount())
	}

	originalManifest := zipMember(t, src, zipPathManifest)
	savedManifest := zipMember(t, dst, zipPathManifest)
	if savedManifest != originalManifest {
		t.Fatalf("opened document manifest should be preserved\noriginal:\n%s\nsaved:\n%s", originalManifest, savedManifest)
	}
	if !strings.Contains(savedManifest, `manifest:full-path="manifest.rdf" manifest:media-type="application/rdf+xml"`) {
		t.Fatalf("manifest.rdf media type was not preserved:\n%s", savedManifest)
	}
	if !strings.Contains(savedManifest, `manifest:full-path="Thumbnails/thumbnail.png" manifest:media-type="image/png"`) {
		t.Fatalf("thumbnail media type was not preserved:\n%s", savedManifest)
	}

	content := zipMember(t, dst, zipPathContent)
	if !strings.Contains(content, `style:name="ta1" style:family="table" style:master-page-name="Default"`) {
		t.Fatalf("table style master page was not preserved:\n%s", content)
	}
	if !strings.Contains(content, `<style:table-properties table:display="true" style:writing-mode="lr-tb"`) {
		t.Fatalf("table style properties were not preserved:\n%s", content)
	}
	for _, want := range []string{
		`<office:font-face-decls>`,
		`<table:calculation-settings table:automatic-find-labels="false" table:use-regular-expressions="false" table:use-wildcards="true" table:null-year="1950"`,
		`<table:named-expressions>`,
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("content.xml did not preserve %q:\n%s", want, content)
		}
	}
}

func TestNoopSavePreservesXMLParts(t *testing.T) {
	src := filepath.Join("tmp", "old.ods")
	if _, err := os.Stat(src); err != nil {
		src = filepath.Join("tmp", "Blank.ods")
	}
	if _, err := os.Stat(src); err != nil {
		t.Skip("sample ODS not available")
	}

	doc, err := Open(src)
	if err != nil {
		t.Fatal(err)
	}
	dst := filepath.Join(t.TempDir(), "noop.ods")
	if err := doc.SaveAs(dst); err != nil {
		t.Fatal(err)
	}

	for _, name := range []string{zipPathContent, zipPathMeta, zipPathSettings, zipPathManifest} {
		if got, want := zipMember(t, dst, name), zipMember(t, src, name); got != want {
			t.Fatalf("%s changed on no-op save", name)
		}
	}
}

func TestReadOnlyCellAccessPreservesContentXML(t *testing.T) {
	src := filepath.Join("tmp", "old.ods")
	if _, err := os.Stat(src); err != nil {
		src = filepath.Join("tmp", "Blank.ods")
	}
	if _, err := os.Stat(src); err != nil {
		t.Skip("sample ODS not available")
	}

	doc, err := Open(src)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := doc.Sheet(0).Cell(0, 0).String(); err != nil {
		t.Fatal(err)
	}
	dst := filepath.Join(t.TempDir(), "read-only.ods")
	if err := doc.SaveAs(dst); err != nil {
		t.Fatal(err)
	}
	if got, want := zipMember(t, dst, zipPathContent), zipMember(t, src, zipPathContent); got != want {
		t.Fatal("content.xml changed after read-only cell access")
	}
}

func TestAddedSheetInheritsStructuralStyles(t *testing.T) {
	src := filepath.Join("tmp", "old.ods")
	if _, err := os.Stat(src); err != nil {
		src = filepath.Join("tmp", "Blank.ods")
	}
	if _, err := os.Stat(src); err != nil {
		t.Skip("sample ODS not available")
	}

	doc, err := Open(src)
	if err != nil {
		t.Fatal(err)
	}
	doc.AddSheet("Sales")
	doc.SheetByName("Sales").Cell(0, 0).Set("Hello")
	dst := filepath.Join(t.TempDir(), "after-save.ods")
	if err := doc.SaveAs(dst); err != nil {
		t.Fatal(err)
	}

	content := zipMember(t, dst, zipPathContent)
	for _, want := range []string{
		`<table:table table:name="Sales" table:style-name="ta1">`,
		`<table:table-column table:style-name="co1" table:default-cell-style-name="Default">`,
		`<table:table-row table:style-name="ro1">`,
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("added sheet did not preserve structural style %q:\n%s", want, content)
		}
	}

	settings := zipMember(t, dst, zipPathSettings)
	if !strings.Contains(settings, `<config:config-item-map-entry config:name="Sales">`) {
		t.Fatalf("settings.xml was not synced for added sheet:\n%s", settings)
	}
}

func TestMetadataStatisticsTrackSavedContent(t *testing.T) {
	src := filepath.Join("tmp", "old.ods")
	if _, err := os.Stat(src); err != nil {
		src = filepath.Join("tmp", "Blank.ods")
	}
	if _, err := os.Stat(src); err != nil {
		t.Skip("sample ODS not available")
	}

	doc, err := Open(src)
	if err != nil {
		t.Fatal(err)
	}
	doc.AddSheet("Sales").Cell(0, 0).Set("Hello")
	dst := filepath.Join(t.TempDir(), "stats.ods")
	if err := doc.SaveAs(dst); err != nil {
		t.Fatal(err)
	}

	meta := zipMember(t, dst, zipPathMeta)
	if !strings.Contains(meta, `meta:table-count="2"`) {
		t.Fatalf("table count did not reflect added sheet:\n%s", meta)
	}
	if !strings.Contains(meta, `meta:cell-count="1"`) {
		t.Fatalf("cell count did not reflect written cell:\n%s", meta)
	}
}

func TestRepeatedTailIsNotFullyExpanded(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "sparse.ods"))
	if err != nil {
		t.Fatal(err)
	}
	doc.Sheet(0).Cell(1000, 100).Set("x")
	if rows := len(doc.content.Body.Spreadsheet.Table[0].TableRow); rows > 4 {
		t.Fatalf("materialized too many rows: %d", rows)
	}
	if cells := len(doc.content.Body.Spreadsheet.Table[0].TableRow[1].TableCell); cells > 4 {
		t.Fatalf("materialized too many cells: %d", cells)
	}
}

func TestLowercaseFormulaReferencesAreNormalized(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "lower-formula.ods"))
	if err != nil {
		t.Fatal(err)
	}
	doc.Sheet(0).Cell(0, 0).SetFormula("sum(a1:b2)")
	got, err := doc.Sheet(0).Cell(0, 0).Formula()
	if err != nil {
		t.Fatal(err)
	}
	if got != "of:=sum([.A1:.B2])" {
		t.Fatalf("formula = %q", got)
	}
}

func TestInsertRowUnsupportedValueDoesNotInsertPartialRow(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "insert-unsupported.ods"))
	if err != nil {
		t.Fatal(err)
	}
	sheet := doc.Sheet(0)
	before := sheet.RowCount()

	sheet.InsertRow(0, "ok", complex64(1+2i))

	if doc.Err() == nil {
		t.Fatal("unsupported row value did not set document error")
	}
	if got := sheet.RowCount(); got != before {
		t.Fatalf("row count changed after failed insert: got %d want %d", got, before)
	}
	if got, err := sheet.Cell(0, 0).String(); err != nil || got != "" {
		t.Fatalf("failed insert wrote partial cell: %q, %v", got, err)
	}
}

func TestStyleDeduplication(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "styles.ods"))
	if err != nil {
		t.Fatal(err)
	}
	doc.Sheet(0).Cell(0, 0).Style().Bold(true).Color("f00").Apply()
	doc.Sheet(0).Cell(0, 1).Style().Bold(true).Color("#f00").Apply()

	if got := len(doc.content.AutomaticStyles.Styles); got != 1 {
		t.Fatalf("style count = %d", got)
	}
}

func TestUnchangedStyleApplyDoesNotCreateDuplicateStyle(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "unchanged-style.ods"))
	if err != nil {
		t.Fatal(err)
	}
	sheet := doc.Sheet(0)

	sheet.Cell(0, 0).Style().Bold(true).Apply()
	sheet.Row(0).Style().Bold(true).Apply()
	sheet.Col(0).Style().Bold(true).Apply()
	styleCount := len(doc.content.AutomaticStyles.Styles)

	sheet.Cell(0, 0).Style().Apply()
	sheet.Row(0).Style().Apply()
	sheet.Col(0).Style().Apply()

	if got := len(doc.content.AutomaticStyles.Styles); got != styleCount {
		t.Fatalf("unchanged style apply created styles: got %d want %d", got, styleCount)
	}
}

func TestStyleClearResetsStyle(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "reset-style.ods"))
	if err != nil {
		t.Fatal(err)
	}
	sheet := doc.Sheet(0)

	sheet.Cell(0, 0).Style().Bold(true).Apply()
	sheet.Row(0).Style().Bold(true).Apply()
	sheet.Col(0).Style().Bold(true).Apply()

	sheet.Cell(0, 0).Style().Clear().Apply()
	sheet.Row(0).Style().Clear().Apply()
	sheet.Col(0).Style().Clear().Apply()

	table := &doc.content.Body.Spreadsheet.Table[0]
	if got := ensureCell(ensureRow(table, 0), 0).StyleName; got != "" {
		t.Fatalf("cell style was not reset: %q", got)
	}
	if got := ensureRow(table, 0).StyleName; got != "" {
		t.Fatalf("row style was not reset: %q", got)
	}
	if got := ensureColumn(table, 0).StyleName; got != "" {
		t.Fatalf("column style was not reset: %q", got)
	}
}

func TestStyleBuilderPreservesExistingProperties(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "preserve-style.ods"))
	if err != nil {
		t.Fatal(err)
	}
	sheet := doc.Sheet(0)

	sheet.Cell(0, 0).Style().Color("#ff0000").Apply()
	sheet.Cell(0, 0).Style().BackgroundColor("#ffffff").Apply()
	cell := ensureCell(ensureRow(&doc.content.Body.Spreadsheet.Table[0], 0), 0)
	cellDef := doc.styles.cellStyleDef(cell.StyleName)
	if cellDef.Color != "#ff0000" || cellDef.BackgroundColor != "#ffffff" {
		t.Fatalf("cell style was not preserved: %#v", cellDef)
	}

	sheet.Row(0).Style().Height("1cm").Bold(true).Apply()
	row := ensureRow(&doc.content.Body.Spreadsheet.Table[0], 0)
	rowDef := doc.styles.rowStyleDef(row.StyleName)
	if rowDef.Height != "1cm" || !rowDef.Bold {
		t.Fatalf("row style was not preserved: %#v", rowDef)
	}

	sheet.Col(0).Style().Width("2cm").Italic(true).Apply()
	col := ensureColumn(&doc.content.Body.Spreadsheet.Table[0], 0)
	colDef := doc.styles.colStyleDef(col.StyleName)
	if colDef.Width != "2cm" || !colDef.Italic {
		t.Fatalf("column style was not preserved: %#v", colDef)
	}
}

func TestConvenienceStyleMethodsPreserveExistingProperties(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "convenience-style.ods"))
	if err != nil {
		t.Fatal(err)
	}
	sheet := doc.Sheet(0)

	sheet.Row(0).Style().Bold(true).Apply()
	sheet.Row(0).Style().Height("1cm").Apply()
	row := ensureRow(&doc.content.Body.Spreadsheet.Table[0], 0)
	rowDef := doc.styles.rowStyleDef(row.StyleName)
	if rowDef.Height != "1cm" || !rowDef.Bold {
		t.Fatalf("row convenience style did not preserve properties: %#v", rowDef)
	}

	sheet.Col(0).Style().Italic(true).Apply()
	sheet.Col(0).Style().Width("2cm").Apply()
	col := ensureColumn(&doc.content.Body.Spreadsheet.Table[0], 0)
	colDef := doc.styles.colStyleDef(col.StyleName)
	if colDef.Width != "2cm" || !colDef.Italic {
		t.Fatalf("column convenience style did not preserve properties: %#v", colDef)
	}
}

func TestNoEmptyStringAttributes(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "attrs.ods"))
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := doc.writeTo(&buf); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), `=""`) {
		t.Fatal("archive XML contains empty-string attributes")
	}
}

func TestOpenRejectsInvalidMimetype(t *testing.T) {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	fw, err := zw.Create(zipPathMimetype)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fw.Write([]byte("application/zip")); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}

	if _, err := OpenReader(bytes.NewReader(buf.Bytes()), int64(buf.Len())); err == nil {
		t.Fatal("OpenReader accepted a non-ODS archive")
	}
}

func TestCreateWithoutPathCanSaveAs(t *testing.T) {
	doc, err := Create("")
	if err != nil {
		t.Fatal(err)
	}
	if err := doc.Save(); err == nil {
		t.Fatal("Save accepted a document without a path")
	}
	path := filepath.Join(t.TempDir(), "created-later.ods")
	if err := doc.SaveAs(path); err != nil {
		t.Fatal(err)
	}
	assertSpecZip(t, path)
}

func TestUnsupportedCellValueDoesNotModifyExistingCell(t *testing.T) {
	doc, err := Create(filepath.Join(t.TempDir(), "unsupported.ods"))
	if err != nil {
		t.Fatal(err)
	}
	cell := doc.Sheet(0).Cell(0, 0)
	cell.Set("original").Style().Bold(true).Apply()
	before := *cell.cell()

	cell.Set(complex64(1 + 2i))

	if doc.Err() == nil {
		t.Fatal("unsupported cell value did not set document error")
	}
	if after := *cell.cell(); after != before {
		t.Fatalf("unsupported cell value modified cell\nbefore: %#v\nafter:  %#v", before, after)
	}
}

func TestSettingsSheetNamesAreEscaped(t *testing.T) {
	settings := []byte(`<config:config-item-map-named config:name="Tables"><config:config-item-map-entry config:name="Sheet1"></config:config-item-map-entry></config:config-item-map-named>`)
	updated := string(syncSettingsSheets(settings, []string{`Sales "Q1" & More`}))
	if !strings.Contains(updated, `config:name="Sales &quot;Q1&quot; &amp; More"`) {
		t.Fatalf("sheet name was not escaped in settings.xml: %s", updated)
	}
	if err := xml.Unmarshal([]byte(`<root xmlns:config="x">`+updated+`</root>`), new(any)); err != nil {
		t.Fatalf("settings snippet is invalid XML: %v\n%s", err, updated)
	}
}

func TestInvalidHandlesDoNotPanic(t *testing.T) {
	var doc *Document
	if doc.SetTitle("x") != nil {
		t.Fatal("nil document SetTitle returned non-nil document")
	}
	if err := doc.Save(); err == nil {
		t.Fatalf("failed to save document: %v", err)
	}

	var sheet *Sheet
	cell := sheet.Cell(0, 0)
	if _, err := cell.String(); err == nil {
		t.Fatalf("failed to get cell string value: %v", err)
	}
	if got := cell.Style().Apply(); got != nil {
		t.Fatalf("invalid style apply returned %#v", got)
	}
	if got := sheet.InsertRow(-1); got != nil {
		t.Fatalf("nil sheet InsertRow returned %#v", got)
	}
}

func assertSpecZip(t *testing.T, path string) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	assertRawMimetypeHeader(t, raw)
	assertNoLocalExtraFields(t, raw)

	reader, err := zip.OpenReader(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	if len(reader.File) == 0 || reader.File[0].Name != zipPathMimetype {
		t.Fatalf("first zip member = %q, want mimetype", reader.File[0].Name)
	}
	if reader.File[0].Method != zip.Store {
		t.Fatalf("mimetype method = %d, want Store", reader.File[0].Method)
	}
	rc, err := reader.File[0].Open()
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(rc)
	if closeErr := rc.Close(); closeErr != nil && err == nil {
		err = closeErr
	}
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != odsMimetype {
		t.Fatalf("mimetype = %q", data)
	}

	for _, name := range []string{zipPathContent, zipPathMeta, zipPathManifest} {
		file := zipFile(reader, name)
		if file == nil {
			t.Fatalf("missing %s", name)
		}
		rc, err := file.Open()
		if err != nil {
			t.Fatal(err)
		}
		data, err := io.ReadAll(rc)
		if closeErr := rc.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			t.Fatal(err)
		}
		if err := xml.Unmarshal(data, new(any)); err != nil {
			t.Fatalf("%s is invalid XML: %v", name, err)
		}
	}
}

func assertNoLocalExtraFields(t *testing.T, raw []byte) {
	t.Helper()
	for offset := 0; offset+30 < len(raw); {
		if string(raw[offset:offset+4]) != "PK\x03\x04" {
			return
		}
		flags := binary.LittleEndian.Uint16(raw[offset+6 : offset+8])
		compressedSize := binary.LittleEndian.Uint32(raw[offset+18 : offset+22])
		nameLen := binary.LittleEndian.Uint16(raw[offset+26 : offset+28])
		extraLen := binary.LittleEndian.Uint16(raw[offset+28 : offset+30])
		nameStart := offset + 30
		nameEnd := nameStart + int(nameLen)
		if nameEnd > len(raw) {
			t.Fatal("truncated zip local header")
		}
		if extraLen != 0 {
			t.Fatalf("%s local extra length = %d, want 0", raw[nameStart:nameEnd], extraLen)
		}
		dataStart := nameEnd + int(extraLen)
		if flags&0x08 != 0 {
			next := bytes.Index(raw[dataStart:], []byte("PK\x03\x04"))
			if next < 0 {
				return
			}
			offset = dataStart + next
			continue
		}
		offset = dataStart + int(compressedSize)
	}
}

func assertRawMimetypeHeader(t *testing.T, raw []byte) {
	t.Helper()
	if len(raw) < 38 {
		t.Fatal("zip archive too small")
	}
	if string(raw[:4]) != "PK\x03\x04" {
		t.Fatal("zip does not start with a local file header")
	}
	flags := binary.LittleEndian.Uint16(raw[6:8])
	method := binary.LittleEndian.Uint16(raw[8:10])
	nameLen := binary.LittleEndian.Uint16(raw[26:28])
	extraLen := binary.LittleEndian.Uint16(raw[28:30])
	if flags&0x08 != 0 {
		t.Fatal("mimetype uses a data descriptor")
	}
	if method != zip.Store {
		t.Fatalf("mimetype local method = %d, want Store", method)
	}
	if nameLen != uint16(len(zipPathMimetype)) {
		t.Fatalf("first local file name length = %d", nameLen)
	}
	if extraLen != 0 {
		t.Fatalf("mimetype local extra length = %d, want 0", extraLen)
	}
	nameStart := 30
	nameEnd := nameStart + int(nameLen)
	dataEnd := nameEnd + len(odsMimetype)
	if len(raw) < dataEnd {
		t.Fatal("zip archive truncated before mimetype data")
	}
	if string(raw[nameStart:nameEnd]) != zipPathMimetype {
		t.Fatalf("first local file name = %q", raw[nameStart:nameEnd])
	}
	if string(raw[nameEnd:dataEnd]) != odsMimetype {
		t.Fatalf("raw mimetype payload = %q", raw[nameEnd:dataEnd])
	}
}

func zipFile(reader *zip.ReadCloser, name string) *zip.File {
	for _, file := range reader.File {
		if file.Name == name {
			return file
		}
	}
	return nil
}

func zipMember(t *testing.T, path, name string) string {
	t.Helper()
	reader, err := zip.OpenReader(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	file := zipFile(reader, name)
	if file == nil {
		t.Fatalf("missing %s", name)
	}
	rc, err := file.Open()
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(rc)
	if closeErr := rc.Close(); closeErr != nil && err == nil {
		err = closeErr
	}
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
