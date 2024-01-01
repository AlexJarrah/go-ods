package ods

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
)

func Read(filepath string) (content Content, r *zip.ReadCloser, err error) {
	if r, err = zip.OpenReader(filepath); err != nil {
		return Content{}, nil, fmt.Errorf("error opening ODS file: %v", err)
	}

	for _, file := range r.File {
		if file.Name != "content.xml" {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return Content{}, nil, fmt.Errorf("error opening content.xml: %v", err)
		}
		defer rc.Close()

		if err = xml.NewDecoder(rc).Decode(&content); err != nil {
			return Content{}, nil, fmt.Errorf("error decoding content.xml: %v", err)
		}
		break
	}

	return content, r, nil
}
