package ods

import "encoding/xml"

//go:generate go run ./internal/genodfxml -type-file xml_types.go -out odf_marshal_gen.go

const (
	odsMimetype  = "application/vnd.oasis.opendocument.spreadsheet"
	xmlMediaType = "text/xml"
	odfVersion   = "1.3"
	generator    = "go-ods"

	zipPathRoot     = "/"
	zipPathMimetype = "mimetype"
	zipPathContent  = "content.xml"
	zipPathMeta     = "meta.xml"
	zipPathSettings = "settings.xml"
	zipPathStyles   = "styles.xml"
	zipPathManifest = "META-INF/manifest.xml"

	cellValueTypeString  = "string"
	cellValueTypeFloat   = "float"
	cellValueTypeBoolean = "boolean"
	cellValueTypeDate    = "date"

	styleFamilyTableCell   = "table-cell"
	styleFamilyTableRow    = "table-row"
	styleFamilyTableColumn = "table-column"
	fontWeightBold         = "bold"
	fontStyleItalic        = "italic"

	defaultSheetNameFormat = "Sheet%d"
	defaultFirstSheetName  = "Sheet1"
	defaultColumnRepeat    = "1024"
	defaultRowRepeat       = "1048576"

	nsOffice   = "urn:oasis:names:tc:opendocument:xmlns:office:1.0"
	nsTable    = "urn:oasis:names:tc:opendocument:xmlns:table:1.0"
	nsText     = "urn:oasis:names:tc:opendocument:xmlns:text:1.0"
	nsStyle    = "urn:oasis:names:tc:opendocument:xmlns:style:1.0"
	nsFO       = "urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0"
	nsSVG      = "urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0"
	nsNumber   = "urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0"
	nsMeta     = "urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
	nsDC       = "http://purl.org/dc/elements/1.1/"
	nsOf       = "urn:oasis:names:tc:opendocument:xmlns:of:1.2"
	nsCalcext  = "urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0"
	nsManifest = "urn:oasis:names:tc:opendocument:xmlns:manifest:1.0"
)

// requiredZipPaths lists the zip entries that must be present in every valid ODS file.
var requiredZipPaths = [...]string{
	zipPathContent,
	zipPathMeta,
	zipPathSettings,
	zipPathStyles,
	zipPathManifest,
}

// manifestPaths lists the XML parts that are registered in the ODS manifest.
var manifestPaths = [...]string{
	zipPathContent,
	zipPathMeta,
	zipPathSettings,
	zipPathStyles,
}

// Content is the root element of an ODF spreadsheet document (office:document-content).
// Local-name xml tags allow encoding/xml to read files with any prefix; odf tags
// generate MarshalXML methods that write canonical prefixes.
//
//odf:marshal office:document-content
type Content struct {
	XMLName xml.Name `xml:"document-content"`

	Office  string `xml:"office,attr,omitempty" odf:"xmlns:office,attr,omitempty"`
	Table   string `xml:"table,attr,omitempty" odf:"xmlns:table,attr,omitempty"`
	Text    string `xml:"text,attr,omitempty" odf:"xmlns:text,attr,omitempty"`
	Style   string `xml:"style,attr,omitempty" odf:"xmlns:style,attr,omitempty"`
	FO      string `xml:"fo,attr,omitempty" odf:"xmlns:fo,attr,omitempty"`
	SVG     string `xml:"svg,attr,omitempty" odf:"xmlns:svg,attr,omitempty"`
	Number  string `xml:"number,attr,omitempty" odf:"xmlns:number,attr,omitempty"`
	MetaNS  string `xml:"meta,attr,omitempty" odf:"xmlns:meta,attr,omitempty"`
	DC      string `xml:"dc,attr,omitempty" odf:"xmlns:dc,attr,omitempty"`
	Of      string `xml:"of,attr,omitempty" odf:"xmlns:of,attr,omitempty"`
	Calcext string `xml:"calcext,attr,omitempty" odf:"xmlns:calcext,attr,omitempty"`
	Version string `xml:"version,attr,omitempty" odf:"office:version,attr,omitempty"`

	Scripts         *EmptyElement   `xml:"scripts,omitempty" odf:"office:scripts,omitempty"`
	FontFaceDecls   FontFaceDecls   `xml:"font-face-decls,omitempty" odf:"office:font-face-decls,omitempty"`
	AutomaticStyles AutomaticStyles `xml:"automatic-styles" odf:"office:automatic-styles"`
	Body            Body            `xml:"body" odf:"office:body"`
}

