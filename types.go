package ods

import "encoding/xml"

type Content struct {
	XMLName         xml.Name        `xml:"document-content" json:"document-content,omitempty"`
	Text            string          `xml:",chardata" json:"text,omitempty"`
	Presentation    string          `xml:"presentation,attr" json:"presentation,omitempty"`
	Css3t           string          `xml:"css3t,attr" json:"css3t,omitempty"`
	Grddl           string          `xml:"grddl,attr" json:"grddl,omitempty"`
	Xhtml           string          `xml:"xhtml,attr" json:"xhtml,omitempty"`
	Xsi             string          `xml:"xsi,attr" json:"xsi,omitempty"`
	Xsd             string          `xml:"xsd,attr" json:"xsd,omitempty"`
	Xforms          string          `xml:"xforms,attr" json:"xforms,omitempty"`
	Dom             string          `xml:"dom,attr" json:"dom,omitempty"`
	Script          string          `xml:"script,attr" json:"script,omitempty"`
	Form            string          `xml:"form,attr" json:"form,omitempty"`
	Math            string          `xml:"math,attr" json:"math,omitempty"`
	Office          string          `xml:"office,attr" json:"office,omitempty"`
	Ooo             string          `xml:"ooo,attr" json:"ooo,omitempty"`
	Fo              string          `xml:"fo,attr" json:"fo,omitempty"`
	Ooow            string          `xml:"ooow,attr" json:"ooow,omitempty"`
	Xlink           string          `xml:"xlink,attr" json:"xlink,omitempty"`
	Drawooo         string          `xml:"drawooo,attr" json:"drawooo,omitempty"`
	Oooc            string          `xml:"oooc,attr" json:"oooc,omitempty"`
	Dc              string          `xml:"dc,attr" json:"dc,omitempty"`
	Calcext         string          `xml:"calcext,attr" json:"calcext,omitempty"`
	Style           string          `xml:"style,attr" json:"style,omitempty"`
	AttrText        string          `xml:"text,attr" json:"attrText,omitempty"`
	Of              string          `xml:"of,attr" json:"of,omitempty"`
	Tableooo        string          `xml:"tableooo,attr" json:"tableooo,omitempty"`
	Draw            string          `xml:"draw,attr" json:"draw,omitempty"`
	Dr3d            string          `xml:"dr3d,attr" json:"dr3d,omitempty"`
	Rpt             string          `xml:"rpt,attr" json:"rpt,omitempty"`
	Formx           string          `xml:"formx,attr" json:"formx,omitempty"`
	SVG             string          `xml:"svg,attr" json:"svg,omitempty"`
	Chart           string          `xml:"chart,attr" json:"chart,omitempty"`
	Table           string          `xml:"table,attr" json:"table,omitempty"`
	Meta            string          `xml:"meta,attr" json:"meta,omitempty"`
	Loext           string          `xml:"loext,attr" json:"loext,omitempty"`
	Number          string          `xml:"number,attr" json:"number,omitempty"`
	Field           string          `xml:"field,attr" json:"field,omitempty"`
	Version         string          `xml:"version,attr" json:"version,omitempty"`
	Scripts         string          `xml:"scripts" json:"scripts,omitempty"`
	FontFaceDecls   FontFaceDecls   `xml:"font-face-decls" json:"font-face-decls,omitempty"`
	AutomaticStyles AutomaticStyles `xml:"automatic-styles" json:"automatic-styles,omitempty"`
	Body            Body            `xml:"body" json:"body,omitempty"`
}

type FontFaceDecls struct {
	Text     string     `xml:",chardata" json:"text,omitempty"`
	FontFace []FontFace `xml:"font-face" json:"font-face,omitempty"`
}

type FontFace struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	Name              string `xml:"name,attr" json:"name,omitempty"`
	FontFamily        string `xml:"font-family,attr" json:"font-family,omitempty"`
	FontFamilyGeneric string `xml:"font-family-generic,attr" json:"font-family-generic,omitempty"`
	FontPitch         string `xml:"font-pitch,attr" json:"font-pitch,omitempty"`
}

