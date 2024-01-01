package ods

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Updates the specified ODS file with the provided content
func Write(filepath string, content Content, r *zip.ReadCloser) error {
	// Translate the content for XML compatibility
	contentMarshal, err := translate(content)
	if err != nil {
		return err
	}

	// Marshal the translated content into XML bytes
	data, err := xml.Marshal(contentMarshal)
	if err != nil {
		return err
	}

	// Create a new zip archive in memory
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	// Add files to the archive, updating "content.xml" with the modified data
	for _, file := range r.File {
		// Create a new file entry in the archive
		f, err := w.Create(file.Name)
		if err != nil {
			return fmt.Errorf("error creating file in archive: %v", err)
		}

		// Open the original file from the input ODS archive
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("error opening file in archive: %v", err)
		}
		defer rc.Close()

		if file.Name == "content.xml" {
			// Write either the modified "content.xml"
			data = append([]byte(xml.Header), data...)
			if _, err = f.Write(data); err != nil {
				return fmt.Errorf("error writing content.xml: %v", err)
			}
		} else {
			// Copy all other files as-is
			if _, err = io.Copy(f, rc); err != nil {
				return fmt.Errorf("error writing file to archive: %v", err)
			}
		}
	}

	// Close the zip writer
	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing archive: %v", err)
	}

	outputFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Write the new ODS file
	_, err = outputFile.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

// Translates the Content struct to a ContentMarshal struct for XML compatibility
func translate(content Content) (ContentMarshal, error) {
	// Modify the Content struct to ensure XML compatibility
	for _, sheet := range content.Body.Spreadsheet.Table {
		for _, row := range sheet.TableRow {
			for i := range row.TableCell {
				cell := &row.TableCell[i]
				cell.ValueType0 = cell.ValueType
			}
		}
	}

	// Marshal the content to JSON, then unmarshal it back to a new struct
	// to handle XML compatibility issues
	jsonData, err := json.Marshal(content)
	if err != nil {
		return ContentMarshal{}, err
	}

	var contentMarshal ContentMarshal
	if err = json.Unmarshal(jsonData, &contentMarshal); err != nil {
		return ContentMarshal{}, err
	}

	return contentMarshal, nil
}