// Body represents the office:body element containing the spreadsheet data.
//
//odf:marshal office:body
type Body struct {
	Spreadsheet Spreadsheet `xml:"spreadsheet" odf:"office:spreadsheet"`
}

// Spreadsheet represents the office:spreadsheet element containing all tables,
// calculation settings, and named expressions.
//
//odf:marshal office:spreadsheet
type Spreadsheet struct {
	CalculationSettings *CalculationSettings `xml:"calculation-settings,omitempty" odf:"table:calculation-settings,omitempty"`
	Table               []Table              `xml:"table" odf:"table:table"`
	NamedExpressions    *EmptyElement        `xml:"named-expressions,omitempty" odf:"table:named-expressions,omitempty"`
}

type EmptyElement struct{}

// FontFaceDecls represents the office:font-face-decls element containing
// font face declarations used in the document.
//
//odf:marshal office:font-face-decls
type FontFaceDecls struct {
	FontFace []FontFace `xml:"font-face,omitempty" odf:"style:font-face,omitempty"`
}

// FontFace represents a single font face declaration within font-face-decls.
//
//odf:marshal style:font-face
type FontFace struct {
	Name              string `xml:"name,attr,omitempty" odf:"style:name,attr,omitempty"`
	FontFamily        string `xml:"font-family,attr,omitempty" odf:"svg:font-family,attr,omitempty"`
	FontFamilyGeneric string `xml:"font-family-generic,attr,omitempty" odf:"style:font-family-generic,attr,omitempty"`
	FontPitch         string `xml:"font-pitch,attr,omitempty" odf:"style:font-pitch,attr,omitempty"`
}

// CalculationSettings represents the table:calculation-settings element
// configuring spreadsheet calculation behavior.
//
//odf:marshal table:calculation-settings
type CalculationSettings struct {
	AutomaticFindLabels   string `xml:"automatic-find-labels,attr,omitempty" odf:"table:automatic-find-labels,attr,omitempty"`
	UseRegularExpressions string `xml:"use-regular-expressions,attr,omitempty" odf:"table:use-regular-expressions,attr,omitempty"`
	UseWildcards          string `xml:"use-wildcards,attr,omitempty" odf:"table:use-wildcards,attr,omitempty"`
	NullYear              string `xml:"null-year,attr,omitempty" odf:"table:null-year,attr,omitempty"`
}

// Table represents a single worksheet (table:table) within the spreadsheet.
//
//odf:marshal table:table
type Table struct {
	Name        string        `xml:"name,attr,omitempty" odf:"table:name,attr,omitempty"`
	StyleName   string        `xml:"style-name,attr,omitempty" odf:"table:style-name,attr,omitempty"`
	TableColumn []TableColumn `xml:"table-column,omitempty" odf:"table:table-column,omitempty"`
	TableRow    []TableRow    `xml:"table-row,omitempty" odf:"table:table-row,omitempty"`
}

// TableColumn represents a table:table-column element defining column properties.
//
//odf:marshal table:table-column
type TableColumn struct {
	StyleName             string `xml:"style-name,attr,omitempty" odf:"table:style-name,attr,omitempty"`
	NumberColumnsRepeated string `xml:"number-columns-repeated,attr,omitempty" odf:"table:number-columns-repeated,attr,omitempty"`
	DefaultCellStyleName  string `xml:"default-cell-style-name,attr,omitempty" odf:"table:default-cell-style-name,attr,omitempty"`
}