type AutomaticStyles struct {
	Text          string          `xml:",chardata" json:"text,omitempty"`
	Style         []Style         `xml:"style" json:"style,omitempty"`
	DateStyle     DateStyle       `xml:"date-style" json:"date-style,omitempty"`
	CurrencyStyle []CurrencyStyle `xml:"currency-style" json:"currency-style,omitempty"`
}

type Style struct {
	Text                  string                `xml:",chardata" json:"text,omitempty"`
	Name                  string                `xml:"name,attr" json:"name,omitempty"`
	Family                string                `xml:"family,attr" json:"family,omitempty"`
	MasterPageName        string                `xml:"master-page-name,attr" json:"master-page-name,omitempty"`
	ParentStyleName       string                `xml:"parent-style-name,attr" json:"parent-style-name,omitempty"`
	DataStyleName         string                `xml:"data-style-name,attr" json:"data-style-name,omitempty"`
	TableColumnProperties TableColumnProperties `xml:"table-column-properties" json:"table-column-properties,omitempty"`
	TableRowProperties    TableRowProperties    `xml:"table-row-properties" json:"table-row-properties,omitempty"`
	TableProperties       TableProperties       `xml:"table-properties" json:"table-properties,omitempty"`
	TableCellProperties   TableCellProperties   `xml:"table-cell-properties" json:"table-cell-properties,omitempty"`
	ParagraphProperties   ParagraphProperties   `xml:"paragraph-properties" json:"paragraph-properties,omitempty"`
	TextProperties        TextProperties        `xml:"text-properties" json:"text-properties,omitempty"`
}

type TableColumnProperties struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	BreakBefore string `xml:"break-before,attr" json:"break-before,omitempty"`
	ColumnWidth string `xml:"column-width,attr" json:"column-width,omitempty"`
}

type TableRowProperties struct {
	Text                string `xml:",chardata" json:"text,omitempty"`
	RowHeight           string `xml:"row-height,attr" json:"row-height,omitempty"`
	BreakBefore         string `xml:"break-before,attr" json:"break-before,omitempty"`
	UseOptimalRowHeight string `xml:"use-optimal-row-height,attr" json:"use-optimal-row-height,omitempty"`
}

type TableProperties struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	Display     string `xml:"display,attr" json:"display,omitempty"`
	WritingMode string `xml:"writing-mode,attr" json:"writing-mode,omitempty"`
}

type TableCellProperties struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	BackgroundColor string `xml:"background-color,attr" json:"background-color,omitempty"`
	TextAlignSource string `xml:"text-align-source,attr" json:"text-align-source,omitempty"`
	RepeatContent   string `xml:"repeat-content,attr" json:"repeat-content,omitempty"`
	Border          string `xml:"border,attr" json:"border,omitempty"`
}

type ParagraphProperties struct {
	Text       string `xml:",chardata" json:"text,omitempty"`
	TextAlign  string `xml:"text-align,attr" json:"text-align,omitempty"`
	MarginLeft string `xml:"margin-left,attr" json:"margin-left,omitempty"`
}

type TextProperties struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	FontSize          string `xml:"font-size,attr" json:"font-size,omitempty"`
	FontWeight        string `xml:"font-weight,attr" json:"font-weight,omitempty"`
	FontSizeAsian     string `xml:"font-size-asian,attr" json:"font-size-asian,omitempty"`
	FontWeightAsian   string `xml:"font-weight-asian,attr" json:"font-weight-asian,omitempty"`
	FontSizeComplex   string `xml:"font-size-complex,attr" json:"font-size-complex,omitempty"`
	FontWeightComplex string `xml:"font-weight-complex,attr" json:"font-weight-complex,omitempty"`
}

type DateStyle struct {
	Chardata       string   `xml:",chardata" json:"chardata,omitempty"`
	Name           string   `xml:"name,attr" json:"name,omitempty"`
	AutomaticOrder string   `xml:"automatic-order,attr" json:"automatic-order,omitempty"`
	Month          Month    `xml:"month" json:"month,omitempty"`
	Text           []string `xml:"text" json:"text,omitempty"`
	Day            Day      `xml:"day" json:"day,omitempty"`
	Year           string   `xml:"year" json:"year,omitempty"`
}

type Month struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"style,attr" json:"style,omitempty"`
}

type Day struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"style,attr" json:"style,omitempty"`
}

