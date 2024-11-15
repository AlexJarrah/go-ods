package ods

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Read the spreadsheet data from an ODS file and provides access to the archive.
func Read(filepath string) (ODS, *zip.ReadCloser, error) {
	// Open the ODS file as a zip archive
	files, err := os.Open(filepath)
	if err != nil {
		return ODS{}, nil, fmt.Errorf("error opening ODS file: %v", err)
	}
	defer files.Close()

	// Get the file stats
	fileInfo, err := files.Stat()
	if err != nil {
		return ODS{}, nil, fmt.Errorf("error getting file info: %v", err)
	}

	return ReadFrom(files, fileInfo.Size())
}

// ReadFrom Read the data from io reader of provided size.
func ReadFrom(reader io.Reader, size int64) (ODS, *zip.ReadCloser, error) {
	buf := new(bytes.Buffer)
	data := ODS{}

	// Copy data from the reader to the buffer
	_, err := io.Copy(buf, reader)
	if err != nil {
		return data, nil, fmt.Errorf("error copying data: %v", err)
	}

	files, err := zip.NewReader(bytes.NewReader(buf.Bytes()), size)
	if err != nil {
		return data, nil, fmt.Errorf("error creating zip archive: %v", err)
	}

	// Iterate through the files within the zip archive
	for _, file := range files.File {
		// Open the file for processing
		rc, err := file.Open()
		if err != nil {
			return data, nil, fmt.Errorf("error opening file: %v", err)
		}
		defer rc.Close()

		// Decode the file contents into the ODS struct
		switch file.Name {
		case "content.xml":
			if err = xml.NewDecoder(rc).Decode(&data.Content); err != nil {
				return ODS{}, nil, fmt.Errorf("error decoding file: %v", err)
			}
		case "meta.xml":
			if err = xml.NewDecoder(rc).Decode(&data.Meta); err != nil {
				return ODS{}, nil, fmt.Errorf("error decoding file: %v", err)
			}
		case "manifest.rdf":
			if err = xml.NewDecoder(rc).Decode(&data.Manifest); err != nil {
				return ODS{}, nil, fmt.Errorf("error decoding file: %v", err)
			}
		case "settings.xml":
			if err = xml.NewDecoder(rc).Decode(&data.Settings); err != nil {
				return ODS{}, nil, fmt.Errorf("error decoding file: %v", err)
			}
		case "styles.xml":
			if err = xml.NewDecoder(rc).Decode(&data.Styles); err != nil {
				return ODS{}, nil, fmt.Errorf("error decoding file: %v", err)
			}
		case "mimetype":
			contents, err := io.ReadAll(rc)
			if err != nil {
				return ODS{}, nil, fmt.Errorf("error reading file: %v", err)
			}

			data.Mimetype = Mimetype(contents)
		}
	}

	// Return the decoded content, the zip reader (for potential further use), and any error
	return data, &zip.ReadCloser{Reader: *files}, nil
}
