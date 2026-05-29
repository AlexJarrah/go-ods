package ods

import "strconv"

// repeatCount parses an ODS repeat-count attribute string and returns the
// count. Returns 1 for empty strings or invalid values.
func repeatCount(raw string) int {
	if raw == "" {
		return 1
	}

	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		return 1
	}

	return n
}

// setRepeat sets an ODS repeat-count attribute string. If n is 1 or less,
// the string is cleared (since repetition of 1 is the default).
func setRepeat(raw *string, n int) {
	if n <= 1 {
		*raw = ""
		return
	}
	*raw = strconv.Itoa(n)
}

// ensureRow returns the TableRow at the given logical index, materializing
// repeated rows as needed. If the row is within a repeated block, the block
// is split into before/target/after segments. If the index exceeds the
// current table size, new rows are appended. Returns nil if rowIndex is negative.
func ensureRow(table *Table, rowIndex int) *TableRow {
	if rowIndex < 0 {
		return nil
	}

	logical := 0

	for i := range table.TableRow {
		repeat := repeatCount(table.TableRow[i].NumberRowsRepeated)
		if rowIndex >= logical && rowIndex < logical+repeat {
			if repeat == 1 {
				return &table.TableRow[i]
			}

			base := table.TableRow[i]
			base.NumberRowsRepeated = ""
			replacement := make([]TableRow, 0, 3)
			if before := rowIndex - logical; before > 0 {
				beforeRow := table.TableRow[i]
				setRepeat(&beforeRow.NumberRowsRepeated, before)
				replacement = append(replacement, beforeRow)
			}

			replacement = append(replacement, base)
			if after := logical + repeat - rowIndex - 1; after > 0 {
				afterRow := table.TableRow[i]
				setRepeat(&afterRow.NumberRowsRepeated, after)
				replacement = append(replacement, afterRow)
			}

			table.TableRow = append(table.TableRow[:i], append(replacement, table.TableRow[i+1:]...)...)
			if len(replacement) == 1 {
				return &table.TableRow[i]
			}

			if rowIndex == logical {
				return &table.TableRow[i]
			}

			return &table.TableRow[i+1]
		}
		logical += repeat
	}

	template := TableRow{}
	if len(table.TableRow) > 0 {
		template.StyleName = table.TableRow[0].StyleName
	}

	for logical <= rowIndex {
		table.TableRow = append(table.TableRow, template)
		logical++
	}

	return &table.TableRow[len(table.TableRow)-1]
}

// materializeRow ensures the row at rowIndex exists and returns its position
// in the TableRow slice after materialization.
func materializeRow(table *Table, rowIndex int) int {
	ensureRow(table, rowIndex)
	logical := 0
	for i, row := range table.TableRow {
		repeat := repeatCount(row.NumberRowsRepeated)
		if rowIndex >= logical && rowIndex < logical+repeat {
			return i
		}
		logical += repeat
	}

	return len(table.TableRow) - 1
}

// existingRow returns the TableRow at the given logical index without
// auto-creating it. Returns nil if the row has not been materialized.
func existingRow(table *Table, rowIndex int) *TableRow {
	if table == nil || rowIndex < 0 {
		return nil
	}

	logical := 0
	for i := range table.TableRow {
		repeat := repeatCount(table.TableRow[i].NumberRowsRepeated)
		if rowIndex >= logical && rowIndex < logical+repeat {
			return &table.TableRow[i]
		}

		logical += repeat
	}

	return nil
}

