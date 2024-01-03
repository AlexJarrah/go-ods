# go-ods

## Overview

go-ods is a Go package designed to simplify reading and writing ODS (OpenDocument Spreadsheet) files. With this library, you can seamlessly work with ODS files in your Go applications, providing support for common spreadsheet operations.

## Features

- Read ODS Files: Easily parse and extract data from ODS files.
- Write ODS Files: Update ODS files with new/updated data.
- Cell Manipulation: Perform operations on individual cells, such as setting values, formatting, and more.
- Sheet Handling: Manage multiple sheets within a single ODS file effortlessly.

## Installation

To use this go-ods, [install Go](https://go.dev) and run:

```bash
go get -u github.com/AlexJarrah/go-ods
```

## Usage

````go
package main

import "github.com/AlexJarrah/go-ods"

func main() {
	// Specify a filepath
	const path = "/home/user/Documents/example.ods"

	// Reading data from the file
	data, files, err := ods.Read(path)
	if err != nil {
		panic(err)
	}
	defer files.Close()

	// Uncompress cells to improve consistency
	data.Content = ods.Uncompress(data.Content, 20)

	// Updating content data in a specific sheet
	sheet := data.Content.Body.Spreadsheet.Table[1]
	sheet.TableRow[1].TableCell[0].P = "updated value"
	sheet.TableRow[2].TableCell[0].P = "updated value"
	sheet.TableRow[2].TableCell[3].P = "updated value"

	// Updating metadata
	data.Meta.Meta.DocumentStatistic.CellCount = "1000"

	// Writing new data to the file
	data.Content.Body.Spreadsheet.Table[1] = sheet
	if err := ods.Write(path, data, files); err != nil {
		panic(err)
	}
}```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/AlexJarrah/go-ods/blob/main/LICENSE) file for details.
````
