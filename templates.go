package ods

import _ "embed"

// minimalSettingsXML is a ODS settings.xml template.
//
//go:embed templates/settings.xml
var minimalSettingsXML []byte

// minimalStylesXML is a ODS styles.xml template.
//
//go:embed templates/styles.xml
var minimalStylesXML []byte