// TableRow represents a table:table-row element containing cells.
//
//odf:marshal table:table-row
type TableRow struct {
	StyleName          string      `xml:"style-name,attr,omitempty" odf:"table:style-name,attr,omitempty"`
	NumberRowsRepeated string      `xml:"number-rows-repeated,attr,omitempty" odf:"table:number-rows-repeated,attr,omitempty"`
	TableCell          []TableCell `xml:"table-cell,omitempty" odf:"table:table-cell,omitempty"`
}

// TableCell represents a table:table-cell element containing a single cell's
// value, formula, style, and formatting information.
//
//odf:marshal table:table-cell
type TableCell struct {
	StyleName             string `xml:"style-name,attr,omitempty" odf:"table:style-name,attr,omitempty"`
	ValueType             string `xml:"value-type,attr,omitempty" odf:"office:value-type,attr,omitempty"`
	CalcextValueType      string `xml:"-" odf:"calcext:value-type,attr,omitempty"`
	NumberColumnsRepeated string `xml:"number-columns-repeated,attr,omitempty" odf:"table:number-columns-repeated,attr,omitempty"`
	Formula               string `xml:"formula,attr,omitempty" odf:"table:formula,attr,omitempty"`
	Value                 string `xml:"value,attr,omitempty" odf:"office:value,attr,omitempty"`
	Currency              string `xml:"currency,attr,omitempty" odf:"office:currency,attr,omitempty"`
	DateValue             string `xml:"date-value,attr,omitempty" odf:"office:date-value,attr,omitempty"`
	TimeValue             string `xml:"time-value,attr,omitempty" odf:"office:time-value,attr,omitempty"`
	P                     string `xml:"p,omitempty" odf:"text:p,omitempty"`
}

// AutomaticStyles represents the office:automatic-styles element containing
// styles generated by the library (as opposed to named styles in styles.xml).
//
//odf:marshal office:automatic-styles
type AutomaticStyles struct {
	Styles []Style `xml:"style,omitempty" odf:"style:style,omitempty"`
}

// Style represents a style:style element defining visual properties for
// cells, rows, columns, or tables.
//
//odf:marshal style:style
type Style struct {
	Name                  string                 `xml:"name,attr,omitempty" odf:"style:name,attr,omitempty"`
	Family                string                 `xml:"family,attr,omitempty" odf:"style:family,attr,omitempty"`
	ParentStyleName       string                 `xml:"parent-style-name,attr,omitempty" odf:"style:parent-style-name,attr,omitempty"`
	MasterPageName        string                 `xml:"master-page-name,attr,omitempty" odf:"style:master-page-name,attr,omitempty"`
	TableProperties       *TableProperties       `xml:"table-properties,omitempty" odf:"style:table-properties,omitempty"`
	TableColumnProperties *TableColumnProperties `xml:"table-column-properties,omitempty" odf:"style:table-column-properties,omitempty"`
	TableRowProperties    *TableRowProperties    `xml:"table-row-properties,omitempty" odf:"style:table-row-properties,omitempty"`
	TableCellProperties   *TableCellProperties   `xml:"table-cell-properties,omitempty" odf:"style:table-cell-properties,omitempty"`
	ParagraphProperties   *ParagraphProperties   `xml:"paragraph-properties,omitempty" odf:"style:paragraph-properties,omitempty"`
	TextProperties        *TextProperties        `xml:"text-properties,omitempty" odf:"style:text-properties,omitempty"`
}

// TableProperties represents the style:table-properties element.
//
//odf:marshal style:table-properties
type TableProperties struct {
	Display     string `xml:"display,attr,omitempty" odf:"table:display,attr,omitempty"`
	WritingMode string `xml:"writing-mode,attr,omitempty" odf:"style:writing-mode,attr,omitempty"`
}

// TableColumnProperties represents the style:table-column-properties element.
//
//odf:marshal style:table-column-properties
type TableColumnProperties struct {
	ColumnWidth string `xml:"column-width,attr,omitempty" odf:"style:column-width,attr,omitempty"`
	BreakBefore string `xml:"break-before,attr,omitempty" odf:"fo:break-before,attr,omitempty"`
}

