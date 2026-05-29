package ods

import "fmt"

// Sheet represents a single worksheet within an ODS document.
type Sheet struct {
	doc   *Document
	index int
}

// Row represents a single row within a sheet.
type Row struct {
	sheet *Sheet
	index int
}

// Col represents a single column within a sheet.
type Col struct {
	sheet *Sheet
	index int
}

// table returns a pointer to the underlying Table for this sheet if valid.
func (s *Sheet) table() *Table {
	if s == nil || s.doc == nil || s.index < 0 || s.index >= len(s.doc.content.Body.Spreadsheet.Table) {
		return nil
	}
	return &s.doc.content.Body.Spreadsheet.Table[s.index]
}

// Name returns the name of the sheet.
func (s *Sheet) Name() string {
	if table := s.table(); table != nil {
		return table.Name
	}
	return ""
}

// SetName renames the sheet.
func (s *Sheet) SetName(name string) *Sheet {
	if table := s.table(); table != nil {
		table.Name = name
		s.doc.contentDirty = true
	}
	return s
}

// Cell returns a Cell handle for the given zero-based row and column indices.
// The cell is not materialized until a read or write operation is performed.
func (s *Sheet) Cell(row, col int) *Cell {
	if s == nil || s.doc == nil {
		return &Cell{sheet: s, row: row, col: col}
	}

	if row < 0 || col < 0 {
		s.doc.setErr(fmt.Errorf("cell coordinates must be non-negative: row=%d col=%d", row, col))
		return &Cell{sheet: s, row: row, col: col}
	}

	table := s.table()
	if table == nil {
		s.doc.setErr(fmt.Errorf("sheet is nil"))
		return &Cell{sheet: s, row: row, col: col}
	}

	return &Cell{sheet: s, row: row, col: col}
}

// Row returns a Row handle for the given zero-based index. The row is created
// if it does not yet exist in the underlying table.
func (s *Sheet) Row(index int) *Row {
	if s == nil || s.doc == nil {
		return &Row{sheet: s, index: index}
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("row index must be non-negative: %d", index))
		return &Row{sheet: s, index: index}
	}

	if table := s.table(); table != nil {
		ensureRow(table, index)
		s.doc.contentDirty = true
	}

	return &Row{sheet: s, index: index}
}

// Col returns a Col handle for the given zero-based index. The column is
// created if it does not yet exist in the underlying table.
func (s *Sheet) Col(index int) *Col {
	if s == nil || s.doc == nil {
		return &Col{sheet: s, index: index}
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("column index must be non-negative: %d", index))
		return &Col{sheet: s, index: index}
	}

	if table := s.table(); table != nil {
		ensureColumn(table, index)
		s.doc.contentDirty = true
	}

	return &Col{sheet: s, index: index}
}

// AppendRow appends a new row at the end of the sheet with the given values.
func (s *Sheet) AppendRow(values ...any) *Sheet {
	return s.InsertRow(s.RowCount(), values...)
}

// InsertRow inserts a new row at the given zero-based index, shifting
// existing rows down. The row is populated with the provided values.
func (s *Sheet) InsertRow(index int, values ...any) *Sheet {
	if s == nil || s.doc == nil {
		return s
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("row index must be non-negative: %d", index))
		return s
	}

	table := s.table()
	if table == nil {
		return s
	}

	row := TableRow{}
	for i, value := range values {
		cell := ensureCell(&row, i)
		if err := setTypedValue(cell, value); err != nil {
			s.doc.setErr(err)
			return s
		}
	}

	pos := materializeRow(table, index)
	table.TableRow = append(table.TableRow[:pos], append([]TableRow{row}, table.TableRow[pos:]...)...)
	s.doc.contentDirty = true
	return s
}

