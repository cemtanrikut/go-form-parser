package parser

import (
	"strings"
	"testing"
)

// Test case: Geçerli XML verisini parse etme
func TestParseXMLValid(t *testing.T) {
	xmlData := `
	<Form>
		<Field Name="language" Type="String" Optional="False" FieldType="TextBox">
			<Caption>Favorite Language</Caption>
		</Field>
	</Form>`

	// StringReader ile XML simüle ediyoruz
	form, err := parseXMLFromReader(strings.NewReader(xmlData))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Beklenen değerleri kontrol et
	if len(form.Fields) != 1 || form.Fields[0].Name != "language" {
		t.Errorf("XML parsing failed. Got: %+v", form)
	}
}

// Test case: Bozuk XML verisi
func TestParseXMLInvalid(t *testing.T) {
	brokenXML := `<Form><Field Name="test"`

	_, err := parseXMLFromReader(strings.NewReader(brokenXML))
	if err == nil {
		t.Fatal("Expected error for invalid XML but got none")
	}
}
