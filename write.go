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
func Write(filepath string, ods ODS, files *zip.ReadCloser) error {
	// Translate the content for XML compatibility
	odsMarshal, err := translate(ods)
	if err != nil {
		return err
	}

	// Marshal the translated content into XML bytes
	content, err := xml.Marshal(odsMarshal.Content)
	if err != nil {
		return err
	}

	meta, err := xml.Marshal(odsMarshal.Meta)
	if err != nil {
		return err
	}

	manifest, err := xml.Marshal(odsMarshal.Manifest)
	if err != nil {
		return err
	}

	settings, err := xml.Marshal(odsMarshal.Settings)
	if err != nil {
		return err
	}

	styles, err := xml.Marshal(odsMarshal.Styles)
	if err != nil {
		return err
	}

	// Create a new zip archive in memory
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	// Add files to the archive, updating "content.xml" with the modified data
	for _, file := range files.File {
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

		switch file.Name {
		case "content.xml":
			content = append([]byte(xml.Header), content...)
			if _, err = f.Write(content); err != nil {
				return fmt.Errorf("error writing content.xml: %v", err)
			}
		case "meta.xml":
			meta = append([]byte(xml.Header), meta...)
			if _, err = f.Write(meta); err != nil {
				return fmt.Errorf("error writing meta.xml: %v", err)
			}
		case "manifest.rdf":
			manifest = append([]byte(xml.Header), manifest...)
			if _, err = f.Write(manifest); err != nil {
				return fmt.Errorf("error writing manifest.rdf: %v", err)
			}
		case "settings.xml":
			settings = append([]byte(xml.Header), settings...)
			if _, err = f.Write(settings); err != nil {
				return fmt.Errorf("error writing settings.xml: %v", err)
			}
		case "styles.xml":
			styles = append([]byte(xml.Header), styles...)
			if _, err = f.Write(styles); err != nil {
				return fmt.Errorf("error writing styles.xml: %v", err)
			}
		case "mimetype":
			if _, err = f.Write([]byte(ods.Mimetype)); err != nil {
				return fmt.Errorf("error writing styles.xml: %v", err)
			}
		default:
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

// Translates the ODS struct to a ODSMarshal struct for XML compatibility
func translate(data ODS) (odsMarshal, error) {
	// Modify the ContentMarshal struct to ensure XML compatibility
	for _, sheet := range data.Content.Body.Spreadsheet.Table {
		for _, row := range sheet.TableRow {
			for i := range row.TableCell {
				cell := &row.TableCell[i]
				cell.ValueType0 = cell.ValueType
			}
		}
	}

	// Marshal the contents to JSON, then unmarshal it back to a new struct
	// to handle XML compatibility issues
	jsonData, err := json.Marshal(data)
	if err != nil {
		return odsMarshal{}, err
	}

	var resp odsMarshal
	if err = json.Unmarshal(jsonData, &resp); err != nil {
		return odsMarshal{}, err
	}

	return resp, nil
}
