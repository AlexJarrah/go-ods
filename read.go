package ods

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

// Extracts the spreadsheet data from an ODS file and provides access to the archive.
func Read(filepath string) (data ODS, files *zip.ReadCloser, err error) {
	// Open the ODS file as a zip archive
	if files, err = zip.OpenReader(filepath); err != nil {
		return ODS{}, nil, fmt.Errorf("error opening ODS file: %v", err)
	}

	// Iterate through the files within the zip archive
	for _, file := range files.File {
		// Open the file for processing
		rc, err := file.Open()
		if err != nil {
			return ODS{}, nil, fmt.Errorf("error opening file: %v", err)
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
	return data, files, nil
}
