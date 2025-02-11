package pdf

import (
	"os"
	"testing"

	"main.go/internal/parser"
)

// PDF oluşturma fonksiyonunu test etme
func TestGeneratePDF(t *testing.T) {
	form := &parser.Form{
		Fields: []parser.Field{
			{Name: "language", Type: "String", Caption: "Favorite Language"},
		},
	}

	outputFile := "output/test_output.pdf"

	err := GeneratePDF(form, outputFile)
	if err != nil {
		t.Fatalf("Failed to generate PDF: %v", err)
	}

	// Dosya gerçekten oluşturulmuş mu?
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatal("PDF file was not created")
	}

	// Test tamamlandıktan sonra test dosyasını silebiliriz
	os.Remove(outputFile)
}