type CurrencyStyle struct {
	Chardata       string          `xml:",chardata" json:"chardata,omitempty"`
	Name           string          `xml:"name,attr" json:"name,omitempty"`
	Volatile       string          `xml:"volatile,attr" json:"volatile,omitempty"`
	Language       string          `xml:"language,attr" json:"language,omitempty"`
	Country        string          `xml:"country,attr" json:"country,omitempty"`
	CurrencySymbol CurrencySymbol  `xml:"currency-symbol" json:"currency-symbol,omitempty"`
	Number         Number          `xml:"number" json:"number,omitempty"`
	TextProperties TextProperties0 `xml:"text-properties" json:"text-properties,omitempty"`
	Text           string          `xml:"text" json:"text,omitempty"`
	Map            Map             `xml:"map" json:"map,omitempty"`
}

type CurrencySymbol struct {
	Text     string `xml:",chardata" json:"text,omitempty"`
	Language string `xml:"language,attr" json:"language,omitempty"`
	Country  string `xml:"country,attr" json:"country,omitempty"`
}

type Number struct {
	Text             string `xml:",chardata" json:"text,omitempty"`
	DecimalPlaces    string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
	MinDecimalPlaces string `xml:"min-decimal-places,attr" json:"min-decimal-places,omitempty"`
	MinIntegerDigits string `xml:"min-integer-digits,attr" json:"min-integer-digits,omitempty"`
	Grouping         string `xml:"grouping,attr" json:"grouping,omitempty"`
}

type TextProperties0 struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Color string `xml:"color,attr" json:"color,omitempty"`
}

type Map struct {
	Text           string `xml:",chardata" json:"text,omitempty"`
	Condition      string `xml:"condition,attr" json:"condition,omitempty"`
	ApplyStyleName string `xml:"apply-style-name,attr" json:"apply-style-name,omitempty"`
}

type Body struct {
	Text        string      `xml:",chardata" json:"text,omitempty"`
	Spreadsheet Spreadsheet `xml:"spreadsheet" json:"spreadsheet,omitempty"`
}

type Spreadsheet struct {
	Text                string              `xml:",chardata" json:"text,omitempty"`
	CalculationSettings CalculationSettings `xml:"calculation-settings" json:"calculation-settings,omitempty"`
	Table               []Table             `xml:"table" json:"table,omitempty"`
	NamedExpressions    string              `xml:"named-expressions" json:"named-expressions,omitempty"`
}

type CalculationSettings struct {
	Text                  string    `xml:",chardata" json:"text,omitempty"`
	CaseSensitive         string    `xml:"case-sensitive,attr" json:"case-sensitive,omitempty"`
	AutomaticFindLabels   string    `xml:"automatic-find-labels,attr" json:"automatic-find-labels,omitempty"`
	UseRegularExpressions string    `xml:"use-regular-expressions,attr" json:"use-regular-expressions,omitempty"`
	UseWildcards          string    `xml:"use-wildcards,attr" json:"use-wildcards,omitempty"`
	Iteration             Iteration `xml:"iteration" json:"iteration,omitempty"`
}

type Iteration struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	MaximumDifference string `xml:"maximum-difference,attr" json:"maximum-difference,omitempty"`
}

type Table struct {
	Text        string        `xml:",chardata" json:"text,omitempty"`
	Name        string        `xml:"name,attr" json:"name,omitempty"`
	StyleName   string        `xml:"style-name,attr" json:"style-name,omitempty"`
	Forms       Forms         `xml:"forms" json:"forms,omitempty"`
	TableColumn []TableColumn `xml:"table-column" json:"table-column,omitempty"`
	TableRow    []TableRow    `xml:"table-row" json:"table-row,omitempty"`
}

type Forms struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	AutomaticFocus  string `xml:"automatic-focus,attr" json:"automatic-focus,omitempty"`
	ApplyDesignMode string `xml:"apply-design-mode,attr" json:"apply-design-mode,omitempty"`
}

type TableColumn struct {
	Text                  string `xml:",chardata" json:"text,omitempty"`
	StyleName             string `xml:"style-name,attr" json:"style-name,omitempty"`
	NumberColumnsRepeated string `xml:"number-columns-repeated,attr" json:"number-columns-repeated,omitempty"`
	DefaultCellStyleName  string `xml:"default-cell-style-name,attr" json:"default-cell-style-name,omitempty"`
}