// ensureCell returns the TableCell at the given logical column index within
// a row, materializing repeated cells as needed. If the cell is within a
// repeated block, the block is split. If the index exceeds the current row
// size, new cells are appended. Returns nil if colIndex is negative.
func ensureCell(row *TableRow, colIndex int) *TableCell {
	if colIndex < 0 {
		return nil
	}

	logical := 0
	for i := 0; i < len(row.TableCell); i++ {
		repeat := repeatCount(row.TableCell[i].NumberColumnsRepeated)
		if colIndex >= logical && colIndex < logical+repeat {
			if repeat == 1 {
				return &row.TableCell[i]
			}

			base := row.TableCell[i]
			base.NumberColumnsRepeated = ""
			replacement := make([]TableCell, 0, 3)
			if before := colIndex - logical; before > 0 {
				beforeCell := row.TableCell[i]
				setRepeat(&beforeCell.NumberColumnsRepeated, before)
				replacement = append(replacement, beforeCell)
			}

			replacement = append(replacement, base)
			if after := logical + repeat - colIndex - 1; after > 0 {
				afterCell := row.TableCell[i]
				setRepeat(&afterCell.NumberColumnsRepeated, after)
				replacement = append(replacement, afterCell)
			}

			row.TableCell = append(row.TableCell[:i], append(replacement, row.TableCell[i+1:]...)...)
			if len(replacement) == 1 {
				return &row.TableCell[i]
			}

			if colIndex == logical {
				return &row.TableCell[i]
			}

			return &row.TableCell[i+1]
		}

		logical += repeat
	}

	for logical <= colIndex {
		row.TableCell = append(row.TableCell, TableCell{})
		logical++
	}

	return &row.TableCell[len(row.TableCell)-1]
}

// materializeCell ensures the cell at colIndex exists and returns its position
// in the TableCell slice after materialization.
func materializeCell(row *TableRow, colIndex int) int {
	ensureCell(row, colIndex)
	logical := 0
	for i, cell := range row.TableCell {
		repeat := repeatCount(cell.NumberColumnsRepeated)
		if colIndex >= logical && colIndex < logical+repeat {
			return i
		}
		logical += repeat
	}

	return len(row.TableCell) - 1
}

// existingCell returns the TableCell at the given logical column index without
// auto-creating it. Returns nil if the cell has not been materialized.
func existingCell(row *TableRow, colIndex int) *TableCell {
	if row == nil || colIndex < 0 {
		return nil
	}

	logical := 0
	for i := range row.TableCell {
		repeat := repeatCount(row.TableCell[i].NumberColumnsRepeated)
		if colIndex >= logical && colIndex < logical+repeat {
			return &row.TableCell[i]
		}
		logical += repeat
	}

	return nil
}

// ensureColumn returns the TableColumn at the given logical index, materializing
// repeated columns as needed. If the column is within a repeated block, the block
// is split. If the index exceeds the current table size, new columns are appended.
// Returns nil if colIndex is negative.
func ensureColumn(table *Table, colIndex int) *TableColumn {
	if colIndex < 0 {
		return nil
	}

	logical := 0
	for i := 0; i < len(table.TableColumn); i++ {
		repeat := repeatCount(table.TableColumn[i].NumberColumnsRepeated)
		if colIndex >= logical && colIndex < logical+repeat {
			if repeat == 1 {
				return &table.TableColumn[i]
			}

			base := table.TableColumn[i]
			base.NumberColumnsRepeated = ""
			replacement := make([]TableColumn, 0, 3)
			if before := colIndex - logical; before > 0 {
				beforeCol := table.TableColumn[i]
				setRepeat(&beforeCol.NumberColumnsRepeated, before)
				replacement = append(replacement, beforeCol)
			}

			replacement = append(replacement, base)
			if after := logical + repeat - colIndex - 1; after > 0 {
				afterCol := table.TableColumn[i]
				setRepeat(&afterCol.NumberColumnsRepeated, after)
				replacement = append(replacement, afterCol)
			}

			table.TableColumn = append(table.TableColumn[:i], append(replacement, table.TableColumn[i+1:]...)...)
			if len(replacement) == 1 || colIndex == logical {
				return &table.TableColumn[i]
			}

			return &table.TableColumn[i+1]
		}

		logical += repeat
	}

	for logical <= colIndex {
		table.TableColumn = append(table.TableColumn, TableColumn{})
		logical++
	}

	return &table.TableColumn[len(table.TableColumn)-1]
}

