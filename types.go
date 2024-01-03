package ods

import "encoding/xml"

type ODS struct {
	Content  Content  `json:"content"`
	Meta     Meta     `json:"meta"`
	Manifest Manifest `json:"manifest"`
	Mimetype Mimetype `json:"mimetype"`
	Settings Settings `json:"settings"`
	Styles   Styles   `json:"styles"`
}

type odsMarshal struct {
	Content  contentMarshal  `json:"content"`
	Meta     metaMarshal     `json:"meta"`
	Manifest manifestMarshal `json:"manifest"`
	Mimetype Mimetype        `json:"mimetype"`
	Settings settingsMarshal `json:"settings"`
	Styles   stylesMarshal   `json:"styles"`
}

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

type contentMarshal struct {
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
	FontFaceDecls   fontFaceDeclsMarshal   `xml:"office:font-face-decls" json:"font-face-decls,omitempty"`
	AutomaticStyles automaticStylesMarshal `xml:"office:automatic-styles" json:"automatic-styles,omitempty"`
	Body            bodyMarshal            `xml:"office:body" json:"body,omitempty"`
}

type fontFaceDeclsMarshal struct {
	Text     string            `xml:",chardata" json:"text,omitempty"`
	FontFace []fontFaceMarshal `xml:"style:font-face" json:"font-face,omitempty"`
}

type fontFaceMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	Name              string `xml:"style:name,attr" json:"name,omitempty"`
	FontFamily        string `xml:"svg:font-family,attr" json:"font-family,omitempty"`
	FontFamilyGeneric string `xml:"style:font-family-generic,attr" json:"font-family-generic,omitempty"`
	FontPitch         string `xml:"style:font-pitch,attr" json:"font-pitch,omitempty"`
}

type automaticStylesMarshal struct {
	Text          string                 `xml:",chardata" json:"text,omitempty"`
	Style         []styleMarshal         `xml:"style:style" json:"style,omitempty"`
	DateStyle     dateStyleMarshal       `xml:"number:date-style" json:"date-style,omitempty"`
	CurrencyStyle []CurrencyStyleMarshal `xml:"number:currency-style" json:"currency-style,omitempty"`
}

type styleMarshal struct {
	Text                  string                       `xml:",chardata" json:"text,omitempty"`
	Name                  string                       `xml:"style:name,attr" json:"name,omitempty"`
	Family                string                       `xml:"style:family,attr" json:"family,omitempty"`
	MasterPageName        string                       `xml:"style:master-page-name,attr" json:"master-page-name,omitempty"`
	ParentStyleName       string                       `xml:"style:parent-style-name,attr" json:"parent-style-name,omitempty"`
	DataStyleName         string                       `xml:"style:data-style-name,attr" json:"data-style-name,omitempty"`
	TableColumnProperties tableColumnPropertiesMarshal `xml:"style:table-column-properties" json:"table-column-properties,omitempty"`
	TableRowProperties    tableRowPropertiesMarshal    `xml:"style:table-row-properties" json:"table-row-properties,omitempty"`
	TableProperties       tablePropertiesMarshal       `xml:"style:table-properties" json:"table-properties,omitempty"`
	TableCellProperties   tableCellPropertiesMarshal   `xml:"style:table-cell-properties" json:"table-cell-properties,omitempty"`
	ParagraphProperties   paragraphPropertiesMarshal   `xml:"style:paragraph-properties" json:"paragraph-properties,omitempty"`
	TextProperties        textPropertiesMarshal        `xml:"style:text-properties" json:"text-properties,omitempty"`
}

type tableColumnPropertiesMarshal struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	BreakBefore string `xml:"fo:break-before,attr" json:"break-before,omitempty"`
	ColumnWidth string `xml:"style:column-width,attr" json:"column-width,omitempty"`
}

type tableRowPropertiesMarshal struct {
	Text                string `xml:",chardata" json:"text,omitempty"`
	RowHeight           string `xml:"style:row-height,attr" json:"row-height,omitempty"`
	BreakBefore         string `xml:"fo:break-before,attr" json:"break-before,omitempty"`
	UseOptimalRowHeight string `xml:"style:use-optimal-row-height,attr" json:"use-optimal-row-height,omitempty"`
}

type tablePropertiesMarshal struct {
	Text        string `xml:",chardata" json:"text,omitempty"`
	Display     string `xml:"table:display,attr" json:"display,omitempty"`
	WritingMode string `xml:"style:writing-mode,attr" json:"writing-mode,omitempty"`
}

type tableCellPropertiesMarshal struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	BackgroundColor string `xml:"fo:background-color,attr" json:"background-color,omitempty"`
	TextAlignSource string `xml:"style:text-align-source,attr" json:"text-align-source,omitempty"`
	RepeatContent   string `xml:"style:repeat-content,attr" json:"repeat-content,omitempty"`
	Border          string `xml:"fo:border,attr" json:"border,omitempty"`
}

type paragraphPropertiesMarshal struct {
	Text       string `xml:",chardata" json:"text,omitempty"`
	TextAlign  string `xml:"fo:text-align,attr" json:"text-align,omitempty"`
	MarginLeft string `xml:"fo:margin-left,attr" json:"margin-left,omitempty"`
}

type textPropertiesMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	FontSize          string `xml:"fo:font-size,attr" json:"font-size,omitempty"`
	FontWeight        string `xml:"fo:font-weight,attr" json:"font-weight,omitempty"`
	FontSizeAsian     string `xml:"style:font-size-asian,attr" json:"font-size-asian,omitempty"`
	FontWeightAsian   string `xml:"style:font-weight-asian,attr" json:"font-weight-asian,omitempty"`
	FontSizeComplex   string `xml:"style:font-size-complex,attr" json:"font-size-complex,omitempty"`
	FontWeightComplex string `xml:"style:font-weight-complex,attr" json:"font-weight-complex,omitempty"`
}

type dateStyleMarshal struct {
	Chardata       string       `xml:",chardata" json:"chardata,omitempty"`
	Name           string       `xml:"style:name,attr" json:"name,omitempty"`
	AutomaticOrder string       `xml:"number:automatic-order,attr" json:"automatic-order,omitempty"`
	Month          monthMarshal `xml:"number:month" json:"month,omitempty"`
	Text           []string     `xml:"number:text" json:"text,omitempty"`
	Day            dayMarshal   `xml:"number:day" json:"day,omitempty"`
	Year           string       `xml:"number:year" json:"year,omitempty"`
}

type monthMarshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"number:style,attr" json:"style,omitempty"`
}

type dayMarshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Style string `xml:"number:style,attr" json:"style,omitempty"`
}

type CurrencyStyleMarshal struct {
	Chardata       string                 `xml:",chardata" json:"chardata,omitempty"`
	Name           string                 `xml:"number:name,attr" json:"name,omitempty"`
	Volatile       string                 `xml:"style:volatile,attr" json:"volatile,omitempty"`
	Language       string                 `xml:"number:language,attr" json:"language,omitempty"`
	Country        string                 `xml:"number:country,attr" json:"country,omitempty"`
	CurrencySymbol currencySymbolMarshal  `xml:"number:currency-symbol" json:"currency-symbol,omitempty"`
	Number         numberMarshal          `xml:"number:number" json:"number,omitempty"`
	TextProperties textProperties0Marshal `xml:"style:text-properties" json:"text-properties,omitempty"`
	Text           string                 `xml:"number:text" json:"text,omitempty"`
	Map            mapMarshal             `xml:"style:map" json:"map,omitempty"`
}

type currencySymbolMarshal struct {
	Text     string `xml:",chardata" json:"text,omitempty"`
	Language string `xml:"number:language,attr" json:"language,omitempty"`
	Country  string `xml:"number:country,attr" json:"country,omitempty"`
}

type numberMarshal struct {
	Text             string `xml:",chardata" json:"text,omitempty"`
	DecimalPlaces    string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
	MinDecimalPlaces string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
	MinIntegerDigits string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
	Grouping         string `xml:"number:grouping,attr" json:"grouping,omitempty"`
}

type textProperties0Marshal struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Color string `xml:"fo:color,attr" json:"color,omitempty"`
}

type mapMarshal struct {
	Text           string `xml:",chardata" json:"text,omitempty"`
	Condition      string `xml:"style:condition,attr" json:"condition,omitempty"`
	ApplyStyleName string `xml:"style:apply-style-name,attr" json:"apply-style-name,omitempty"`
}

type bodyMarshal struct {
	Text        string             `xml:",chardata" json:"text,omitempty"`
	Spreadsheet spreadsheetMarshal `xml:"office:spreadsheet" json:"spreadsheet,omitempty"`
}

type spreadsheetMarshal struct {
	Text                string                     `xml:",chardata" json:"text,omitempty"`
	CalculationSettings calculationSettingsMarshal `xml:"table:calculation-settings" json:"calculation-settings,omitempty"`
	Table               []tableMarshal             `xml:"table:table" json:"table,omitempty"`
	NamedExpressions    string                     `xml:"table:named-expressions" json:"named-expressions,omitempty"`
}

type calculationSettingsMarshal struct {
	Text                  string           `xml:",chardata" json:"text,omitempty"`
	CaseSensitive         string           `xml:"table:case-sensitive,attr" json:"case-sensitive,omitempty"`
	AutomaticFindLabels   string           `xml:"table:automatic-find-labels,attr" json:"automatic-find-labels,omitempty"`
	UseRegularExpressions string           `xml:"table:use-regular-expressions,attr" json:"use-regular-expressions,omitempty"`
	UseWildcards          string           `xml:"table:use-wildcards,attr" json:"use-wildcards,omitempty"`
	Iteration             iterationMarshal `xml:"table:iteration" json:"iteration,omitempty"`
}

type iterationMarshal struct {
	Text              string `xml:",chardata" json:"text,omitempty"`
	MaximumDifference string `xml:"table:maximum-difference,attr" json:"maximum-difference,omitempty"`
}

type tableMarshal struct {
	Text        string               `xml:",chardata" json:"text,omitempty"`
	Name        string               `xml:"table:name,attr" json:"name,omitempty"`
	StyleName   string               `xml:"table:style-name,attr" json:"style-name,omitempty"`
	Forms       formsMarshal         `xml:"office:forms" json:"forms,omitempty"`
	TableColumn []tableColumnMarshal `xml:"table:table-column" json:"table-column,omitempty"`
	TableRow    []tableRowMarshal    `xml:"table:table-row" json:"table-row,omitempty"`
}

type formsMarshal struct {
	Text            string `xml:",chardata" json:"text,omitempty"`
	AutomaticFocus  string `xml:"form:automatic-focus,attr" json:"automatic-focus,omitempty"`
	ApplyDesignMode string `xml:"form:apply-design-mode,attr" json:"apply-design-mode,omitempty"`
}

type tableColumnMarshal struct {
	Text                  string `xml:",chardata" json:"text,omitempty"`
	StyleName             string `xml:"table:style-name,attr" json:"style-name,omitempty"`
	NumberColumnsRepeated string `xml:"table:number-columns-repeated,attr" json:"number-columns-repeated,omitempty"`
	DefaultCellStyleName  string `xml:"table:default-cell-style-name,attr" json:"default-cell-style-name,omitempty"`
}

type tableRowMarshal struct {
	Text               string             `xml:",chardata" json:"text,omitempty"`
	StyleName          string             `xml:"table:style-name,attr" json:"style-name,omitempty"`
	NumberRowsRepeated string             `xml:"table:number-rows-repeated,attr" json:"number-rows-repeated,omitempty"`
	TableCell          []tableCellMarshal `xml:"table:table-cell" json:"table-cell,omitempty"`
}

type tableCellMarshal struct {
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

type Manifest struct {
	XMLName     xml.Name `xml:"RDF" json:"rdf,omitempty"`
	Text        string   `xml:",chardata" json:"text,omitempty"`
	Rdf         string   `xml:"rdf,attr" json:"rdf0,omitempty"`
	Description []struct {
		Text  string `xml:",chardata" json:"text,omitempty"`
		About string `xml:"about,attr" json:"about,omitempty"`
		Type  struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			Resource string `xml:"resource,attr" json:"resource,omitempty"`
		} `xml:"type" json:"type,omitempty"`
		HasPart struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			Ns0      string `xml:"ns0,attr" json:"ns0,omitempty"`
			Resource string `xml:"resource,attr" json:"resource,omitempty"`
		} `xml:"hasPart" json:"haspart,omitempty"`
	} `xml:"Description" json:"description,omitempty"`
}

type manifestMarshal struct {
	XMLName     xml.Name `xml:"rdf:RDF" json:"rdf,omitempty"`
	Text        string   `xml:",chardata" json:"text,omitempty"`
	Rdf         string   `xml:"xmlns:rdf,attr" json:"rdf0,omitempty"`
	Description []struct {
		Text  string `xml:",chardata" json:"text,omitempty"`
		About string `xml:"rdf:about,attr" json:"about,omitempty"`
		Type  struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			Resource string `xml:"rdf:resource,attr" json:"resource,omitempty"`
		} `xml:"rdf:type" json:"type,omitempty"`
		HasPart struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			Ns0      string `xml:"xmlns:ns0,attr" json:"ns0,omitempty"`
			Resource string `xml:"rdf:resource,attr" json:"resource,omitempty"`
		} `xml:"ns0:hasPart" json:"haspart,omitempty"`
	} `xml:"rdf:Description" json:"description,omitempty"`
}

type Meta struct {
	XMLName  xml.Name `xml:"document-meta" json:"document-meta,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Grddl    string   `xml:"grddl,attr" json:"grddl,omitempty"`
	AttrMeta string   `xml:"meta,attr" json:"attrMeta,omitempty"`
	Dc       string   `xml:"dc,attr" json:"dc,omitempty"`
	Xlink    string   `xml:"xlink,attr" json:"xlink,omitempty"`
	Ooo      string   `xml:"ooo,attr" json:"ooo,omitempty"`
	Office   string   `xml:"office,attr" json:"office,omitempty"`
	Version  string   `xml:"version,attr" json:"version,omitempty"`
	Meta     struct {
		Text              string `xml:",chardata" json:"text,omitempty"`
		Generator         string `xml:"generator"`
		Date              string `xml:"date"`
		EditingDuration   string `xml:"editing-duration"`
		EditingCycles     string `xml:"editing-cycles"`
		DocumentStatistic struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			TableCount  string `xml:"table-count,attr" json:"table-count,omitempty"`
			CellCount   string `xml:"cell-count,attr" json:"cell-count,omitempty"`
			ObjectCount string `xml:"object-count,attr" json:"object-count,omitempty"`
		} `xml:"document-statistic" json:"document-statistic,omitempty"`
	} `xml:"meta" json:"meta,omitempty"`
}