type TableRow struct {
	Text               string      `xml:",chardata" json:"text,omitempty"`
	StyleName          string      `xml:"style-name,attr" json:"style-name,omitempty"`
	NumberRowsRepeated string      `xml:"number-rows-repeated,attr" json:"number-rows-repeated,omitempty"`
	TableCell          []TableCell `xml:"table-cell" json:"table-cell,omitempty"`
}

type TableCell struct {
	Text                  string `xml:",chardata" json:"text,omitempty"`
	StyleName             string `xml:"style-name,attr" json:"style-name,omitempty"`
	ValueType             string `xml:"value-type,attr" json:"value-type,omitempty"`
	ValueType0            string `json:"value-type0,omitempty"`
	NumberColumnsRepeated string `xml:"number-columns-repeated,attr" json:"number-columns-repeated,omitempty"`
	Formula               string `xml:"formula,attr" json:"formula,omitempty"`
	Value                 string `xml:"value,attr" json:"value,omitempty"`
	Currency              string `xml:"currency,attr" json:"currency,omitempty"`
	DateValue             string `xml:"date-value,attr" json:"date-value,omitempty"`
	P                     string `xml:"p" json:"p,omitempty"`
}

type ContentMarshal struct {
	XMLName         xml.Name               `xml:"office:document-content" json:"document-content,omitempty"`
	Text            string                 `xml:",chardata" json:"text,omitempty"`
	Presentation    string                 `xml:"xmlns:presentation,attr" json:"presentation,omitempty"`
	Css3t           string                 `xml:"xmlns:css3t,attr" json:"css3t,omitempty"`
	Grddl           string                 `xml:"xmlns:grddl,attr" json:"grddl,omitempty"`
	Xhtml           string                 `xml:"xmlns:xhtml,attr" json:"xhtml,omitempty"`
	Xsi             string                 `xml:"xmlns:xsi,attr" json:"xsi,omitempty"`
	Xsd             string                 `xml:"xmlns:xsd,attr" json:"xsd,omitempty"`
	Xforms          string                 `xml:"xmlns:xforms,attr" json:"xforms,omitempty"`
	Dom             string                 `xml:"xmlns:dom,attr" json:"dom,omitempty"`
	Script          string                 `xml:"xmlns:script,attr" json:"script,omitempty"`
	Form            string                 `xml:"xmlns:form,attr" json:"form,omitempty"`
	Math            string                 `xml:"xmlns:math,attr" json:"math,omitempty"`
	Office          string                 `xml:"xmlns:office,attr" json:"office,omitempty"`
	Ooo             string                 `xml:"xmlns:ooo,attr" json:"ooo,omitempty"`
	Fo              string                 `xml:"xmlns:fo,attr" json:"fo,omitempty"`
	Ooow            string                 `xml:"xmlns:ooow,attr" json:"ooow,omitempty"`
	Xlink           string                 `xml:"xmlns:xlink,attr" json:"xlink,omitempty"`
	Drawooo         string                 `xml:"xmlns:drawooo,attr" json:"drawooo,omitempty"`
	Oooc            string                 `xml:"xmlns:oooc,attr" json:"oooc,omitempty"`
	Dc              string                 `xml:"xmlns:dc,attr" json:"dc,omitempty"`
	Calcext         string                 `xml:"xmlns:calcext,attr" json:"calcext,omitempty"`
	Style           string                 `xml:"xmlns:style,attr" json:"style,omitempty"`
	AttrText        string                 `xml:"xmlns:text,attr" json:"attrText,omitempty"`
	Of              string                 `xml:"xmlns:of,attr" json:"of,omitempty"`
	Tableooo        string                 `xml:"xmlns:tableooo,attr" json:"tableooo,omitempty"`
	Draw            string                 `xml:"xmlns:draw,attr" json:"draw,omitempty"`
	Dr3d            string                 `xml:"xmlns:dr3d,attr" json:"dr3d,omitempty"`
	Rpt             string                 `xml:"xmlns:rpt,attr" json:"rpt,omitempty"`
	Formx           string                 `xml:"xmlns:formx,attr" json:"formx,omitempty"`
	SVG             string                 `xml:"xmlns:svg,attr" json:"svg,omitempty"`
	Chart           string                 `xml:"xmlns:chart,attr" json:"chart,omitempty"`
	Table           string                 `xml:"xmlns:table,attr" json:"table,omitempty"`
	Meta            string                 `xml:"xmlns:meta,attr" json:"meta,omitempty"`
	Loext           string                 `xml:"xmlns:loext,attr" json:"loext,omitempty"`
	Number          string                 `xml:"xmlns:number,attr" json:"number,omitempty"`
	Field           string                 `xml:"xmlns:field,attr" json:"field,omitempty"`
	Version         string                 `xml:"office:version,attr" json:"version,omitempty"`
	Scripts         string                 `xml:"office:scripts" json:"scripts,omitempty"`
	FontFaceDecls   FontFaceDeclsMarshal   `xml:"office:font-face-decls" json:"font-face-decls,omitempty"`
	AutomaticStyles AutomaticStylesMarshal `xml:"office:automatic-styles" json:"automatic-styles,omitempty"`
	Body            BodyMarshal            `xml:"office:body" json:"body,omitempty"`
}