// TableRowProperties represents the style:table-row-properties element.
//
//odf:marshal style:table-row-properties
type TableRowProperties struct {
	RowHeight           string `xml:"row-height,attr,omitempty" odf:"style:row-height,attr,omitempty"`
	BreakBefore         string `xml:"break-before,attr,omitempty" odf:"fo:break-before,attr,omitempty"`
	UseOptimalRowHeight string `xml:"use-optimal-row-height,attr,omitempty" odf:"style:use-optimal-row-height,attr,omitempty"`
}

// TableCellProperties represents the style:table-cell-properties element.
//
//odf:marshal style:table-cell-properties
type TableCellProperties struct {
	BackgroundColor string `xml:"background-color,attr,omitempty" odf:"fo:background-color,attr,omitempty"`
	Border          string `xml:"border,attr,omitempty" odf:"fo:border,attr,omitempty"`
	VerticalAlign   string `xml:"vertical-align,attr,omitempty" odf:"style:vertical-align,attr,omitempty"`
}

// ParagraphProperties represents the style:paragraph-properties element.
//
//odf:marshal style:paragraph-properties
type ParagraphProperties struct {
	TextAlign string `xml:"text-align,attr,omitempty" odf:"fo:text-align,attr,omitempty"`
}

// TextProperties represents the style:text-properties element containing
// font, color, and weight settings.
//
//odf:marshal style:text-properties
type TextProperties struct {
	Color             string `xml:"color,attr,omitempty" odf:"fo:color,attr,omitempty"`
	FontWeight        string `xml:"font-weight,attr,omitempty" odf:"fo:font-weight,attr,omitempty"`
	FontWeightAsian   string `xml:"font-weight-asian,attr,omitempty" odf:"style:font-weight-asian,attr,omitempty"`
	FontWeightComplex string `xml:"font-weight-complex,attr,omitempty" odf:"style:font-weight-complex,attr,omitempty"`
	FontStyle         string `xml:"font-style,attr,omitempty" odf:"fo:font-style,attr,omitempty"`
	FontStyleAsian    string `xml:"font-style-asian,attr,omitempty" odf:"style:font-style-asian,attr,omitempty"`
	FontStyleComplex  string `xml:"font-style-complex,attr,omitempty" odf:"style:font-style-complex,attr,omitempty"`
	FontSize          string `xml:"font-size,attr,omitempty" odf:"fo:font-size,attr,omitempty"`
	FontFamily        string `xml:"font-family,attr,omitempty" odf:"fo:font-family,attr,omitempty"`
}

// Meta represents the root element of the ODF metadata document
// (office:document-meta).
//
//odf:marshal office:document-meta
type Meta struct {
	XMLName xml.Name `xml:"document-meta"`

	Office  string `xml:"office,attr,omitempty" odf:"xmlns:office,attr,omitempty"`
	MetaNS  string `xml:"meta,attr,omitempty" odf:"xmlns:meta,attr,omitempty"`
	DC      string `xml:"dc,attr,omitempty" odf:"xmlns:dc,attr,omitempty"`
	Version string `xml:"version,attr,omitempty" odf:"office:version,attr,omitempty"`

	Body MetaBody `xml:"meta" odf:"office:meta"`
}

// MetaBody represents the office:meta element containing document metadata
// such as title, author, dates, and statistics.
//
//odf:marshal office:meta
type MetaBody struct {
	Generator         string            `xml:"generator,omitempty" odf:"meta:generator,omitempty"`
	Title             string            `xml:"title,omitempty" odf:"dc:title,omitempty"`
	Subject           string            `xml:"subject,omitempty" odf:"dc:subject,omitempty"`
	Description       string            `xml:"description,omitempty" odf:"dc:description,omitempty"`
	Keywords          string            `xml:"keyword,omitempty" odf:"meta:keyword,omitempty"`
	InitialCreator    string            `xml:"initial-creator,omitempty" odf:"meta:initial-creator,omitempty"`
	Creator           string            `xml:"creator,omitempty" odf:"dc:creator,omitempty"`
	CreationDate      string            `xml:"creation-date,omitempty" odf:"meta:creation-date,omitempty"`
	Date              string            `xml:"date,omitempty" odf:"dc:date,omitempty"`
	Language          string            `xml:"language,omitempty" odf:"dc:language,omitempty"`
	EditingDuration   string            `xml:"editing-duration,omitempty" odf:"meta:editing-duration,omitempty"`
	EditingCycles     string            `xml:"editing-cycles,omitempty" odf:"meta:editing-cycles,omitempty"`
	DocumentStatistic DocumentStatistic `xml:"document-statistic,omitempty" odf:"meta:document-statistic,omitempty"`
}

