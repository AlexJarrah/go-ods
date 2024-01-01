package ods

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
)

// Extracts the spreadsheet data from an ODS file and provides access to the archive.
func Read(filepath string) (content Content, r *zip.ReadCloser, err error) {
	// Open the ODS file as a zip archive
	if r, err = zip.OpenReader(filepath); err != nil {
		return Content{}, nil, fmt.Errorf("error opening ODS file: %v", err)
	}

	// Iterate through the files within the zip archive
	for _, file := range r.File {
		// Find the "content.xml" file, which contains the spreadsheet data
		if file.Name != "content.xml" {
			continue
		}

		// Open the "content.xml" file for processing
		rc, err := file.Open()
		if err != nil {
			return Content{}, nil, fmt.Errorf("error opening content.xml: %v", err)
		}
		defer rc.Close()

		// Decode the XML content into the Content struct
		if err = xml.NewDecoder(rc).Decode(&content); err != nil {
			return Content{}, nil, fmt.Errorf("error decoding content.xml: %v", err)
		}

		// Break the loop, no other files need to be processed
		break
	}

	// Return the decoded content, the zip reader (for potential further use), and any error
	return content, r, nil
}