type FontFaceDeclsMarshal struct {
	Text     string            `xml:",chardata" json:"text,omitempty"`
	FontFace []FontFaceMarshal `xml:"style:font-face" json:"font-face,omitempty"`
}

type FontFaceMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	Name              string `xml:"style:name,attr" json:"name,omitempty"`
	FontFamily        string `xml:"svg:font-family,attr" json:"font-family,omitempty"`
	FontFamilyGeneric string `xml:"style:font-family-generic,attr" json:"font-family-generic,omitempty"`
	FontPitch         string `xml:"style:font-pitch,attr" json:"font-pitch,omitempty"`
}

type AutomaticStylesMarshal struct {
	Text          string                 `xml:",chardata" json:"text,omitempty"`
	Style         []StyleMarshal         `xml:"style:style" json:"style,omitempty"`
	DateStyle     DateStyleMarshal       `xml:"number:date-style" json:"date-style,omitempty"`
	CurrencyStyle []CurrencyStyleMarshal `xml:"number:currency-style" json:"currency-style,omitempty"`
}

type StyleMarshal struct {
	Text                  string                       `xml:",chardata" json:"text,omitempty"`
	Name                  string                       `xml:"style:name,attr" json:"name,omitempty"`
	Family                string                       `xml:"style:family,attr" json:"family,omitempty"`
	MasterPageName        string                       `xml:"style:master-page-name,attr" json:"master-page-name,omitempty"`
	ParentStyleName       string                       `xml:"style:parent-style-name,attr" json:"parent-style-name,omitempty"`
	DataStyleName         string                       `xml:"style:data-style-name,attr" json:"data-style-name,omitempty"`
	TableColumnProperties TableColumnPropertiesMarshal `xml:"style:table-column-properties" json:"table-column-properties,omitempty"`
	TableRowProperties    TableRowPropertiesMarshal    `xml:"style:table-row-properties" json:"table-row-properties,omitempty"`
	TableProperties       TablePropertiesMarshal       `xml:"style:table-properties" json:"table-properties,omitempty"`
	TableCellProperties   TableCellPropertiesMarshal   `xml:"style:table-cell-properties" json:"table-cell-properties,omitempty"`
	ParagraphProperties   ParagraphPropertiesMarshal   `xml:"style:paragraph-properties" json:"paragraph-properties,omitempty"`
	TextProperties        TextPropertiesMarshal        `xml:"style:text-properties" json:"text-properties,omitempty"`
}

type TableColumnPropertiesMarshal struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	BreakBefore string `xml:"fo:break-before,attr" json:"break-before,omitempty"`
	ColumnWidth string `xml:"style:column-width,attr" json:"column-width,omitempty"`
}

type TableRowPropertiesMarshal struct {
	Text                string `xml:",chardata" json:"text,omitempty"`
	RowHeight           string `xml:"style:row-height,attr" json:"row-height,omitempty"`
	BreakBefore         string `xml:"fo:break-before,attr" json:"break-before,omitempty"`
	UseOptimalRowHeight string `xml:"style:use-optimal-row-height,attr" json:"use-optimal-row-height,omitempty"`
}

