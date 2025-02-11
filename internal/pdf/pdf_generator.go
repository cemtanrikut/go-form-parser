package pdf

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
	"main.go/internal/parser"
)

// GeneratePDF takes a parsed Form and creates a PDF representation
func GeneratePDF(form *parser.Form, outputFilename string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Form Submission")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	// Process fields
	for _, field := range form.Fields {
		pdf.Cell(40, 10, fmt.Sprintf("%s: %s", field.Caption, field.Type))
		pdf.Ln(8)
	}

	// Process sections
	for _, section := range form.Sections {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(40, 10, section.Title)
		pdf.Ln(8)
		pdf.SetFont("Arial", "", 12)
		for _, field := range section.Contents {
			pdf.Cell(40, 10, fmt.Sprintf("%s: %s", field.Caption, field.Type))
			pdf.Ln(8)
		}
	}
	return pdf.OutputFileAndClose(outputFilename)
}