type metaMarshal struct {
	XMLName  xml.Name `xml:"office:document-meta" json:"document-meta,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Grddl    string   `xml:"xmlns:grddl,attr" json:"grddl,omitempty"`
	AttrMeta string   `xml:"xmlns:meta,attr" json:"attrMeta,omitempty"`
	Dc       string   `xml:"xmlns:dc,attr" json:"dc,omitempty"`
	Xlink    string   `xml:"xmlns:xlink,attr" json:"xlink,omitempty"`
	Ooo      string   `xml:"xmlns:ooo,attr" json:"ooo,omitempty"`
	Office   string   `xml:"xmlns:office,attr" json:"office,omitempty"`
	Version  string   `xml:"office:version,attr" json:"version,omitempty"`
	Meta     struct {
		Text              string `xml:",chardata" json:"text,omitempty"`
		Generator         string `xml:"meta:generator"`
		Date              string `xml:"dc:date"`
		EditingDuration   string `xml:"meta:editing-duration"`
		EditingCycles     string `xml:"meta:editing-cycles"`
		DocumentStatistic struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			TableCount  string `xml:"meta:table-count,attr" json:"table-count,omitempty"`
			CellCount   string `xml:"meta:cell-count,attr" json:"cell-count,omitempty"`
			ObjectCount string `xml:"meta:object-count,attr" json:"object-count,omitempty"`
		} `xml:"meta:document-statistic" json:"document-statistic,omitempty"`
	} `xml:"office:meta" json:"meta,omitempty"`
}

type Mimetype string

type Styles struct {
	XMLName       xml.Name `xml:"document-styles" json:"document-styles,omitempty"`
	Text          string   `xml:",chardata" json:"text,omitempty"`
	Presentation  string   `xml:"presentation,attr" json:"presentation,omitempty"`
	Css3t         string   `xml:"css3t,attr" json:"css3t,omitempty"`
	Grddl         string   `xml:"grddl,attr" json:"grddl,omitempty"`
	Xhtml         string   `xml:"xhtml,attr" json:"xhtml,omitempty"`
	Dom           string   `xml:"dom,attr" json:"dom,omitempty"`
	Script        string   `xml:"script,attr" json:"script,omitempty"`
	Form          string   `xml:"form,attr" json:"form,omitempty"`
	Math          string   `xml:"math,attr" json:"math,omitempty"`
	Office        string   `xml:"office,attr" json:"office,omitempty"`
	Ooo           string   `xml:"ooo,attr" json:"ooo,omitempty"`
	Fo            string   `xml:"fo,attr" json:"fo,omitempty"`
	Ooow          string   `xml:"ooow,attr" json:"ooow,omitempty"`
	Xlink         string   `xml:"xlink,attr" json:"xlink,omitempty"`
	Drawooo       string   `xml:"drawooo,attr" json:"drawooo,omitempty"`
	Oooc          string   `xml:"oooc,attr" json:"oooc,omitempty"`
	Dc            string   `xml:"dc,attr" json:"dc,omitempty"`
	Calcext       string   `xml:"calcext,attr" json:"calcext,omitempty"`
	Style         string   `xml:"style,attr" json:"style,omitempty"`
	AttrText      string   `xml:"text,attr" json:"attrText,omitempty"`
	Of            string   `xml:"of,attr" json:"of,omitempty"`
	Tableooo      string   `xml:"tableooo,attr" json:"tableooo,omitempty"`
	Draw          string   `xml:"draw,attr" json:"draw,omitempty"`
	Dr3d          string   `xml:"dr3d,attr" json:"dr3d,omitempty"`
	Rpt           string   `xml:"rpt,attr" json:"rpt,omitempty"`
	SVG           string   `xml:"svg,attr" json:"svg,omitempty"`
	Chart         string   `xml:"chart,attr" json:"chart,omitempty"`
	Table         string   `xml:"table,attr" json:"table,omitempty"`
	Meta          string   `xml:"meta,attr" json:"meta,omitempty"`
	Loext         string   `xml:"loext,attr" json:"loext,omitempty"`
	Number        string   `xml:"number,attr" json:"number,omitempty"`
	Field         string   `xml:"field,attr" json:"field,omitempty"`
	Version       string   `xml:"version,attr" json:"version,omitempty"`
	FontFaceDecls struct {
		Text     string `xml:",chardata" json:"text,omitempty"`
		FontFace []struct {
			Text              string `xml:",chardata" json:"text,omitempty"`
			Name              string `xml:"name,attr" json:"name,omitempty"`
			FontFamily        string `xml:"font-family,attr" json:"font-family,omitempty"`
			FontFamilyGeneric string `xml:"font-family-generic,attr" json:"font-family-generic,omitempty"`
			FontPitch         string `xml:"font-pitch,attr" json:"font-pitch,omitempty"`
		} `xml:"font-face" json:"font-face,omitempty"`
	} `xml:"font-face-decls" json:"font-face-decls,omitempty"`
	Styles struct {
		Text         string `xml:",chardata" json:"text,omitempty"`
		DefaultStyle []struct {
			Text                string `xml:",chardata" json:"text,omitempty"`
			Family              string `xml:"family,attr" json:"family,omitempty"`
			ParagraphProperties struct {
				Text                       string `xml:",chardata" json:"text,omitempty"`
				TabStopDistance            string `xml:"tab-stop-distance,attr" json:"tab-stop-distance,omitempty"`
				TextAutospace              string `xml:"text-autospace,attr" json:"text-autospace,omitempty"`
				PunctuationWrap            string `xml:"punctuation-wrap,attr" json:"punctuation-wrap,omitempty"`
				LineBreak                  string `xml:"line-break,attr" json:"line-break,omitempty"`
				WritingMode                string `xml:"writing-mode,attr" json:"writing-mode,omitempty"`
				FontIndependentLineSpacing string `xml:"font-independent-line-spacing,attr" json:"font-independent-line-spacing,omitempty"`
				TabStops                   string `xml:"tab-stops"`
			} `xml:"paragraph-properties" json:"paragraph-properties,omitempty"`
			TextProperties struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				FontName           string `xml:"font-name,attr" json:"font-name,omitempty"`
				FontSize           string `xml:"font-size,attr" json:"font-size,omitempty"`
				Language           string `xml:"language,attr" json:"language,omitempty"`
				Country            string `xml:"country,attr" json:"country,omitempty"`
				FontNameAsian      string `xml:"font-name-asian,attr" json:"font-name-asian,omitempty"`
				FontSizeAsian      string `xml:"font-size-asian,attr" json:"font-size-asian,omitempty"`
				LanguageAsian      string `xml:"language-asian,attr" json:"language-asian,omitempty"`
				CountryAsian       string `xml:"country-asian,attr" json:"country-asian,omitempty"`
				FontNameComplex    string `xml:"font-name-complex,attr" json:"font-name-complex,omitempty"`
				FontSizeComplex    string `xml:"font-size-complex,attr" json:"font-size-complex,omitempty"`
				LanguageComplex    string `xml:"language-complex,attr" json:"language-complex,omitempty"`
				CountryComplex     string `xml:"country-complex,attr" json:"country-complex,omitempty"`
				UseWindowFontColor string `xml:"use-window-font-color,attr" json:"use-window-font-color,omitempty"`
				Opacity            string `xml:"opacity,attr" json:"opacity,omitempty"`
				FontFamily         string `xml:"font-family,attr" json:"font-family,omitempty"`
				FontFamilyGeneric  string `xml:"font-family-generic,attr" json:"font-family-generic,omitempty"`
				FontPitch          string `xml:"font-pitch,attr" json:"font-pitch,omitempty"`
				LetterKerning      string `xml:"letter-kerning,attr" json:"letter-kerning,omitempty"`
			} `xml:"text-properties" json:"text-properties,omitempty"`
			GraphicProperties struct {
				Text          string `xml:",chardata" json:"text,omitempty"`
				StrokeColor   string `xml:"stroke-color,attr" json:"stroke-color,omitempty"`
				FillColor     string `xml:"fill-color,attr" json:"fill-color,omitempty"`
				WrapOption    string `xml:"wrap-option,attr" json:"wrap-option,omitempty"`
				ShadowOffsetX string `xml:"shadow-offset-x,attr" json:"shadow-offset-x,omitempty"`
				ShadowOffsetY string `xml:"shadow-offset-y,attr" json:"shadow-offset-y,omitempty"`
				WritingMode   string `xml:"writing-mode,attr" json:"writing-mode,omitempty"`
			} `xml:"graphic-properties" json:"graphic-properties,omitempty"`
		} `xml:"default-style" json:"default-style,omitempty"`
		Style []struct {
			Text              string `xml:",chardata" json:"text,omitempty"`
			Name              string `xml:"name,attr" json:"name,omitempty"`
			Family            string `xml:"family,attr" json:"family,omitempty"`
			ParentStyleName   string `xml:"parent-style-name,attr" json:"parent-style-name,omitempty"`
			DisplayName       string `xml:"display-name,attr" json:"display-name,omitempty"`
			GraphicProperties struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				Stroke            string `xml:"stroke,attr" json:"stroke,omitempty"`
				MarkerStart       string `xml:"marker-start,attr" json:"marker-start,omitempty"`
				MarkerStartWidth  string `xml:"marker-start-width,attr" json:"marker-start-width,omitempty"`
				MarkerStartCenter string `xml:"marker-start-center,attr" json:"marker-start-center,omitempty"`
				Fill              string `xml:"fill,attr" json:"fill,omitempty"`
				FillColor         string `xml:"fill-color,attr" json:"fill-color,omitempty"`
				AutoGrowHeight    string `xml:"auto-grow-height,attr" json:"auto-grow-height,omitempty"`
				AutoGrowWidth     string `xml:"auto-grow-width,attr" json:"auto-grow-width,omitempty"`
				PaddingTop        string `xml:"padding-top,attr" json:"padding-top,omitempty"`
				PaddingBottom     string `xml:"padding-bottom,attr" json:"padding-bottom,omitempty"`
				PaddingLeft       string `xml:"padding-left,attr" json:"padding-left,omitempty"`
				PaddingRight      string `xml:"padding-right,attr" json:"padding-right,omitempty"`
				Shadow            string `xml:"shadow,attr" json:"shadow,omitempty"`
				ShadowOffsetX     string `xml:"shadow-offset-x,attr" json:"shadow-offset-x,omitempty"`
				ShadowOffsetY     string `xml:"shadow-offset-y,attr" json:"shadow-offset-y,omitempty"`
			} `xml:"graphic-properties" json:"graphic-properties,omitempty"`
			TextProperties struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				FontName               string `xml:"font-name,attr" json:"font-name,omitempty"`
				FontFamily             string `xml:"font-family,attr" json:"font-family,omitempty"`
				FontSize               string `xml:"font-size,attr" json:"font-size,omitempty"`
				FontNameAsian          string `xml:"font-name-asian,attr" json:"font-name-asian,omitempty"`
				FontFamilyAsian        string `xml:"font-family-asian,attr" json:"font-family-asian,omitempty"`
				FontFamilyGenericAsian string `xml:"font-family-generic-asian,attr" json:"font-family-generic-asian,omitempty"`
				FontPitchAsian         string `xml:"font-pitch-asian,attr" json:"font-pitch-asian,omitempty"`
				FontSizeAsian          string `xml:"font-size-asian,attr" json:"font-size-asian,omitempty"`
				FontNameComplex        string `xml:"font-name-complex,attr" json:"font-name-complex,omitempty"`
				FontFamilyComplex      string `xml:"font-family-complex,attr" json:"font-family-complex,omitempty"`
				FontSizeComplex        string `xml:"font-size-complex,attr" json:"font-size-complex,omitempty"`
				Color                  string `xml:"color,attr" json:"color,omitempty"`
				FontStyle              string `xml:"font-style,attr" json:"font-style,omitempty"`
				FontWeight             string `xml:"font-weight,attr" json:"font-weight,omitempty"`
				TextUnderlineStyle     string `xml:"text-underline-style,attr" json:"text-underline-style,omitempty"`
				TextUnderlineWidth     string `xml:"text-underline-width,attr" json:"text-underline-width,omitempty"`
				TextUnderlineColor     string `xml:"text-underline-color,attr" json:"text-underline-color,omitempty"`
				FontStyleAsian         string `xml:"font-style-asian,attr" json:"font-style-asian,omitempty"`
				FontWeightAsian        string `xml:"font-weight-asian,attr" json:"font-weight-asian,omitempty"`
				FontStyleComplex       string `xml:"font-style-complex,attr" json:"font-style-complex,omitempty"`
				FontWeightComplex      string `xml:"font-weight-complex,attr" json:"font-weight-complex,omitempty"`
			} `xml:"text-properties" json:"text-properties,omitempty"`
			TableCellProperties struct {
				Text            string `xml:",chardata" json:"text,omitempty"`
				RotationAlign   string `xml:"rotation-align,attr" json:"rotation-align,omitempty"`
				VerticalAlign   string `xml:"vertical-align,attr" json:"vertical-align,omitempty"`
				BackgroundColor string `xml:"background-color,attr" json:"background-color,omitempty"`
				DiagonalBlTr    string `xml:"diagonal-bl-tr,attr" json:"diagonal-bl-tr,omitempty"`
				DiagonalTlBr    string `xml:"diagonal-tl-br,attr" json:"diagonal-tl-br,omitempty"`
				Border          string `xml:"border,attr" json:"border,omitempty"`
				WrapOption      string `xml:"wrap-option,attr" json:"wrap-option,omitempty"`
				ShrinkToFit     string `xml:"shrink-to-fit,attr" json:"shrink-to-fit,omitempty"`
			} `xml:"table-cell-properties" json:"table-cell-properties,omitempty"`
		} `xml:"style" json:"style,omitempty"`
		NumberStyle []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     string `xml:"name,attr" json:"name,omitempty"`
			Volatile string `xml:"volatile,attr" json:"volatile,omitempty"`
			Language string `xml:"language,attr" json:"language,omitempty"`
			Country  string `xml:"country,attr" json:"country,omitempty"`
			Number   struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				MinIntegerDigits string `xml:"min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				DecimalPlaces    string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				Grouping         string `xml:"grouping,attr" json:"grouping,omitempty"`
			} `xml:"number" json:"number,omitempty"`
			ScientificNumber struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces      string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces   string `xml:"min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits   string `xml:"min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				MinExponentDigits  string `xml:"min-exponent-digits,attr" json:"min-exponent-digits,omitempty"`
				ExponentInterval   string `xml:"exponent-interval,attr" json:"exponent-interval,omitempty"`
				ForcedExponentSign string `xml:"forced-exponent-sign,attr" json:"forced-exponent-sign,omitempty"`
			} `xml:"scientific-number" json:"scientific-number,omitempty"`
			Text           []string `xml:"text"`
			TextProperties struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Color string `xml:"color,attr" json:"color,omitempty"`
			} `xml:"text-properties" json:"text-properties,omitempty"`
			Map struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"map" json:"map,omitempty"`
			FillCharacter string `xml:"fill-character"`
		} `xml:"number-style" json:"number-style,omitempty"`
		TimeStyle []struct {
			Chardata           string `xml:",chardata" json:"chardata,omitempty"`
			Name               string `xml:"name,attr" json:"name,omitempty"`
			TruncateOnOverflow string `xml:"truncate-on-overflow,attr" json:"truncate-on-overflow,omitempty"`
			Language           string `xml:"language,attr" json:"language,omitempty"`
			Country            string `xml:"country,attr" json:"country,omitempty"`
			Minutes            struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"style,attr" json:"style,omitempty"`
			} `xml:"minutes" json:"minutes,omitempty"`
			Text    []string `xml:"text"`
			Seconds struct {
				Text          string `xml:",chardata" json:"text,omitempty"`
				Style         string `xml:"style,attr" json:"style,omitempty"`
				DecimalPlaces string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
			} `xml:"seconds" json:"seconds,omitempty"`
			Hours string `xml:"hours"`
			AmPm  string `xml:"am-pm"`
		} `xml:"time-style" json:"time-style,omitempty"`
		DateStyle []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     string `xml:"name,attr" json:"name,omitempty"`
			Language string `xml:"language,attr" json:"language,omitempty"`
			Country  string `xml:"country,attr" json:"country,omitempty"`
			Month    struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Textual string `xml:"textual,attr" json:"textual,omitempty"`
			} `xml:"month" json:"month,omitempty"`
			Text []string `xml:"text"`
			Day  string   `xml:"day"`
			Year struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"style,attr" json:"style,omitempty"`
			} `xml:"year" json:"year,omitempty"`
			Hours   string `xml:"hours"`
			Minutes struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"style,attr" json:"style,omitempty"`
			} `xml:"minutes" json:"minutes,omitempty"`
		} `xml:"date-style" json:"date-style,omitempty"`
		CurrencyStyle []struct {
			Chardata       string `xml:",chardata" json:"chardata,omitempty"`
			Name           string `xml:"name,attr" json:"name,omitempty"`
			Volatile       string `xml:"volatile,attr" json:"volatile,omitempty"`
			Language       string `xml:"language,attr" json:"language,omitempty"`
			Country        string `xml:"country,attr" json:"country,omitempty"`
			CurrencySymbol string `xml:"currency-symbol"`
			Number         struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces    string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits string `xml:"min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				Grouping         string `xml:"grouping,attr" json:"grouping,omitempty"`
			} `xml:"number" json:"number,omitempty"`
			Text []string `xml:"text"`
			Map  struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"map" json:"map,omitempty"`
			TextProperties struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Color string `xml:"color,attr" json:"color,omitempty"`
			} `xml:"text-properties" json:"text-properties,omitempty"`
			FillCharacter string `xml:"fill-character"`
		} `xml:"currency-style" json:"currency-style,omitempty"`
		TextStyle []struct {
			Chardata    string   `xml:",chardata" json:"chardata,omitempty"`
			Name        string   `xml:"name,attr" json:"name,omitempty"`
			Language    string   `xml:"language,attr" json:"language,omitempty"`
			Country     string   `xml:"country,attr" json:"country,omitempty"`
			Text        []string `xml:"text"`
			TextContent string   `xml:"text-content"`
			Map         []struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"map" json:"map,omitempty"`
		} `xml:"text-style" json:"text-style,omitempty"`
		Marker struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			Name        string `xml:"name,attr" json:"name,omitempty"`
			DisplayName string `xml:"display-name,attr" json:"display-name,omitempty"`
			ViewBox     string `xml:"viewBox,attr" json:"viewbox,omitempty"`
			D           string `xml:"d,attr" json:"d,omitempty"`
		} `xml:"marker" json:"marker,omitempty"`
	} `xml:"styles" json:"styles,omitempty"`
	AutomaticStyles struct {
		Text        string `xml:",chardata" json:"text,omitempty"`
		NumberStyle struct {
			Text   string `xml:",chardata" json:"text,omitempty"`
			Name   string `xml:"name,attr" json:"name,omitempty"`
			Number struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces    string `xml:"decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits string `xml:"min-integer-digits,attr" json:"min-integer-digits,omitempty"`
			} `xml:"number" json:"number,omitempty"`
		} `xml:"number-style" json:"number-style,omitempty"`
		PageLayout []struct {
			Text                 string `xml:",chardata" json:"text,omitempty"`
			Name                 string `xml:"name,attr" json:"name,omitempty"`
			PageLayoutProperties struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				FirstPageNumber  string `xml:"first-page-number,attr" json:"first-page-number,omitempty"`
				WritingMode      string `xml:"writing-mode,attr" json:"writing-mode,omitempty"`
				NumFormat        string `xml:"num-format,attr" json:"num-format,omitempty"`
				PrintOrientation string `xml:"print-orientation,attr" json:"print-orientation,omitempty"`
				MarginTop        string `xml:"margin-top,attr" json:"margin-top,omitempty"`
				MarginBottom     string `xml:"margin-bottom,attr" json:"margin-bottom,omitempty"`
				MarginLeft       string `xml:"margin-left,attr" json:"margin-left,omitempty"`
				MarginRight      string `xml:"margin-right,attr" json:"margin-right,omitempty"`
				PrintPageOrder   string `xml:"print-page-order,attr" json:"print-page-order,omitempty"`
				ScaleTo          string `xml:"scale-to,attr" json:"scale-to,omitempty"`
				Print            string `xml:"print,attr" json:"print,omitempty"`
			} `xml:"page-layout-properties" json:"page-layout-properties,omitempty"`
			HeaderStyle struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				HeaderFooterProperties struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					MinHeight       string `xml:"min-height,attr" json:"min-height,omitempty"`
					MarginLeft      string `xml:"margin-left,attr" json:"margin-left,omitempty"`
					MarginRight     string `xml:"margin-right,attr" json:"margin-right,omitempty"`
					MarginBottom    string `xml:"margin-bottom,attr" json:"margin-bottom,omitempty"`
					Border          string `xml:"border,attr" json:"border,omitempty"`
					Padding         string `xml:"padding,attr" json:"padding,omitempty"`
					BackgroundColor string `xml:"background-color,attr" json:"background-color,omitempty"`
					BackgroundImage string `xml:"background-image"`
				} `xml:"header-footer-properties" json:"header-footer-properties,omitempty"`
			} `xml:"header-style" json:"header-style,omitempty"`
			FooterStyle struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				HeaderFooterProperties struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					MinHeight       string `xml:"min-height,attr" json:"min-height,omitempty"`
					MarginLeft      string `xml:"margin-left,attr" json:"margin-left,omitempty"`
					MarginRight     string `xml:"margin-right,attr" json:"margin-right,omitempty"`
					MarginTop       string `xml:"margin-top,attr" json:"margin-top,omitempty"`
					Border          string `xml:"border,attr" json:"border,omitempty"`
					Padding         string `xml:"padding,attr" json:"padding,omitempty"`
					BackgroundColor string `xml:"background-color,attr" json:"background-color,omitempty"`
					BackgroundImage string `xml:"background-image"`
				} `xml:"header-footer-properties" json:"header-footer-properties,omitempty"`
			} `xml:"footer-style" json:"footer-style,omitempty"`
		} `xml:"page-layout" json:"page-layout,omitempty"`
	} `xml:"automatic-styles" json:"automatic-styles,omitempty"`
	MasterStyles struct {
		Text       string `xml:",chardata" json:"text,omitempty"`
		MasterPage []struct {
			Text           string `xml:",chardata" json:"text,omitempty"`
			Name           string `xml:"name,attr" json:"name,omitempty"`
			PageLayoutName string `xml:"page-layout-name,attr" json:"page-layout-name,omitempty"`
			DisplayName    string `xml:"display-name,attr" json:"display-name,omitempty"`
			Header         struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
				P       struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					SheetName string `xml:"sheet-name"`
				} `xml:"p" json:"p,omitempty"`
				RegionLeft struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					P    struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						SheetName string `xml:"sheet-name"`
						S         string `xml:"s"`
						Title     string `xml:"title"`
					} `xml:"p" json:"p,omitempty"`
				} `xml:"region-left" json:"region-left,omitempty"`
				RegionRight struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					P    struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Date struct {
							Text          string `xml:",chardata" json:"text,omitempty"`
							DataStyleName string `xml:"data-style-name,attr" json:"data-style-name,omitempty"`
							DateValue     string `xml:"date-value,attr" json:"date-value,omitempty"`
						} `xml:"date" json:"date,omitempty"`
						Time struct {
							Text          string `xml:",chardata" json:"text,omitempty"`
							DataStyleName string `xml:"data-style-name,attr" json:"data-style-name,omitempty"`
							TimeValue     string `xml:"time-value,attr" json:"time-value,omitempty"`
						} `xml:"time" json:"time,omitempty"`
					} `xml:"p" json:"p,omitempty"`
				} `xml:"region-right" json:"region-right,omitempty"`
			} `xml:"header" json:"header,omitempty"`
			HeaderLeft struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
			} `xml:"header-left" json:"header-left,omitempty"`
			HeaderFirst struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
			} `xml:"header-first" json:"header-first,omitempty"`
			Footer struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
				P       struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					PageNumber string `xml:"page-number"`
					S          string `xml:"s"`
					PageCount  string `xml:"page-count"`
				} `xml:"p" json:"p,omitempty"`
			} `xml:"footer" json:"footer,omitempty"`
			FooterLeft struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
			} `xml:"footer-left" json:"footer-left,omitempty"`
			FooterFirst struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"display,attr" json:"display,omitempty"`
			} `xml:"footer-first" json:"footer-first,omitempty"`
		} `xml:"master-page" json:"master-page,omitempty"`
	} `xml:"master-styles" json:"master-styles,omitempty"`
}

type stylesMarshal struct {
	XMLName       xml.Name `xml:"office:document-styles" json:"document-styles,omitempty"`
	Text          string   `xml:",chardata" json:"text,omitempty"`
	Presentation  string   `xml:"xmlns:presentation,attr" json:"presentation,omitempty"`
	Css3t         string   `xml:"xmlns:css3t,attr" json:"css3t,omitempty"`
	Grddl         string   `xml:"xmlns:grddl,attr" json:"grddl,omitempty"`
	Xhtml         string   `xml:"xmlns:xhtml,attr" json:"xhtml,omitempty"`
	Dom           string   `xml:"xmlns:dom,attr" json:"dom,omitempty"`
	Script        string   `xml:"xmlns:script,attr" json:"script,omitempty"`
	Form          string   `xml:"xmlns:form,attr" json:"form,omitempty"`
	Math          string   `xml:"xmlns:math,attr" json:"math,omitempty"`
	Office        string   `xml:"xmlns:office,attr" json:"office,omitempty"`
	Ooo           string   `xml:"xmlns:ooo,attr" json:"ooo,omitempty"`
	Fo            string   `xml:"xmlns:fo,attr" json:"fo,omitempty"`
	Ooow          string   `xml:"xmlns:ooow,attr" json:"ooow,omitempty"`
	Xlink         string   `xml:"xmlns:xlink,attr" json:"xlink,omitempty"`
	Drawooo       string   `xml:"xmlns:drawooo,attr" json:"drawooo,omitempty"`
	Oooc          string   `xml:"xmlns:oooc,attr" json:"oooc,omitempty"`
	Dc            string   `xml:"xmlns:dc,attr" json:"dc,omitempty"`
	Calcext       string   `xml:"xmlns:calcext,attr" json:"calcext,omitempty"`
	Style         string   `xml:"xmlns:style,attr" json:"style,omitempty"`
	AttrText      string   `xml:"xmlns:text,attr" json:"attrText,omitempty"`
	Of            string   `xml:"xmlns:of,attr" json:"of,omitempty"`
	Tableooo      string   `xml:"xmlns:tableooo,attr" json:"tableooo,omitempty"`
	Draw          string   `xml:"xmlns:draw,attr" json:"draw,omitempty"`
	Dr3d          string   `xml:"xmlns:dr3d,attr" json:"dr3d,omitempty"`
	Rpt           string   `xml:"xmlns:rpt,attr" json:"rpt,omitempty"`
	SVG           string   `xml:"xmlns:svg,attr" json:"svg,omitempty"`
	Chart         string   `xml:"xmlns:chart,attr" json:"chart,omitempty"`
	Table         string   `xml:"xmlns:table,attr" json:"table,omitempty"`
	Meta          string   `xml:"xmlns:meta,attr" json:"meta,omitempty"`
	Loext         string   `xml:"xmlns:loext,attr" json:"loext,omitempty"`
	Number        string   `xml:"xmlns:number,attr" json:"number,omitempty"`
	Field         string   `xml:"xmlns:field,attr" json:"field,omitempty"`
	Version       string   `xml:"office:version,attr" json:"version,omitempty"`
	FontFaceDecls struct {
		Text     string `xml:",chardata" json:"text,omitempty"`
		FontFace []struct {
			Text              string `xml:",chardata" json:"text,omitempty"`
			Name              string `xml:"style:name,attr" json:"name,omitempty"`
			FontFamily        string `xml:"svg:font-family,attr" json:"font-family,omitempty"`
			FontFamilyGeneric string `xml:"style:font-family-generic,attr" json:"font-family-generic,omitempty"`
			FontPitch         string `xml:"style:font-pitch,attr" json:"font-pitch,omitempty"`
		} `xml:"style:font-face" json:"font-face,omitempty"`
	} `xml:"office:font-face-decls" json:"font-face-decls,omitempty"`
	Styles struct {
		Text         string `xml:",chardata" json:"text,omitempty"`
		DefaultStyle []struct {
			Text                string `xml:",chardata" json:"text,omitempty"`
			Family              string `xml:"style:family,attr" json:"family,omitempty"`
			ParagraphProperties struct {
				Text                       string `xml:",chardata" json:"text,omitempty"`
				TabStopDistance            string `xml:"style:tab-stop-distance,attr" json:"tab-stop-distance,omitempty"`
				TextAutospace              string `xml:"style:text-autospace,attr" json:"text-autospace,omitempty"`
				PunctuationWrap            string `xml:"style:punctuation-wrap,attr" json:"punctuation-wrap,omitempty"`
				LineBreak                  string `xml:"style:line-break,attr" json:"line-break,omitempty"`
				WritingMode                string `xml:"style:writing-mode,attr" json:"writing-mode,omitempty"`
				FontIndependentLineSpacing string `xml:"style:font-independent-line-spacing,attr" json:"font-independent-line-spacing,omitempty"`
				TabStops                   string `xml:"style:tab-stops"`
			} `xml:"style:paragraph-properties" json:"paragraph-properties,omitempty"`
			TextProperties struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				FontName           string `xml:"style:font-name,attr" json:"font-name,omitempty"`
				FontSize           string `xml:"fo:font-size,attr" json:"font-size,omitempty"`
				Language           string `xml:"fo:language,attr" json:"language,omitempty"`
				Country            string `xml:"fo:country,attr" json:"country,omitempty"`
				FontNameAsian      string `xml:"style:font-name-asian,attr" json:"font-name-asian,omitempty"`
				FontSizeAsian      string `xml:"style:font-size-asian,attr" json:"font-size-asian,omitempty"`
				LanguageAsian      string `xml:"style:language-asian,attr" json:"language-asian,omitempty"`
				CountryAsian       string `xml:"style:country-asian,attr" json:"country-asian,omitempty"`
				FontNameComplex    string `xml:"style:font-name-complex,attr" json:"font-name-complex,omitempty"`
				FontSizeComplex    string `xml:"style:font-size-complex,attr" json:"font-size-complex,omitempty"`
				LanguageComplex    string `xml:"style:language-complex,attr" json:"language-complex,omitempty"`
				CountryComplex     string `xml:"style:country-complex,attr" json:"country-complex,omitempty"`
				UseWindowFontColor string `xml:"style:use-window-font-color,attr" json:"use-window-font-color,omitempty"`
				Opacity            string `xml:"loext:opacity,attr" json:"opacity,omitempty"`
				FontFamily         string `xml:"fo:font-family,attr" json:"font-family,omitempty"`
				FontFamilyGeneric  string `xml:"style:font-family-generic,attr" json:"font-family-generic,omitempty"`
				FontPitch          string `xml:"style:font-pitch,attr" json:"font-pitch,omitempty"`
				LetterKerning      string `xml:"style:letter-kerning,attr" json:"letter-kerning,omitempty"`
			} `xml:"style:text-properties" json:"text-properties,omitempty"`
			GraphicProperties struct {
				Text          string `xml:",chardata" json:"text,omitempty"`
				StrokeColor   string `xml:"svg:stroke-color,attr" json:"stroke-color,omitempty"`
				FillColor     string `xml:"draw:fill-color,attr" json:"fill-color,omitempty"`
				WrapOption    string `xml:"fo:wrap-option,attr" json:"wrap-option,omitempty"`
				ShadowOffsetX string `xml:"draw:shadow-offset-x,attr" json:"shadow-offset-x,omitempty"`
				ShadowOffsetY string `xml:"draw:shadow-offset-y,attr" json:"shadow-offset-y,omitempty"`
				WritingMode   string `xml:"style:writing-mode,attr" json:"writing-mode,omitempty"`
			} `xml:"style:graphic-properties" json:"graphic-properties,omitempty"`
		} `xml:"style:default-style" json:"default-style,omitempty"`
		Style []struct {
			Text              string `xml:",chardata" json:"text,omitempty"`
			Name              string `xml:"style:name,attr" json:"name,omitempty"`
			Family            string `xml:"style:family,attr" json:"family,omitempty"`
			ParentStyleName   string `xml:"style:parent-style-name,attr" json:"parent-style-name,omitempty"`
			DisplayName       string `xml:"style:display-name,attr" json:"display-name,omitempty"`
			GraphicProperties struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				Stroke            string `xml:"draw:stroke,attr" json:"stroke,omitempty"`
				MarkerStart       string `xml:"draw:marker-start,attr" json:"marker-start,omitempty"`
				MarkerStartWidth  string `xml:"draw:marker-start-width,attr" json:"marker-start-width,omitempty"`
				MarkerStartCenter string `xml:"draw:marker-start-center,attr" json:"marker-start-center,omitempty"`
				Fill              string `xml:"draw:fill,attr" json:"fill,omitempty"`
				FillColor         string `xml:"draw:fill-color,attr" json:"fill-color,omitempty"`
				AutoGrowHeight    string `xml:"draw:auto-grow-height,attr" json:"auto-grow-height,omitempty"`
				AutoGrowWidth     string `xml:"draw:auto-grow-width,attr" json:"auto-grow-width,omitempty"`
				PaddingTop        string `xml:"fo:padding-top,attr" json:"padding-top,omitempty"`
				PaddingBottom     string `xml:"fo:padding-bottom,attr" json:"padding-bottom,omitempty"`
				PaddingLeft       string `xml:"fo:padding-left,attr" json:"padding-left,omitempty"`
				PaddingRight      string `xml:"fo:padding-right,attr" json:"padding-right,omitempty"`
				Shadow            string `xml:"draw:shadow,attr" json:"shadow,omitempty"`
				ShadowOffsetX     string `xml:"draw:shadow-offset-x,attr" json:"shadow-offset-x,omitempty"`
				ShadowOffsetY     string `xml:"draw:shadow-offset-y,attr" json:"shadow-offset-y,omitempty"`
			} `xml:"style:graphic-properties" json:"graphic-properties,omitempty"`
			TextProperties struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				FontName               string `xml:"style:font-name,attr" json:"font-name,omitempty"`
				FontFamily             string `xml:"fo:font-family,attr" json:"font-family,omitempty"`
				FontSize               string `xml:"fo:font-size,attr" json:"font-size,omitempty"`
				FontNameAsian          string `xml:"style:font-name-asian,attr" json:"font-name-asian,omitempty"`
				FontFamilyAsian        string `xml:"style:font-family-asian,attr" json:"font-family-asian,omitempty"`
				FontFamilyGenericAsian string `xml:"style:font-family-generic-asian,attr" json:"font-family-generic-asian,omitempty"`
				FontPitchAsian         string `xml:"style:font-pitch-asian,attr" json:"font-pitch-asian,omitempty"`
				FontSizeAsian          string `xml:"style:font-size-asian,attr" json:"font-size-asian,omitempty"`
				FontNameComplex        string `xml:"style:font-name-complex,attr" json:"font-name-complex,omitempty"`
				FontFamilyComplex      string `xml:"style:font-family-complex,attr" json:"font-family-complex,omitempty"`
				FontSizeComplex        string `xml:"style:font-size-complex,attr" json:"font-size-complex,omitempty"`
				Color                  string `xml:"fo:color,attr" json:"color,omitempty"`
				FontStyle              string `xml:"fo:font-style,attr" json:"font-style,omitempty"`
				FontWeight             string `xml:"fo:font-weight,attr" json:"font-weight,omitempty"`
				TextUnderlineStyle     string `xml:"style:text-underline-style,attr" json:"text-underline-style,omitempty"`
				TextUnderlineWidth     string `xml:"style:text-underline-width,attr" json:"text-underline-width,omitempty"`
				TextUnderlineColor     string `xml:"style:text-underline-color,attr" json:"text-underline-color,omitempty"`
				FontStyleAsian         string `xml:"style:font-style-asian,attr" json:"font-style-asian,omitempty"`
				FontWeightAsian        string `xml:"style:font-weight-asian,attr" json:"font-weight-asian,omitempty"`
				FontStyleComplex       string `xml:"style:font-style-complex,attr" json:"font-style-complex,omitempty"`
				FontWeightComplex      string `xml:"style:font-weight-complex,attr" json:"font-weight-complex,omitempty"`
			} `xml:"style:text-properties" json:"text-properties,omitempty"`
			TableCellProperties struct {
				Text            string `xml:",chardata" json:"text,omitempty"`
				RotationAlign   string `xml:"style:rotation-align,attr" json:"rotation-align,omitempty"`
				VerticalAlign   string `xml:"style:vertical-align,attr" json:"vertical-align,omitempty"`
				BackgroundColor string `xml:"fo:background-color,attr" json:"background-color,omitempty"`
				DiagonalBlTr    string `xml:"style:diagonal-bl-tr,attr" json:"diagonal-bl-tr,omitempty"`
				DiagonalTlBr    string `xml:"style:diagonal-tl-br,attr" json:"diagonal-tl-br,omitempty"`
				Border          string `xml:"fo:border,attr" json:"border,omitempty"`
				WrapOption      string `xml:"fo:wrap-option,attr" json:"wrap-option,omitempty"`
				ShrinkToFit     string `xml:"style:shrink-to-fit,attr" json:"shrink-to-fit,omitempty"`
			} `xml:"style:table-cell-properties" json:"table-cell-properties,omitempty"`
		} `xml:"style:style" json:"style,omitempty"`
		NumberStyle []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     string `xml:"style:name,attr" json:"name,omitempty"`
			Volatile string `xml:"style:volatile,attr" json:"volatile,omitempty"`
			Language string `xml:"number:language,attr" json:"language,omitempty"`
			Country  string `xml:"number:country,attr" json:"country,omitempty"`
			Number   struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				MinIntegerDigits string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				DecimalPlaces    string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				Grouping         string `xml:"number:grouping,attr" json:"grouping,omitempty"`
			} `xml:"number:number" json:"number,omitempty"`
			ScientificNumber struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces      string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces   string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits   string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				MinExponentDigits  string `xml:"number:min-exponent-digits,attr" json:"min-exponent-digits,omitempty"`
				ExponentInterval   string `xml:"number:exponent-interval,attr" json:"exponent-interval,omitempty"`
				ForcedExponentSign string `xml:"number:forced-exponent-sign,attr" json:"forced-exponent-sign,omitempty"`
			} `xml:"number:scientific-number" json:"scientific-number,omitempty"`
			Text           []string `xml:"number:text"`
			TextProperties struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Color string `xml:"fo:color,attr" json:"color,omitempty"`
			} `xml:"style:text-properties" json:"text-properties,omitempty"`
			Map struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"style:condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"style:apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"style:map" json:"map,omitempty"`
			FillCharacter string `xml:"number:fill-character"`
		} `xml:"number:number-style" json:"number-style,omitempty"`
		TimeStyle []struct {
			Chardata           string `xml:",chardata" json:"chardata,omitempty"`
			Name               string `xml:"style:name,attr" json:"name,omitempty"`
			TruncateOnOverflow string `xml:"number:truncate-on-overflow,attr" json:"truncate-on-overflow,omitempty"`
			Language           string `xml:"number:language,attr" json:"language,omitempty"`
			Country            string `xml:"number:country,attr" json:"country,omitempty"`
			Minutes            struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"number:style,attr" json:"style,omitempty"`
			} `xml:"number:minutes" json:"minutes,omitempty"`
			Text    []string `xml:"number:text"`
			Seconds struct {
				Text          string `xml:",chardata" json:"text,omitempty"`
				Style         string `xml:"number:style,attr" json:"style,omitempty"`
				DecimalPlaces string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
			} `xml:"number:seconds" json:"seconds,omitempty"`
			Hours string `xml:"number:hours"`
			AmPm  string `xml:"number:am-pm"`
		} `xml:"number:time-style" json:"time-style,omitempty"`
		DateStyle []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     string `xml:"style:name,attr" json:"name,omitempty"`
			Language string `xml:"number:language,attr" json:"language,omitempty"`
			Country  string `xml:"number:country,attr" json:"country,omitempty"`
			Month    struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Textual string `xml:"number:textual,attr" json:"textual,omitempty"`
			} `xml:"number:month" json:"month,omitempty"`
			Text []string `xml:"number:text"`
			Day  string   `xml:"number:day"`
			Year struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"number:style,attr" json:"style,omitempty"`
			} `xml:"number:year" json:"year,omitempty"`
			Hours   string `xml:"number:hours"`
			Minutes struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Style string `xml:"number:style,attr" json:"style,omitempty"`
			} `xml:"number:minutes" json:"minutes,omitempty"`
		} `xml:"number:date-style" json:"date-style,omitempty"`
		CurrencyStyle []struct {
			Chardata       string `xml:",chardata" json:"chardata,omitempty"`
			Name           string `xml:"style:name,attr" json:"name,omitempty"`
			Volatile       string `xml:"style:volatile,attr" json:"volatile,omitempty"`
			Language       string `xml:"number:language,attr" json:"language,omitempty"`
			Country        string `xml:"number:country,attr" json:"country,omitempty"`
			CurrencySymbol string `xml:"number:currency-symbol"`
			Number         struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces    string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
				Grouping         string `xml:"number:grouping,attr" json:"grouping,omitempty"`
			} `xml:"number:number" json:"number,omitempty"`
			Text []string `xml:"number:text"`
			Map  struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"style:condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"style:apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"style:map" json:"map,omitempty"`
			TextProperties struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Color string `xml:"fo:color,attr" json:"color,omitempty"`
			} `xml:"style:text-properties" json:"text-properties,omitempty"`
			FillCharacter string `xml:"number:fill-character"`
		} `xml:"number:currency-style" json:"currency-style,omitempty"`
		TextStyle []struct {
			Chardata    string   `xml:",chardata" json:"chardata,omitempty"`
			Name        string   `xml:"style:name,attr" json:"name,omitempty"`
			Language    string   `xml:"number:language,attr" json:"language,omitempty"`
			Country     string   `xml:"number:country,attr" json:"country,omitempty"`
			Text        []string `xml:"number:text"`
			TextContent string   `xml:"number:text-content"`
			Map         []struct {
				Text           string `xml:",chardata" json:"text,omitempty"`
				Condition      string `xml:"style:condition,attr" json:"condition,omitempty"`
				ApplyStyleName string `xml:"style:apply-style-name,attr" json:"apply-style-name,omitempty"`
			} `xml:"style:map" json:"map,omitempty"`
		} `xml:"number:text-style" json:"text-style,omitempty"`
		Marker struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			Name        string `xml:"draw:name,attr" json:"name,omitempty"`
			DisplayName string `xml:"draw:display-name,attr" json:"display-name,omitempty"`
			ViewBox     string `xml:"svg:viewBox,attr" json:"viewbox,omitempty"`
			D           string `xml:"svg:d,attr" json:"d,omitempty"`
		} `xml:"draw:marker" json:"marker,omitempty"`
	} `xml:"office:styles" json:"styles,omitempty"`
	AutomaticStyles struct {
		Text        string `xml:",chardata" json:"text,omitempty"`
		NumberStyle struct {
			Text   string `xml:",chardata" json:"text,omitempty"`
			Name   string `xml:"style:name,attr" json:"name,omitempty"`
			Number struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				DecimalPlaces    string `xml:"number:decimal-places,attr" json:"decimal-places,omitempty"`
				MinDecimalPlaces string `xml:"number:min-decimal-places,attr" json:"min-decimal-places,omitempty"`
				MinIntegerDigits string `xml:"number:min-integer-digits,attr" json:"min-integer-digits,omitempty"`
			} `xml:"number:number" json:"number,omitempty"`
		} `xml:"number:number-style" json:"number-style,omitempty"`
		PageLayout []struct {
			Text                 string `xml:",chardata" json:"text,omitempty"`
			Name                 string `xml:"style:name,attr" json:"name,omitempty"`
			PageLayoutProperties struct {
				Text             string `xml:",chardata" json:"text,omitempty"`
				FirstPageNumber  string `xml:"style:first-page-number,attr" json:"first-page-number,omitempty"`
				WritingMode      string `xml:"style:writing-mode,attr" json:"writing-mode,omitempty"`
				NumFormat        string `xml:"style:num-format,attr" json:"num-format,omitempty"`
				PrintOrientation string `xml:"style:print-orientation,attr" json:"print-orientation,omitempty"`
				MarginTop        string `xml:"fo:margin-top,attr" json:"margin-top,omitempty"`
				MarginBottom     string `xml:"fo:margin-bottom,attr" json:"margin-bottom,omitempty"`
				MarginLeft       string `xml:"fo:margin-left,attr" json:"margin-left,omitempty"`
				MarginRight      string `xml:"fo:margin-right,attr" json:"margin-right,omitempty"`
				PrintPageOrder   string `xml:"style:print-page-order,attr" json:"print-page-order,omitempty"`
				ScaleTo          string `xml:"style:scale-to,attr" json:"scale-to,omitempty"`
				Print            string `xml:"style:print,attr" json:"print,omitempty"`
			} `xml:"style:page-layout-properties" json:"page-layout-properties,omitempty"`
			HeaderStyle struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				HeaderFooterProperties struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					MinHeight       string `xml:"fo:min-height,attr" json:"min-height,omitempty"`
					MarginLeft      string `xml:"fo:margin-left,attr" json:"margin-left,omitempty"`
					MarginRight     string `xml:"fo:margin-right,attr" json:"margin-right,omitempty"`
					MarginBottom    string `xml:"fo:margin-bottom,attr" json:"margin-bottom,omitempty"`
					Border          string `xml:"fo:border,attr" json:"border,omitempty"`
					Padding         string `xml:"fo:padding,attr" json:"padding,omitempty"`
					BackgroundColor string `xml:"fo:background-color,attr" json:"background-color,omitempty"`
					BackgroundImage string `xml:"style:background-image"`
				} `xml:"style:header-footer-properties" json:"header-footer-properties,omitempty"`
			} `xml:"style:header-style" json:"header-style,omitempty"`
			FooterStyle struct {
				Text                   string `xml:",chardata" json:"text,omitempty"`
				HeaderFooterProperties struct {
					Text            string `xml:",chardata" json:"text,omitempty"`
					MinHeight       string `xml:"fo:min-height,attr" json:"min-height,omitempty"`
					MarginLeft      string `xml:"fo:margin-left,attr" json:"margin-left,omitempty"`
					MarginRight     string `xml:"fo:margin-right,attr" json:"margin-right,omitempty"`
					MarginTop       string `xml:"fo:margin-top,attr" json:"margin-top,omitempty"`
					Border          string `xml:"fo:border,attr" json:"border,omitempty"`
					Padding         string `xml:"fo:padding,attr" json:"padding,omitempty"`
					BackgroundColor string `xml:"fo:background-color,attr" json:"background-color,omitempty"`
					BackgroundImage string `xml:"style:background-image"`
				} `xml:"style:header-footer-properties" json:"header-footer-properties,omitempty"`
			} `xml:"style:footer-style" json:"footer-style,omitempty"`
		} `xml:"style:page-layout" json:"page-layout,omitempty"`
	} `xml:"office:automatic-styles" json:"automatic-styles,omitempty"`
	MasterStyles struct {
		Text       string `xml:",chardata" json:"text,omitempty"`
		MasterPage []struct {
			Text           string `xml:",chardata" json:"text,omitempty"`
			Name           string `xml:"style:name,attr" json:"name,omitempty"`
			PageLayoutName string `xml:"style:page-layout-name,attr" json:"page-layout-name,omitempty"`
			DisplayName    string `xml:"style:display-name,attr" json:"display-name,omitempty"`
			Header         struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
				P       struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					SheetName string `xml:"text:sheet-name"`
				} `xml:"text:p" json:"p,omitempty"`
				RegionLeft struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					P    struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						SheetName string `xml:"text:sheet-name"`
						S         string `xml:"text:s"`
						Title     string `xml:"text:title"`
					} `xml:"text:p" json:"p,omitempty"`
				} `xml:"style:region-left" json:"region-left,omitempty"`
				RegionRight struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					P    struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Date struct {
							Text          string `xml:",chardata" json:"text,omitempty"`
							DataStyleName string `xml:"style:data-style-name,attr" json:"data-style-name,omitempty"`
							DateValue     string `xml:"text:date-value,attr" json:"date-value,omitempty"`
						} `xml:"text:date" json:"date,omitempty"`
						Time struct {
							Text          string `xml:",chardata" json:"text,omitempty"`
							DataStyleName string `xml:"style:data-style-name,attr" json:"data-style-name,omitempty"`
							TimeValue     string `xml:"text:time-value,attr" json:"time-value,omitempty"`
						} `xml:"text:time" json:"time,omitempty"`
					} `xml:"text:p" json:"p,omitempty"`
				} `xml:"style:region-right" json:"region-right,omitempty"`
			} `xml:"style:header" json:"header,omitempty"`
			HeaderLeft struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
			} `xml:"style:header-left" json:"header-left,omitempty"`
			HeaderFirst struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
			} `xml:"style:header-first" json:"header-first,omitempty"`
			Footer struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
				P       struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					PageNumber string `xml:"text:page-number"`
					S          string `xml:"text:s"`
					PageCount  string `xml:"text:page-count"`
				} `xml:"text:p" json:"p,omitempty"`
			} `xml:"style:footer" json:"footer,omitempty"`
			FooterLeft struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
			} `xml:"style:footer-left" json:"footer-left,omitempty"`
			FooterFirst struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Display string `xml:"style:display,attr" json:"display,omitempty"`
			} `xml:"style:footer-first" json:"footer-first,omitempty"`
		} `xml:"style:master-page" json:"master-page,omitempty"`
	} `xml:"office:master-styles" json:"master-styles,omitempty"`
}

type Settings struct {
	XMLName  xml.Name `xml:"document-settings" json:"document-settings,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Config   string   `xml:"config,attr" json:"config,omitempty"`
	Xlink    string   `xml:"xlink,attr" json:"xlink,omitempty"`
	Ooo      string   `xml:"ooo,attr" json:"ooo,omitempty"`
	Office   string   `xml:"office,attr" json:"office,omitempty"`
	Version  string   `xml:"version,attr" json:"version,omitempty"`
	Settings struct {
		Text          string `xml:",chardata" json:"text,omitempty"`
		ConfigItemSet []struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			Name       string `xml:"name,attr" json:"name,omitempty"`
			ConfigItem []struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Name string `xml:"name,attr" json:"name,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"config-item" json:"config-item,omitempty"`
			ConfigItemMapIndexed struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				Name               string `xml:"name,attr" json:"name,omitempty"`
				ConfigItemMapEntry struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					ConfigItem []struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name string `xml:"name,attr" json:"name,omitempty"`
						Type string `xml:"type,attr" json:"type,omitempty"`
					} `xml:"config-item" json:"config-item,omitempty"`
					ConfigItemMapNamed struct {
						Text               string `xml:",chardata" json:"text,omitempty"`
						Name               string `xml:"name,attr" json:"name,omitempty"`
						ConfigItemMapEntry []struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							Name       string `xml:"name,attr" json:"name,omitempty"`
							ConfigItem []struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Name string `xml:"name,attr" json:"name,omitempty"`
								Type string `xml:"type,attr" json:"type,omitempty"`
							} `xml:"config-item" json:"config-item,omitempty"`
						} `xml:"config-item-map-entry" json:"config-item-map-entry,omitempty"`
					} `xml:"config-item-map-named" json:"config-item-map-named,omitempty"`
				} `xml:"config-item-map-entry" json:"config-item-map-entry,omitempty"`
			} `xml:"config-item-map-indexed" json:"config-item-map-indexed,omitempty"`
			ConfigItemMapNamed struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				Name               string `xml:"name,attr" json:"name,omitempty"`
				ConfigItemMapEntry []struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Name       string `xml:"name,attr" json:"name,omitempty"`
					ConfigItem struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name string `xml:"name,attr" json:"name,omitempty"`
						Type string `xml:"type,attr" json:"type,omitempty"`
					} `xml:"config-item" json:"config-item,omitempty"`
				} `xml:"config-item-map-entry" json:"config-item-map-entry,omitempty"`
			} `xml:"config-item-map-named" json:"config-item-map-named,omitempty"`
		} `xml:"config-item-set" json:"config-item-set,omitempty"`
	} `xml:"settings" json:"settings,omitempty"`
}