type TablePropertiesMarshal struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	Display     string `xml:"table:display,attr" json:"display,omitempty"`
	WritingMode string `xml:"style:writing-mode,attr" json:"writing-mode,omitempty"`
}

type TableCellPropertiesMarshal struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	BackgroundColor string `xml:"fo:background-color,attr" json:"background-color,omitempty"`
	TextAlignSource string `xml:"style:text-align-source,attr" json:"text-align-source,omitempty"`
	RepeatContent   string `xml:"style:repeat-content,attr" json:"repeat-content,omitempty"`
	Border          string `xml:"fo:border,attr" json:"border,omitempty"`
}

type ParagraphPropertiesMarshal struct {
	Text       string `xml:",chardata" json:"text,omitempty"`
	TextAlign  string `xml:"fo:text-align,attr" json:"text-align,omitempty"`
	MarginLeft string `xml:"fo:margin-left,attr" json:"margin-left,omitempty"`
}

type TextPropertiesMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	FontSize          string `xml:"fo:font-size,attr" json:"font-size,omitempty"`
	FontWeight        string `xml:"fo:font-weight,attr" json:"font-weight,omitempty"`
	FontSizeAsian     string `xml:"style:font-size-asian,attr" json:"font-size-asian,omitempty"`
	FontWeightAsian   string `xml:"style:font-weight-asian,attr" json:"font-weight-asian,omitempty"`
	FontSizeComplex   string `xml:"style:font-size-complex,attr" json:"font-size-complex,omitempty"`
	FontWeightComplex string `xml:"style:font-weight-complex,attr" json:"font-weight-complex,omitempty"`
}

type DateStyleMarshal struct {
	Chardata       string       `xml:",chardata" json:"chardata,omitempty"`
	Name           string       `xml:"style:name,attr" json:"name,omitempty"`
	AutomaticOrder string       `xml:"number:automatic-order,attr" json:"automatic-order,omitempty"`
	Month          MonthMarshal `xml:"number:month" json:"month,omitempty"`
	Text           []string     `xml:"number:text" json:"text,omitempty"`
	Day            DayMarshal   `xml:"number:day" json:"day,omitempty"`
	Year           string       `xml:"number:year" json:"year,omitempty"`
}

type MonthMarshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"number:style,attr" json:"style,omitempty"`
}

type DayMarshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"number:style,attr" json:"style,omitempty"`
}

type CurrencyStyleMarshal struct {
	Chardata       string                 `xml:",chardata" json:"chardata,omitempty"`
	Name           string                 `xml:"number:name,attr" json:"name,omitempty"`
	Volatile       string                 `xml:"style:volatile,attr" json:"volatile,omitempty"`
	Language       string                 `xml:"number:language,attr" json:"language,omitempty"`
	Country        string                 `xml:"number:country,attr" json:"country,omitempty"`
	CurrencySymbol CurrencySymbolMarshal  `xml:"number:currency-symbol" json:"currency-symbol,omitempty"`
	Number         NumberMarshal          `xml:"number:number" json:"number,omitempty"`
	TextProperties TextProperties0Marshal `xml:"style:text-properties" json:"text-properties,omitempty"`
	Text           string                 `xml:"number:text" json:"text,omitempty"`
	Map            MapMarshal             `xml:"style:map" json:"map,omitempty"`
}

type CurrencySymbolMarshal struct {
	Text     string `xml:",chardata" json:"text,omitempty"`
	Language string `xml:"number:language,attr" json:"language,omitempty"`
	Country  string `xml:"number:country,attr" json:"country,omitempty"`
}

type NumberMarshal struct {
	Text             string `xml:",chardata" json:"text,omitempty"`
	DecimalPlaces    string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
	MinDecimalPlaces string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
	MinIntegerDigits string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
	Grouping         string `xml:"number:grouping,attr" json:"grouping,omitempty"`
}

type TextProperties0Marshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Color string `xml:"fo:color,attr" json:"color,omitempty"`
}

type MapMarshal struct {
	Text           string `xml:",chardata" json:"text,omitempty"`
	Condition      string `xml:"style:condition,attr" json:"condition,omitempty"`
	ApplyStyleName string `xml:"style:apply-style-name,attr" json:"apply-style-name,omitempty"`
}

