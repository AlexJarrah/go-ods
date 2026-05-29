# go-ods

Go library for reading, writing, and modifying OpenDocument Spreadsheet (ODS) files.

Designed to provide a reliable, performant, and simple API for ODS documents. Archives, XML structures, repeated rows, and other spreadsheet internals are handled automatically.

## Features

- Create, read, and modify ODF spreadsheets
- Manage sheets, rows, columns, and cells
- Set values, formulas, and styles
- Automatic row and cell growth
- Efficient row and column repetition handling
- Pure Go, no external dependencies

## Installation

```bash
go get -u github.com/AlexJarrah/go-ods
```

## Example

```go
package main

import (
	"fmt"

	"github.com/AlexJarrah/go-ods"
)

func main() {
	// Open existing document
	doc, err := ods.Open("example.ods")
	if err != nil {
		panic(err)
	}

	// Modify document
	doc.SetTitle("Example")
	doc.SetDescription("Demonstration of go-ods")

	// Read cell value
	sales := doc.SheetByName("Sales")
	cell := sales.Cell(0, 1)
	text, _ := cell.String()
	fmt.Println("A2 Value:", text)

	// Write data to a new sheet
	sheet := doc.AddSheet("Example")
	sheet.Cell(0, 0).Set("Styled")
	sheet.Cell(1, 0).Set(5)
	sheet.Cell(1, 1).Set(10)
	sheet.Cell(2, 0).SetFormula("SUM(A2:B2)")
	sheet.Cell(3, 0).SetFormula("=TODAY()")
	sheet.Cell(0, 0).Style().
		Bold(true).
		Color("#0088ff").
		BackgroundColor("#323232").
		HAlign("center").
		Apply()

	// Save changes
	if err := doc.Save(); err != nil {
		panic(err)
	}
}
```
