package ods

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Cell represents a single cell in an ODS spreadsheet. It is identified by
// its row and column indices (zero-based) within a sheet. Underlying data is
// not materialized until a read or write operation is performed.
type Cell struct {
	sheet *Sheet
	row   int
	col   int
}

// cell returns the underlying TableCell if it exists, materializing the row
// and cell if they do not yet exist.
func (c *Cell) cell() *TableCell {
	if c == nil || c.sheet == nil || c.sheet.doc == nil || c.row < 0 || c.col < 0 {
		return nil
	}

	table := c.sheet.table()
	if table == nil {
		return nil
	}

	return ensureCell(ensureRow(table, c.row), c.col)
}

// existingCell returns the underlying TableCell only if it has already been
// materialized. Returns (nil, nil) if the cell does not exist yet, avoiding
// auto-creation.
func (c *Cell) existingCell() (*TableCell, error) {
	if c == nil || c.sheet == nil || c.sheet.doc == nil || c.row < 0 || c.col < 0 {
		return nil, c.err()
	}

	table := c.sheet.table()
	if table == nil {
		return nil, c.err()
	}

	row := existingRow(table, c.row)
	if row == nil {
		return nil, nil
	}

	return existingCell(row, c.col), nil
}

// Set sets the cell value to the given Go value. Supported types are:
// string, []byte, int/int8/int16/int32/int64, uint/uint8/uint16/uint32/uint64,
// float32, float64, bool, time.Time, and nil (clears the cell).
func (c *Cell) Set(value any) *Cell {
	cell := c.cell()
	if cell == nil {
		return c
	}

	if err := setTypedValue(cell, value); err != nil {
		c.sheet.doc.setErr(err)
		return c
	}

	c.sheet.doc.contentDirty = true
	return c
}

// SetFormula sets the cell to contain a spreadsheet formula. The formula
// string is automatically normalized: A1-style references are converted to
// ODF bracket notation, and the appropriate prefix is added.
func (c *Cell) SetFormula(formula string) *Cell {
	cell := c.cell()
	if cell == nil {
		return c
	}

	cell.Formula = normalizeFormula(formula)
	if cell.ValueType == "" {
		setFloatCell(cell, "0")
	}

	c.sheet.doc.contentDirty = true
	return c
}

// Style returns a CellStyle builder for configuring the cell's appearance.
// Call Apply on the returned builder to save changes.
func (c *Cell) Style() *CellStyle {
	if cell, _ := c.existingCell(); cell != nil {
		return &CellStyle{cell: c, def: c.sheet.doc.styles.cellStyleDef(cell.StyleName)}
	}
	return &CellStyle{cell: c, def: CellStyleDef{}}
}

// String returns the cell's text content as a string. It checks the P
// (paragraph) element first, then the Value attribute. Returns an empty
// string if the cell does not exist.
func (c *Cell) String() (string, error) {
	cell, err := c.existingCell()
	if err != nil {
		return "", err
	}
	if cell == nil {
		return "", nil
	}
	if cell.P != "" {
		return cell.P, nil
	}
	return cell.Value, nil
}

// Float64 returns the cell's numeric value as a float64. It checks the
// Value attribute first, then the P element. Returns 0 if the cell does
// not exist or has no numeric content.
func (c *Cell) Float64() (float64, error) {
	cell, err := c.existingCell()
	if err != nil {
		return 0, err
	}
	if cell == nil {
		return 0, nil
	}
	raw := cell.Value
	if raw == "" {
		raw = cell.P
	}
	if raw == "" {
		return 0, nil
	}
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("cell (%d,%d): failed to parse float %q: %w", c.row, c.col, raw, err)
	}
	return f, nil
}

// Bool returns the cell's boolean value. It checks the Value attribute
// first, then the P element. Returns false if the cell does not exist.
func (c *Cell) Bool() (bool, error) {
	cell, err := c.existingCell()
	if err != nil {
		return false, err
	}
	if cell == nil {
		return false, nil
	}
	raw := cell.Value
	if raw == "" {
		raw = cell.P
	}
	if raw == "" {
		return false, nil
	}
	b, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("cell (%d,%d): failed to parse bool %q: %w", c.row, c.col, raw, err)
	}
	return b, nil
}