// DocumentStatistic represents the meta:document-statistic element containing
// table, cell, and object counts.
//
//odf:marshal meta:document-statistic
type DocumentStatistic struct {
	TableCount  string `xml:"table-count,attr,omitempty" odf:"meta:table-count,attr,omitempty"`
	CellCount   string `xml:"cell-count,attr,omitempty" odf:"meta:cell-count,attr,omitempty"`
	ObjectCount string `xml:"object-count,attr,omitempty" odf:"meta:object-count,attr,omitempty"`
}

// Manifest represents the META-INF/manifest.xml root element listing all
// files in the ODS archive with their media types.
//
//odf:marshal manifest:manifest
type Manifest struct {
	XMLName xml.Name       `xml:"manifest"`
	Xmlns   string         `xml:"manifest,attr,omitempty" odf:"xmlns:manifest,attr,omitempty"`
	Version string         `xml:"version,attr,omitempty" odf:"manifest:version,attr,omitempty"`
	Files   []ManifestFile `xml:"file-entry" odf:"manifest:file-entry"`
}

// ManifestFile represents a manifest:file-entry element mapping a zip path
// to its MIME media type.
//
//odf:marshal manifest:file-entry
type ManifestFile struct {
	FullPath  string `xml:"full-path,attr" odf:"manifest:full-path,attr"`
	MediaType string `xml:"media-type,attr" odf:"manifest:media-type,attr"`
}

// newContent creates a default Content with all ODF namespace attributes
// and a single empty sheet.
func newContent() Content {
	return Content{
		Office:  nsOffice,
		Table:   nsTable,
		Text:    nsText,
		Style:   nsStyle,
		FO:      nsFO,
		SVG:     nsSVG,
		Number:  nsNumber,
		MetaNS:  nsMeta,
		DC:      nsDC,
		Of:      nsOf,
		Calcext: nsCalcext,
		Version: odfVersion,
		Scripts: &EmptyElement{},
		Body: Body{Spreadsheet: Spreadsheet{
			CalculationSettings: &CalculationSettings{},
			NamedExpressions:    &EmptyElement{},
			Table: []Table{{
				Name:        defaultFirstSheetName,
				TableColumn: []TableColumn{{NumberColumnsRepeated: defaultColumnRepeat}},
				TableRow:    []TableRow{{NumberRowsRepeated: defaultRowRepeat, TableCell: []TableCell{{NumberColumnsRepeated: defaultColumnRepeat}}}},
			}},
		}},
	}
}

// newMeta creates a default Meta with the go-ods generator identifier.
func newMeta() Meta {
	return Meta{
		Office:  nsOffice,
		MetaNS:  nsMeta,
		DC:      nsDC,
		Version: odfVersion,
		Body: MetaBody{
			Generator: generator,
		},
	}
}

// newManifest creates a default Manifest listing all standard ODS parts.
func newManifest() Manifest {
	return Manifest{
		Xmlns:   nsManifest,
		Version: odfVersion,
		Files: []ManifestFile{
			{FullPath: zipPathRoot, MediaType: odsMimetype},
			{FullPath: zipPathContent, MediaType: xmlMediaType},
			{FullPath: zipPathMeta, MediaType: xmlMediaType},
			{FullPath: zipPathSettings, MediaType: xmlMediaType},
			{FullPath: zipPathStyles, MediaType: xmlMediaType},
		},
	}
}