type settingsMarshal struct {
	XMLName  xml.Name `xml:"office:document-settings" json:"document-settings,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Config   string   `xml:"xmlns:config,attr" json:"config,omitempty"`
	Xlink    string   `xml:"xmlns:xlink,attr" json:"xlink,omitempty"`
	Ooo      string   `xml:"xmlns:ooo,attr" json:"ooo,omitempty"`
	Office   string   `xml:"xmlns:office,attr" json:"office,omitempty"`
	Version  string   `xml:"office:version,attr" json:"version,omitempty"`
	Settings struct {
		Text          string `xml:",chardata" json:"text,omitempty"`
		ConfigItemSet []struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			Name       string `xml:"config:name,attr" json:"name,omitempty"`
			ConfigItem []struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Name string `xml:"config:name,attr" json:"name,omitempty"`
				Type string `xml:"config:type,attr" json:"type,omitempty"`
			} `xml:"config:config-item" json:"config-item,omitempty"`
			ConfigItemMapIndexed struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				Name               string `xml:"config:name,attr" json:"name,omitempty"`
				ConfigItemMapEntry struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					ConfigItem []struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name string `xml:"config:name,attr" json:"name,omitempty"`
						Type string `xml:"config:type,attr" json:"type,omitempty"`
					} `xml:"config:config-item" json:"config-item,omitempty"`
					ConfigItemMapNamed struct {
						Text               string `xml:",chardata" json:"text,omitempty"`
						Name               string `xml:"config:name,attr" json:"name,omitempty"`
						ConfigItemMapEntry []struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							Name       string `xml:"config:name,attr" json:"name,omitempty"`
							ConfigItem []struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Name string `xml:"config:name,attr" json:"name,omitempty"`
								Type string `xml:"config:type,attr" json:"type,omitempty"`
							} `xml:"config:config-item" json:"config-item,omitempty"`
						} `xml:"config:config-item-map-entry" json:"config-item-map-entry,omitempty"`
					} `xml:"config:config-item-map-named" json:"config-item-map-named,omitempty"`
				} `xml:"config:config-item-map-entry" json:"config-item-map-entry,omitempty"`
			} `xml:"config:config-item-map-indexed" json:"config-item-map-indexed,omitempty"`
			ConfigItemMapNamed struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				Name               string `xml:"config:name,attr" json:"name,omitempty"`
				ConfigItemMapEntry []struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					Name       string `xml:"config:name,attr" json:"name,omitempty"`
					ConfigItem struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name string `xml:"config:name,attr" json:"name,omitempty"`
						Type string `xml:"config:type,attr" json:"type,omitempty"`
					} `xml:"config:config-item" json:"config-item,omitempty"`
				} `xml:"config:config-item-map-entry" json:"config-item-map-entry,omitempty"`
			} `xml:"config:config-item-map-named" json:"config-item-map-named,omitempty"`
		} `xml:"config:config-item-set" json:"config-item-set,omitempty"`
	} `xml:"office:settings" json:"settings,omitempty"`
}
