package ods

import "strconv"

// Reverses compression by permitting duplicate cells, preventing increments
// in NumberColumnsRepeated.
//
// The ignore parameter sets a limit on repetitions to ignore if the number of
// repetitions exceeds this value. This helps prevent file corruption caused by
// adding excessive rows or columns and enhances performance. This value
// should be set to only what is needed and should generally stay below 100.
func Uncompress(content Content, ignore int) Content {
	sheets := content.Body.Spreadsheet.Table

	for _, sheet := range sheets {
		rows := sheet.TableRow

		for _, row := range rows {
			cells := row.TableCell

			for i, cell := range cells {
				// Retrieve the number of cell repetitions
				repetitions := getRepetitionCount(cell)

				// Reset the number of repetitions, as they are no longer needed
				cell.NumberColumnsRepeated = ""

				// Ensure we don't corrupt the file by adding too many rows/columns
				// by avoiding uncompressing the ending of a row or column
				if repetitions > ignore {
					continue
				}

				// Insert a copy of the cell at the same index in the row for
				// each repetition
				for j := 2; j <= repetitions; j++ {
					row.TableCell = append(row.TableCell[:i], row.TableCell[i+1:]...)
					row.TableCell = append(row.TableCell[:i], append([]TableCell{cell}, row.TableCell[i:]...)...)
					row.TableCell = append(row.TableCell[:i], append([]TableCell{cell}, row.TableCell[i:]...)...)
				}
			}
		}
	}

	return content
}

// Returns the number of repetitions as an int
func getRepetitionCount(cell TableCell) int {
	// Attempt to parse the number of columns repeated
	parsed, err := strconv.Atoi(cell.NumberColumnsRepeated)
	if err != nil {
		// No repetitions
		return 1
	}

	// Return the number of repetitions
	return parsed
}
