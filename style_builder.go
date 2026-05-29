package ods

// CellStyle is a builder for configuring a cell's visual appearance.
// Call Apply to persist the accumulated changes to the cell.
type CellStyle struct {
	cell    *Cell
	def     CellStyleDef
	changed bool
	clear   bool
}

// RowStyle is a builder for configuring a row's visual appearance.
// Call Apply to persist the accumulated changes to the row.
type RowStyle struct {
	row     *Row
	def     RowStyleDef
	changed bool
	clear   bool
}

// ColStyle is a builder for configuring a column's visual appearance.
// Call Apply to persist the accumulated changes to the column.
type ColStyle struct {
	col     *Col
	def     ColStyleDef
	changed bool
	clear   bool
}

// Bold toggles bold text on the cell.
func (s *CellStyle) Bold(enabled bool) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.Bold = enabled
	s.changed = true
	return s
}

// Italic toggles italic text on the cell.
func (s *CellStyle) Italic(enabled bool) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.Italic = enabled
	s.changed = true
	return s
}

// Color sets the text color of the cell.
func (s *CellStyle) Color(color string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.Color = normalizeColor(color)
	s.changed = true
	return s
}

// BackgroundColor sets the background color of the cell.
func (s *CellStyle) BackgroundColor(color string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.BackgroundColor = normalizeColor(color)
	s.changed = true
	return s
}

// FontSize sets the font size of the cell (e.g. "12pt", "1.5em").
func (s *CellStyle) FontSize(size string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.FontSize = size
	s.changed = true
	return s
}

// FontFamily sets the font family of the cell (e.g. "Arial", "serif").
func (s *CellStyle) FontFamily(family string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.FontFamily = family
	s.changed = true
	return s
}

// Border sets the cell border using CSS-like syntax (e.g. "1pt solid #000").
func (s *CellStyle) Border(border string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.Border = border
	s.changed = true
	return s
}

// HAlign sets the horizontal text alignment of the cell.
// Accepts: "left", "center", "right", and "justify".
func (s *CellStyle) HAlign(align string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.HAlign = align
	s.changed = true
	return s
}

// VAlign sets the vertical alignment of the cell.
// Accepts: "top", "middle", and "bottom".
func (s *CellStyle) VAlign(align string) *CellStyle {
	if s == nil {
		return nil
	}
	s.def.VAlign = align
	s.changed = true
	return s
}

// Clear marks the cell style to be completely removed when Apply is called.
func (s *CellStyle) Clear() *CellStyle {
	if s == nil {
		return nil
	}
	s.def = CellStyleDef{}
	s.changed = true
	s.clear = true
	return s
}

// Apply saves the style changes to the cell.
func (s *CellStyle) Apply() *Cell {
	if s == nil || s.cell == nil || s.cell.sheet == nil || s.cell.sheet.doc == nil {
		return nil
	}

	if !s.changed {
		return s.cell
	}

	cell := s.cell.cell()
	if cell == nil {
		return s.cell
	}

	if s.clear || s.def == (CellStyleDef{}) {
		cell.StyleName = ""
		s.cell.sheet.doc.contentDirty = true
		return s.cell
	}

	name := s.cell.sheet.doc.styles.cellStyle(s.def)
	cell.StyleName = name
	s.cell.sheet.doc.contentDirty = true
	return s.cell
}

// Bold toggles bold text on the row.
func (s *RowStyle) Bold(enabled bool) *RowStyle {
	if s == nil {
		return nil
	}
	s.def.Bold = enabled
	s.changed = true
	return s
}

// Italic toggles italic text on the row.
func (s *RowStyle) Italic(enabled bool) *RowStyle {
	if s == nil {
		return nil
	}
	s.def.Italic = enabled
	s.changed = true
	return s
}

// Height sets the row height (e.g. "1cm", "72pt").
func (s *RowStyle) Height(height string) *RowStyle {
	if s == nil {
		return nil
	}
	s.def.Height = height
	s.changed = true
	return s
}

// Clear marks the row style to be completely removed when Apply is called.
func (s *RowStyle) Clear() *RowStyle {
	if s == nil {
		return nil
	}
	s.def = RowStyleDef{}
	s.changed = true
	s.clear = true
	return s
}

// Apply saves the accumulated style changes to the row.
func (s *RowStyle) Apply() *Row {
	if s == nil || s.row == nil || s.row.sheet == nil || s.row.sheet.doc == nil {
		return nil
	}

	if !s.changed {
		return s.row
	}

	table := s.row.sheet.table()
	if table == nil || s.row.index < 0 {
		return s.row
	}

	row := ensureRow(table, s.row.index)
	if s.clear || s.def == (RowStyleDef{}) {
		row.StyleName = ""
		s.row.sheet.doc.contentDirty = true
		return s.row
	}

	row.StyleName = s.row.sheet.doc.styles.rowStyle(s.def)
	s.row.sheet.doc.contentDirty = true
	return s.row
}

// Bold toggles bold text on the column.
func (s *ColStyle) Bold(enabled bool) *ColStyle {
	if s == nil {
		return nil
	}
	s.def.Bold = enabled
	s.changed = true
	return s
}

// Italic toggles italic text on the column.
func (s *ColStyle) Italic(enabled bool) *ColStyle {
	if s == nil {
		return nil
	}
	s.def.Italic = enabled
	s.changed = true
	return s
}

// Width sets the column width (e.g. "2cm", "100pt").
func (s *ColStyle) Width(width string) *ColStyle {
	if s == nil {
		return nil
	}
	s.def.Width = width
	s.changed = true
	return s
}

// Clear marks the column style to be completely removed when Apply is called.
func (s *ColStyle) Clear() *ColStyle {
	if s == nil {
		return nil
	}
	s.def = ColStyleDef{}
	s.changed = true
	s.clear = true
	return s
}

// Apply saves the accumulated style changes to the column. If no changes
// were made, the column is returned unchanged. Returns the column for chaining.
func (s *ColStyle) Apply() *Col {
	if s == nil || s.col == nil || s.col.sheet == nil || s.col.sheet.doc == nil {
		return nil
	}

	if !s.changed {
		return s.col
	}

	table := s.col.sheet.table()
	if table == nil || s.col.index < 0 {
		return s.col
	}

	col := ensureColumn(table, s.col.index)
	if s.clear || s.def == (ColStyleDef{}) {
		col.StyleName = ""
		s.col.sheet.doc.contentDirty = true
		return s.col
	}

	col.StyleName = s.col.sheet.doc.styles.colStyle(s.def)
	s.col.sheet.doc.contentDirty = true
	return s.col
}
