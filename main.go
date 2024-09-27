package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/unidoc/unioffice/document"
)

func main() {
	docPath := "input.docx"
	pdfPath := "output.pdf"

	// Open DOCX file
	doc, err := document.Open(docPath)
	if err != nil {
		fmt.Println("Error opening DOCX:", err)
		return
	}
	defer doc.Close()

	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Read and convert DOCX content to PDF
	for _, para := range doc.Paragraphs() {
		text := para.Text()
		pdf.MultiCell(190, 10, strings.TrimSpace(text), "", "", false)
	}

	// Save the PDF
	err = pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF created successfully:", pdfPath)
}