// DeleteRow removes the row at the given zero-based index, shifting
// subsequent rows up.
func (s *Sheet) DeleteRow(index int) *Sheet {
	if s == nil || s.doc == nil {
		return s
	}

	table := s.table()
	if table == nil {
		return s
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("row index out of range: %d", index))
		return s
	}

	pos := materializeRow(table, index)
	if pos < 0 || pos >= len(table.TableRow) {
		s.doc.setErr(fmt.Errorf("row index out of range: %d", index))
		return s
	}

	table.TableRow = append(table.TableRow[:pos], table.TableRow[pos+1:]...)
	s.doc.contentDirty = true
	return s
}

// InsertCol inserts a new column at the given zero-based index, shifting
// existing columns to the right. A new cell is added to every row at that position.
func (s *Sheet) InsertCol(index int) *Sheet {
	if s == nil || s.doc == nil {
		return s
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("column index must be non-negative: %d", index))
		return s
	}

	table := s.table()
	if table == nil {
		return s
	}

	pos := materializeColumn(table, index)
	table.TableColumn = append(table.TableColumn[:pos], append([]TableColumn{{}}, table.TableColumn[pos:]...)...)
	for i := range table.TableRow {
		cellPos := materializeCell(&table.TableRow[i], index)
		table.TableRow[i].TableCell = append(table.TableRow[i].TableCell[:cellPos], append([]TableCell{{}}, table.TableRow[i].TableCell[cellPos:]...)...)
	}

	s.doc.contentDirty = true
	return s
}

// DeleteCol removes the column at the given zero-based index, shifting
// subsequent columns to the left.
func (s *Sheet) DeleteCol(index int) *Sheet {
	if s == nil || s.doc == nil {
		return s
	}

	table := s.table()
	if table == nil {
		return s
	}

	if index < 0 {
		s.doc.setErr(fmt.Errorf("column index out of range: %d", index))
		return s
	}

	pos := materializeColumn(table, index)
	if pos < 0 || pos >= len(table.TableColumn) {
		s.doc.setErr(fmt.Errorf("column index out of range: %d", index))
		return s
	}

	table.TableColumn = append(table.TableColumn[:pos], table.TableColumn[pos+1:]...)
	for i := range table.TableRow {
		cellPos := materializeCell(&table.TableRow[i], index)
		if cellPos < len(table.TableRow[i].TableCell) {
			table.TableRow[i].TableCell = append(table.TableRow[i].TableCell[:cellPos], table.TableRow[i].TableCell[cellPos+1:]...)
		}
	}

	s.doc.contentDirty = true
	return s
}

// RowCount returns the number of rows that contain data or styling in the sheet.
func (s *Sheet) RowCount() int {
	if table := s.table(); table != nil {
		return usedRowCount(table)
	}
	return 0
}

// ColCount returns the maximum number of columns used across all rows in the sheet.
func (s *Sheet) ColCount() int {
	if table := s.table(); table != nil {
		return usedColCount(table)
	}
	return 0
}

// Style returns a RowStyle builder for configuring the row's appearance.
// Call Apply on the returned builder to save changes.
func (r *Row) Style() *RowStyle {
	if r == nil || r.sheet == nil || r.sheet.doc == nil || r.index < 0 {
		return &RowStyle{row: r, def: RowStyleDef{}}
	}

	table := r.sheet.table()
	if table == nil {
		return &RowStyle{row: r, def: RowStyleDef{}}
	}

	row := ensureRow(table, r.index)
	return &RowStyle{row: r, def: r.sheet.doc.styles.rowStyleDef(row.StyleName)}
}

// Style returns a ColStyle builder for configuring the column's appearance.
// Call Apply on the returned builder to save changes.
func (c *Col) Style() *ColStyle {
	if c == nil || c.sheet == nil || c.sheet.doc == nil || c.index < 0 {
		return &ColStyle{col: c, def: ColStyleDef{}}
	}

	table := c.sheet.table()
	if table == nil {
		return &ColStyle{col: c, def: ColStyleDef{}}
	}

	col := ensureColumn(table, c.index)
	return &ColStyle{col: c, def: c.sheet.doc.styles.colStyleDef(col.StyleName)}
}