type BodyMarshal struct {
	Text        string             `xml:",chardata" json:"text,omitempty"`
	Spreadsheet SpreadsheetMarshal `xml:"office:spreadsheet" json:"spreadsheet,omitempty"`
}

type SpreadsheetMarshal struct {
	Text                string                     `xml:",chardata" json:"text,omitempty"`
	CalculationSettings CalculationSettingsMarshal `xml:"table:calculation-settings" json:"calculation-settings,omitempty"`
	Table               []TableMarshal             `xml:"table:table" json:"table,omitempty"`
	NamedExpressions    string                     `xml:"table:named-expressions" json:"named-expressions,omitempty"`
}

type CalculationSettingsMarshal struct {
	Text                  string           `xml:",chardata" json:"text,omitempty"`
	CaseSensitive         string           `xml:"table:case-sensitive,attr" json:"case-sensitive,omitempty"`
	AutomaticFindLabels   string           `xml:"table:automatic-find-labels,attr" json:"automatic-find-labels,omitempty"`
	UseRegularExpressions string           `xml:"table:use-regular-expressions,attr" json:"use-regular-expressions,omitempty"`
	UseWildcards          string           `xml:"table:use-wildcards,attr" json:"use-wildcards,omitempty"`
	Iteration             IterationMarshal `xml:"table:iteration" json:"iteration,omitempty"`
}

type IterationMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	MaximumDifference string `xml:"table:maximum-difference,attr" json:"maximum-difference,omitempty"`
}

type TableMarshal struct {
	Text        string               `xml:",chardata" json:"text,omitempty"`
	Name        string               `xml:"table:name,attr" json:"name,omitempty"`
	StyleName   string               `xml:"table:style-name,attr" json:"style-name,omitempty"`
	Forms       FormsMarshal         `xml:"office:forms" json:"forms,omitempty"`
	TableColumn []TableColumnMarshal `xml:"table:table-column" json:"table-column,omitempty"`
	TableRow    []TableRowMarshal    `xml:"table:table-row" json:"table-row,omitempty"`
}

type FormsMarshal struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	AutomaticFocus  string `xml:"form:automatic-focus,attr" json:"automatic-focus,omitempty"`
	ApplyDesignMode string `xml:"form:apply-design-mode,attr" json:"apply-design-mode,omitempty"`
}

type TableColumnMarshal struct {
	Text                  string `xml:",chardata" json:"text,omitempty"`
	StyleName             string `xml:"table:style-name,attr" json:"style-name,omitempty"`
	NumberColumnsRepeated string `xml:"table:number-columns-repeated,attr" json:"number-columns-repeated,omitempty"`
	DefaultCellStyleName  string `xml:"table:default-cell-style-name,attr" json:"default-cell-style-name,omitempty"`
}

type TableRowMarshal struct {
	Text               string             `xml:",chardata" json:"text,omitempty"`
	StyleName          string             `xml:"table:style-name,attr" json:"style-name,omitempty"`
	NumberRowsRepeated string             `xml:"table:number-rows-repeated,attr" json:"number-rows-repeated,omitempty"`
	TableCell          []TableCellMarshal `xml:"table:table-cell" json:"table-cell,omitempty"`
}

type TableCellMarshal struct {
	Text                  string `xml:",chardata" json:"text,omitempty"`
	StyleName             string `xml:"table:style-name,attr" json:"style-name,omitempty"`
	ValueType             string `xml:"office:value-type,attr" json:"value-type,omitempty"`
	ValueType0            string `xml:"calcext:value-type,attr" json:"value-type0,omitempty"`
	NumberColumnsRepeated string `xml:"table:number-columns-repeated,attr" json:"number-columns-repeated,omitempty"`
	Formula               string `xml:"table:formula,attr" json:"formula,omitempty"`
	Value                 string `xml:"office:value,attr" json:"value,omitempty"`
	Currency              string `xml:"office:currency,attr" json:"currency,omitempty"`
	DateValue             string `xml:"office:date-value,attr" json:"date-value,omitempty"`
	P                     string `xml:"text:p" json:"p,omitempty"`
}
