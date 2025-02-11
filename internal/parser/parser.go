package parser

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
)

func ParseXML(filename string) (*Form, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var form Form
	err = xml.Unmarshal(bytes, &form)
	if err != nil {
		return nil, err
	}

	return &form, nil
}

func parseXMLFromReader(reader io.Reader) (*Form, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var form Form
	err = xml.Unmarshal(bytes, &form)
	if err != nil {
		return nil, err
	}

	return &form, nil
}