// materializeColumn ensures the column at colIndex exists and returns its
// position in the TableColumn slice after materialization.
func materializeColumn(table *Table, colIndex int) int {
	ensureColumn(table, colIndex)
	logical := 0
	for i, col := range table.TableColumn {
		repeat := repeatCount(col.NumberColumnsRepeated)
		if colIndex >= logical && colIndex < logical+repeat {
			return i
		}
		logical += repeat
	}

	return len(table.TableColumn) - 1
}

// compressTable optimizes a table's storage by merging adjacent identical
// rows and columns into repeated entries, and trimming trailing empty rows,
// columns, and cells.
func compressTable(table *Table) {
	for i := range table.TableRow {
		compressCells(&table.TableRow[i])
	}
	table.TableRow = compressRows(table.TableRow)
	table.TableColumn = compressColumns(table.TableColumn)
	trimTrailingEmpty(table)
}

// compressCells merges adjacent identical cells within a row into a single
// cell with an updated NumberColumnsRepeated count.
func compressCells(row *TableRow) {
	if len(row.TableCell) == 0 {
		return
	}

	out := make([]TableCell, 0, len(row.TableCell))
	current := row.TableCell[0]
	count := repeatCount(row.TableCell[0].NumberColumnsRepeated)
	current.NumberColumnsRepeated = ""

	for _, next := range row.TableCell[1:] {
		nextCount := repeatCount(next.NumberColumnsRepeated)
		next.NumberColumnsRepeated = ""
		if cellsEqual(current, next) {
			count += nextCount
			continue
		}

		setRepeat(&current.NumberColumnsRepeated, count)
		out = append(out, current)
		current = next
		count = nextCount
	}

	setRepeat(&current.NumberColumnsRepeated, count)
	out = append(out, current)
	row.TableCell = out
}

// compressRows merges adjacent identical rows into a single row with an
// updated NumberRowsRepeated count.
func compressRows(rows []TableRow) []TableRow {
	if len(rows) == 0 {
		return rows
	}

	out := make([]TableRow, 0, len(rows))
	current := rows[0]
	count := repeatCount(rows[0].NumberRowsRepeated)
	current.NumberRowsRepeated = ""

	for _, next := range rows[1:] {
		nextCount := repeatCount(next.NumberRowsRepeated)
		next.NumberRowsRepeated = ""
		if rowsEqual(current, next) {
			count += nextCount
			continue
		}

		setRepeat(&current.NumberRowsRepeated, count)
		out = append(out, current)
		current = next
		count = nextCount
	}

	setRepeat(&current.NumberRowsRepeated, count)
	out = append(out, current)
	return out
}

// compressColumns merges adjacent identical columns into a single column
// with an updated NumberColumnsRepeated count.
func compressColumns(cols []TableColumn) []TableColumn {
	if len(cols) == 0 {
		return cols
	}

	out := make([]TableColumn, 0, len(cols))
	current := cols[0]
	count := repeatCount(cols[0].NumberColumnsRepeated)
	current.NumberColumnsRepeated = ""

	for _, next := range cols[1:] {
		nextCount := repeatCount(next.NumberColumnsRepeated)
		next.NumberColumnsRepeated = ""
		if columnsEqual(current, next) {
			count += nextCount
			continue
		}
		setRepeat(&current.NumberColumnsRepeated, count)
		out = append(out, current)
		current = next
		count = nextCount
	}

	setRepeat(&current.NumberColumnsRepeated, count)
	out = append(out, current)
	return out
}

// cellsEqual returns whether two TableCell values are equal, ignoring
// their NumberColumnsRepeated fields.
func cellsEqual(a, b TableCell) bool {
	a.NumberColumnsRepeated = ""
	b.NumberColumnsRepeated = ""
	return a == b
}

// rowsEqual returns whether two TableRow values are equal, comparing
// style names and all cells (ignoring repeat counts).
func rowsEqual(a, b TableRow) bool {
	if a.StyleName != b.StyleName || len(a.TableCell) != len(b.TableCell) {
		return false
	}
	for i := range a.TableCell {
		if !cellsEqual(a.TableCell[i], b.TableCell[i]) {
			return false
		}
	}
	return true
}

