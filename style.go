package ods

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

// styleRegistry deduplicates styles within a document. It maps style
// definitions to generated names and ensures that identical styles share
// the same name. It also provides reverse lookups from names to definitions.
type styleRegistry struct {
	cell map[CellStyleDef]string
	row  map[RowStyleDef]string
	col  map[ColStyleDef]string
	next int
	doc  *Document
}

// CellStyleDef describes the complete set of visual properties for a cell style.
// Two CellStyleDef values that differ only in case of color strings (e.g. "f00"
// vs "#f00") will be normalized by the style builder before lookup.
type CellStyleDef struct {
	Bold            bool
	Italic          bool
	Color           string
	BackgroundColor string
	FontSize        string
	FontFamily      string
	Border          string
	HAlign          string
	VAlign          string
}

// RowStyleDef describes the visual properties for a row style, including
// row height and text formatting.
type RowStyleDef struct {
	Height string
	Bold   bool
	Italic bool
}

// ColStyleDef describes the visual properties for a column style, including
// column width and text formatting.
type ColStyleDef struct {
	Width  string
	Italic bool
	Bold   bool
}

// newStyleRegistry creates a new empty style registry.
func newStyleRegistry() *styleRegistry {
	return &styleRegistry{
		cell: map[CellStyleDef]string{},
		row:  map[RowStyleDef]string{},
		col:  map[ColStyleDef]string{},
	}
}

// attach binds the registry to a document so it can read and write styles.
func (r *styleRegistry) attach(doc *Document) {
	r.doc = doc
}

// cellStyle returns the style name for the given cell style definition,
// creating a new style if one does not already exist.
func (r *styleRegistry) cellStyle(def CellStyleDef) string {
	if name, ok := r.cell[def]; ok {
		return name
	}
	name := r.newName("go-cs", fmt.Sprintf("%+v", def))
	r.cell[def] = name
	r.doc.content.AutomaticStyles.Styles = append(r.doc.content.AutomaticStyles.Styles, Style{
		Name:   name,
		Family: styleFamilyTableCell,
		TableCellProperties: &TableCellProperties{
			BackgroundColor: def.BackgroundColor,
			Border:          def.Border,
			VerticalAlign:   def.VAlign,
		},
		ParagraphProperties: &ParagraphProperties{TextAlign: def.HAlign},
		TextProperties:      textProperties(def.Bold, def.Italic, def.Color, def.FontSize, def.FontFamily),
	})
	return name
}

// rowStyle returns the style name for the given row style definition,
// creating a new style if one does not already exist.
func (r *styleRegistry) rowStyle(def RowStyleDef) string {
	if name, ok := r.row[def]; ok {
		return name
	}
	name := r.newName("go-rs", fmt.Sprintf("%+v", def))
	r.row[def] = name
	r.doc.content.AutomaticStyles.Styles = append(r.doc.content.AutomaticStyles.Styles, Style{
		Name:               name,
		Family:             styleFamilyTableRow,
		TableRowProperties: &TableRowProperties{RowHeight: def.Height},
		TextProperties:     textProperties(def.Bold, def.Italic, "", "", ""),
	})
	return name
}

// colStyle returns the style name for the given column style definition,
// creating a new style if one does not already exist.
func (r *styleRegistry) colStyle(def ColStyleDef) string {
	if name, ok := r.col[def]; ok {
		return name
	}
	name := r.newName("go-cols", fmt.Sprintf("%+v", def))
	r.col[def] = name
	r.doc.content.AutomaticStyles.Styles = append(r.doc.content.AutomaticStyles.Styles, Style{
		Name:                  name,
		Family:                styleFamilyTableColumn,
		TableColumnProperties: &TableColumnProperties{ColumnWidth: def.Width},
		TextProperties:        textProperties(def.Bold, def.Italic, "", "", ""),
	})
	return name
}