// Time returns the cell's date/time value. It checks the date-value
// attribute first, then the value attribute. It attempts to parse using
// RFC3339Nano, RFC3339, and "2006-01-02" layouts. Returns a zero time
// if the cell does not exist or has no date content.
func (c *Cell) Time() (time.Time, error) {
	cell, err := c.existingCell()
	if err != nil {
		return time.Time{}, err
	}
	if cell == nil {
		return time.Time{}, nil
	}
	raw := cell.DateValue
	if raw == "" {
		raw = cell.Value
	}
	if raw == "" {
		return time.Time{}, nil
	}
	for _, layout := range []string{time.RFC3339Nano, time.RFC3339, "2006-01-02"} {
		t, err := time.Parse(layout, raw)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cell (%d,%d): failed to parse time %q: %w", c.row, c.col, raw, err)
}

// Formula returns the cell's formula string. Returns an empty string if the
// cell does not exist or contains no formula.
func (c *Cell) Formula() (string, error) {
	cell, err := c.existingCell()
	if err != nil {
		return "", err
	}
	if cell == nil {
		return "", nil
	}
	return cell.Formula, nil
}

// err returns the document-level error if available, or a generic invalid
// cell error.
func (c *Cell) err() error {
	if c != nil && c.sheet != nil && c.sheet.doc != nil && c.sheet.doc.err != nil {
		return c.sheet.doc.err
	}
	return fmt.Errorf("cell is invalid")
}

// setTypedValue dispatches on the Go type of value and sets the appropriate
// cell fields (ValueType, Value, P, DateValue, etc.). Supported types are
// string, []byte, int*, uint*, float*, bool, time.Time, and nil (clear cell).
func setTypedValue(cell *TableCell, value any) error {
	next := TableCell{StyleName: cell.StyleName}

	switch v := value.(type) {
	case string:
		next.ValueType = cellValueTypeString
		next.CalcextValueType = cellValueTypeString
		next.P = v
	case []byte:
		next.ValueType = cellValueTypeString
		next.CalcextValueType = cellValueTypeString
		next.P = string(v)
	case int:
		setFloatCell(&next, strconv.FormatInt(int64(v), 10))
	case int8:
		setFloatCell(&next, strconv.FormatInt(int64(v), 10))
	case int16:
		setFloatCell(&next, strconv.FormatInt(int64(v), 10))
	case int32:
		setFloatCell(&next, strconv.FormatInt(int64(v), 10))
	case int64:
		setFloatCell(&next, strconv.FormatInt(v, 10))
	case uint:
		setFloatCell(&next, strconv.FormatUint(uint64(v), 10))
	case uint8:
		setFloatCell(&next, strconv.FormatUint(uint64(v), 10))
	case uint16:
		setFloatCell(&next, strconv.FormatUint(uint64(v), 10))
	case uint32:
		setFloatCell(&next, strconv.FormatUint(uint64(v), 10))
	case uint64:
		setFloatCell(&next, strconv.FormatUint(v, 10))
	case float32:
		setFloatCell(&next, strconv.FormatFloat(float64(v), 'g', -1, 32))
	case float64:
		setFloatCell(&next, strconv.FormatFloat(v, 'g', -1, 64))
	case bool:
		next.ValueType = cellValueTypeBoolean
		next.CalcextValueType = cellValueTypeBoolean
		next.Value = strconv.FormatBool(v)
		next.P = next.Value
	case time.Time:
		next.ValueType = cellValueTypeDate
		next.CalcextValueType = cellValueTypeDate
		next.DateValue = v.Format(time.RFC3339)
		next.P = v.Format("2006-01-02")
	case nil:
		next = TableCell{}
	default:
		return fmt.Errorf("unsupported cell value type: %T", value)
	}
	*cell = next
	return nil
}

// setFloatCell sets a TableCell to the float value-type with the given
// raw string representation.
func setFloatCell(cell *TableCell, raw string) {
	cell.ValueType = cellValueTypeFloat
	cell.CalcextValueType = cellValueTypeFloat
	cell.Value = raw
	cell.P = raw
}

// normalizeFormula normalizes a formula string for ODS compatibility.
// It prepends "of:" as needed, handles "=" and ":=" prefixes, and
// converts A1-style references to ODF bracket notation.
func normalizeFormula(formula string) string {
	formula = strings.TrimSpace(formula)
	if formula == "" {
		return ""
	}
	if strings.Contains(formula, ":=") {
		return normalizeFormulaRefs(formula)
	}
	if strings.HasPrefix(formula, "=") {
		return "of:" + normalizeFormulaRefs(formula)
	}
	return "of:=" + normalizeFormulaRefs(formula)
}

// normalizeFormulaRefs converts A1-style cell references (e.g. A1, B2:C10)
// within a formula to ODF bracket notation (e.g. [.A1], [.B2:.C10]).
// Existing bracket-quoted references are preserved.
func normalizeFormulaRefs(formula string) string {
	var out strings.Builder
	for i := 0; i < len(formula); {
		if formula[i] == '[' {
			end := strings.IndexByte(formula[i:], ']')
			if end < 0 {
				out.WriteByte(formula[i])
				i++
				continue
			}
			end += i
			out.WriteString(formula[i : end+1])
			i = end + 1
			continue
		}

		ref, next := readA1Ref(formula, i)
		if ref == "" {
			out.WriteByte(formula[i])
			i++
			continue
		}

		if next < len(formula) && formula[next] == ':' {
			ref2, next2 := readA1Ref(formula, next+1)
			if ref2 != "" {
				out.WriteString("[.")
				out.WriteString(ref)
				out.WriteString(":.")
				out.WriteString(ref2)
				out.WriteByte(']')
				i = next2
				continue
			}
		}

		out.WriteString("[.")
		out.WriteString(ref)
		out.WriteByte(']')
		i = next
	}
	return out.String()
}

// readA1Ref parses a single A1-style cell reference starting at position
// start in the formula string. It returns the normalized uppercase reference
// (with $ signs removed) and the position after the reference, or ("", start)
// if no valid reference is found.
func readA1Ref(s string, start int) (string, int) {
	if start > 0 && isFormulaIdent(rune(s[start-1])) {
		return "", start
	}

	i := start
	for i < len(s) && s[i] == '$' {
		i++
	}
	colStart := i
	for i < len(s) && ((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')) {
		i++
	}
	if i == colStart || i-colStart > 3 {
		return "", start
	}
	if i < len(s) && s[i] == '$' {
		i++
	}
	rowStart := i
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	if i == rowStart {
		return "", start
	}
	if i < len(s) && isFormulaIdent(rune(s[i])) {
		return "", start
	}

	ref := strings.ToUpper(strings.ReplaceAll(s[start:i], "$", ""))
	return ref, i
}

// isFormulaIdent returns whether r is a valid identifier character in
// the context of a formula.
func isFormulaIdent(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.'
}
