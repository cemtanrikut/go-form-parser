package main

import (
	"fmt"
	"log"

	"main.go/internal/parser"
	"main.go/internal/pdf"
)

func main() {
	form, err := parser.ParseXML("form.xml")
	if err != nil {
		log.Fatalf("Error parsing XML: %s", err)
	}

	err = pdf.GeneratePDF(form, "form_output.pdf")
	if err != nil {
		log.Fatalf("Error generating PDF: %s", err)
	}

	fmt.Println("PDF generated successfully. Check output/form_output.pdf")
}