// columnsEqual returns whether two TableColumn values are equal, ignoring
// their NumberColumnsRepeated fields.
func columnsEqual(a, b TableColumn) bool {
	a.NumberColumnsRepeated = ""
	b.NumberColumnsRepeated = ""
	return a == b
}

// usedRowCount returns the number of rows that contain data or styling,
// accounting for row repetition.
func usedRowCount(table *Table) int {
	count := 0
	logical := 0
	for _, row := range table.TableRow {
		repeat := repeatCount(row.NumberRowsRepeated)
		if !rowIsEmpty(row) {
			count = logical + repeat
		}
		logical += repeat
	}
	return count
}

// usedColCount returns the maximum number of columns used across all rows,
// accounting for column repetition.
func usedColCount(table *Table) int {
	maxCol := 0
	for _, row := range table.TableRow {
		logical := 0
		for _, cell := range row.TableCell {
			repeat := repeatCount(cell.NumberColumnsRepeated)
			if !cellIsEmpty(cell) && logical+repeat > maxCol {
				maxCol = logical + repeat
			}
			logical += repeat
		}
	}
	return maxCol
}

// rowIsEmpty returns whether a row has no styling and contains only empty cells.
func rowIsEmpty(row TableRow) bool {
	if row.StyleName != "" {
		return false
	}
	for _, cell := range row.TableCell {
		if !cellIsEmpty(cell) {
			return false
		}
	}
	return true
}

// cellIsEmpty returns whether a cell has no content or styling, ignoring
// the NumberColumnsRepeated field.
func cellIsEmpty(cell TableCell) bool {
	cell.NumberColumnsRepeated = ""
	return cell == TableCell{}
}

// trimTrailingEmpty removes trailing empty rows, columns, and cells from
// a table to minimize its on-disk representation.
func trimTrailingEmpty(table *Table) {
	for i := range table.TableRow {
		trimTrailingEmptyCells(&table.TableRow[i])
	}

	for len(table.TableRow) > 1 && rowIsEmpty(table.TableRow[len(table.TableRow)-1]) {
		table.TableRow = table.TableRow[:len(table.TableRow)-1]
	}
	if len(table.TableRow) == 0 {
		table.TableRow = []TableRow{{TableCell: []TableCell{{}}}}
	}
	if len(table.TableRow) == 1 && rowIsEmpty(table.TableRow[0]) {
		table.TableRow[0].NumberRowsRepeated = ""
	}

	for len(table.TableColumn) > 1 && tableColumnIsEmpty(table.TableColumn[len(table.TableColumn)-1]) {
		table.TableColumn = table.TableColumn[:len(table.TableColumn)-1]
	}
	if len(table.TableColumn) == 0 {
		table.TableColumn = []TableColumn{{}}
	}
	if len(table.TableColumn) == 1 && tableColumnIsEmpty(table.TableColumn[0]) {
		table.TableColumn[0].NumberColumnsRepeated = ""
	}
}

// trimTrailingEmptyCells removes trailing empty cells from a row, ensuring
// at least one cell remains.
func trimTrailingEmptyCells(row *TableRow) {
	for len(row.TableCell) > 1 && cellIsEmpty(row.TableCell[len(row.TableCell)-1]) {
		row.TableCell = row.TableCell[:len(row.TableCell)-1]
	}
	if len(row.TableCell) == 0 {
		row.TableCell = []TableCell{{}}
	}
	if len(row.TableCell) == 1 && cellIsEmpty(row.TableCell[0]) {
		row.TableCell[0].NumberColumnsRepeated = ""
	}
}

// tableColumnIsEmpty returns whether a TableColumn has no styling,
// ignoring the NumberColumnsRepeated field.
func tableColumnIsEmpty(col TableColumn) bool {
	col.NumberColumnsRepeated = ""
	return col == TableColumn{}
}

// coalesce returns value or a fallback, depending on if value is zero.
func coalesce[T comparable](value, fallback T) T {
	var zero T
	if value != zero {
		return value
	}
	return fallback
}
