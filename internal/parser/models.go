package parser

import "encoding/xml"

// Form represents the root of the XML form
type Form struct {
	XMLName  xml.Name  `xml:"Form"`
	Fields   []Field   `xml:"Field"`
	Sections []Section `xml:"Section"`
}

// Field represents an individual form field
type Field struct {
	Name      string `xml:"Name,attr"`
	Type      string `xml:"Type,attr"`
	Optional  string `xml:"Optional,attr"`
	FieldType string `xml:"FieldType,attr"`
	Caption   string `xml:"Caption"`
}

// Section represents a section in the form
type Section struct {
	Name     string  `xml:"Name,attr"`
	Optional string  `xml:"Optional,attr"`
	Title    string  `xml:"Title"`
	Contents []Field `xml:"Contents>Field"`
}
