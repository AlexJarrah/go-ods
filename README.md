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

```go
package main

import (
	"log"

	"github.com/AlexJarrah/go-ods"
)

func main() {
    // Specify a filepath
    const path = "/home/user/Documents/example.ods"

    // Reading data from the file
	content, files, err := ods.Read(path)
	if err != nil {
		log.Panic(err)
	}
	defer files.Close()

	// Updating content data in a specific sheet
	sheet := content.Body.Spreadsheet.Table[1]
	sheet.TableRow[1].TableCell[0].P = "updated value"
	sheet.TableRow[2].TableCell[0].P = "updated value"
	sheet.TableRow[2].TableCell[3].P = "updated value"

	// Writing new data to the file
	if err := ods.Write(path, content, files); err != nil {
		log.Panic(err)
	}
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/AlexJarrah/go-ods/blob/main/LICENSE) file for details.
