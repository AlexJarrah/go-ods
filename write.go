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

// Updates the specified file with the provided content
func Write(filepath string, content Content, r *zip.ReadCloser) error {
	contentMarshal, err := translate(content)
	if err != nil {
		return err
	}

	// Marshal Content struct into bytes
	data, err := xml.Marshal(contentMarshal)
	if err != nil {
		return err
	}

	// Create a buffer
	buf := new(bytes.Buffer)

	// Create new zip archive
	w := zip.NewWriter(buf)

	// Add files to archive
	for _, file := range r.File {
		f, err := w.Create(file.Name)
		if err != nil {
			return fmt.Errorf("error creating file in archive: %v", err)
		}

		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("error opening file in archive: %v", err)
		}
		defer rc.Close()

		if file.Name == "content.xml" {
			data = append([]byte(xml.Header), data...)
			if _, err = f.Write(data); err != nil {
				return fmt.Errorf("error writing content.xml: %v", err)
			}
		} else {
			if _, err = io.Copy(f, rc); err != nil {
				return fmt.Errorf("error writing file to archive: %v", err)
			}
		}
	}

	// Close zip writer before writing buffer to file
	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing archive: %v", err)
	}

	outputFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Write the new .ods file
	_, err = outputFile.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func translate(content Content) (ContentMarshal, error) {
	for _, sheet := range content.Body.Spreadsheet.Table {
		for _, row := range sheet.TableRow {
			for i := range row.TableCell {
				cell := &row.TableCell[i]
				cell.ValueType0 = cell.ValueType
			}
		}
	}

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