// cellStyleDef reconstructs a CellStyleDef from a named style in the
// document's automatic styles. Returns an empty CellStyleDef if not found.
func (r *styleRegistry) cellStyleDef(name string) CellStyleDef {
	style := r.style(name)
	if style == nil {
		return CellStyleDef{}
	}

	def := CellStyleDef{}
	if style.TableCellProperties != nil {
		def.BackgroundColor = style.TableCellProperties.BackgroundColor
		def.Border = style.TableCellProperties.Border
		def.VAlign = style.TableCellProperties.VerticalAlign
	}

	if style.ParagraphProperties != nil {
		def.HAlign = style.ParagraphProperties.TextAlign
	}

	if style.TextProperties != nil {
		def.Bold = isBold(style.TextProperties)
		def.Italic = isItalic(style.TextProperties)
		def.Color = style.TextProperties.Color
		def.FontSize = style.TextProperties.FontSize
		def.FontFamily = style.TextProperties.FontFamily
	}

	return def
}

// rowStyleDef reconstructs a RowStyleDef from a named style in the
// document's automatic styles. Returns an empty RowStyleDef if not found.
func (r *styleRegistry) rowStyleDef(name string) RowStyleDef {
	style := r.style(name)
	if style == nil {
		return RowStyleDef{}
	}

	def := RowStyleDef{}
	if style.TableRowProperties != nil {
		def.Height = style.TableRowProperties.RowHeight
	}

	if style.TextProperties != nil {
		def.Bold = isBold(style.TextProperties)
		def.Italic = isItalic(style.TextProperties)
	}

	return def
}

// colStyleDef reconstructs a ColStyleDef from a named style in the
// document's automatic styles. Returns an empty ColStyleDef if not found.
func (r *styleRegistry) colStyleDef(name string) ColStyleDef {
	style := r.style(name)
	if style == nil {
		return ColStyleDef{}
	}

	def := ColStyleDef{}
	if style.TableColumnProperties != nil {
		def.Width = style.TableColumnProperties.ColumnWidth
	}

	if style.TextProperties != nil {
		def.Bold = isBold(style.TextProperties)
		def.Italic = isItalic(style.TextProperties)
	}

	return def
}

// style returns a pointer to the Style with the given name in the document's
// automatic styles, or nil if not found.
func (r *styleRegistry) style(name string) *Style {
	if r == nil || r.doc == nil || name == "" {
		return nil
	}

	for i := range r.doc.content.AutomaticStyles.Styles {
		if r.doc.content.AutomaticStyles.Styles[i].Name == name {
			return &r.doc.content.AutomaticStyles.Styles[i]
		}
	}

	return nil
}

// isBold returns whether the TextProperties is bold.
func isBold(props *TextProperties) bool {
	return props.FontWeight == fontWeightBold ||
		props.FontWeightAsian == fontWeightBold ||
		props.FontWeightComplex == fontWeightBold
}

// isItalic returns whether the TextProperties is italic.
func isItalic(props *TextProperties) bool {
	return props.FontStyle == fontStyleItalic ||
		props.FontStyleAsian == fontStyleItalic ||
		props.FontStyleComplex == fontStyleItalic
}

// newName generates a unique style name by hashing the definition key
// with SHA-1 and appending a short hex prefix.
func (r *styleRegistry) newName(prefix, key string) string {
	sum := sha1.Sum([]byte(key))
	short := hex.EncodeToString(sum[:])[:8]
	name := prefix + "-" + short
	for styleNameExists(r.doc.content.AutomaticStyles.Styles, name) {
		r.next++
		name = fmt.Sprintf("%s-%s-%d", prefix, short, r.next)
	}
	return name
}

// styleNameExists returns whether a style with the specified name exists.
func styleNameExists(styles []Style, name string) bool {
	for _, style := range styles {
		if style.Name == name {
			return true
		}
	}
	return false
}

// textProperties creates a TextProperties struct from the given formatting
// parameters. Returns nil if no properties are set.
func textProperties(bold, italic bool, color, fontSize, fontFamily string) *TextProperties {
	props := &TextProperties{
		Color:      color,
		FontSize:   fontSize,
		FontFamily: fontFamily,
	}

	if bold {
		props.FontWeight = fontWeightBold
		props.FontWeightAsian = fontWeightBold
		props.FontWeightComplex = fontWeightBold
	}

	if italic {
		props.FontStyle = fontStyleItalic
		props.FontStyleAsian = fontStyleItalic
		props.FontStyleComplex = fontStyleItalic
	}

	if *props == (TextProperties{}) {
		return nil
	}

	return props
}

// normalizeColor ensures the color string has a leading "#" prefix.
func normalizeColor(color string) string {
	if color == "" || strings.HasPrefix(color, "#") {
		return color
	}
	return "#" + color
}
