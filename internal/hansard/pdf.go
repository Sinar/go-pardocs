package hansard

import (
	"github.com/ledongthuc/pdf"
)

// All functions dealing with PDFs here; the output to be returned and fed into the metadata extraction in main
// hansard routines ..

type PDFPage struct {
	PageNo           int
	PDFTxtSameStyles []string // combined content with same style .. proxy for changes
	PDFPlainText     string
	PDFRawTxt        []pdf.Text // for debugging purposes ..
}

type PDFDocument struct {
	NumPages int
	Pages    []PDFPage
}

func ExtractPDF(pdfPath string) (*PDFDocument, error) {
	// Guard functions here ..

	// Now all c;ear, do the action
	// Start with the zero value
	var pdfDoc PDFDocument

	// iterate through all the pages one by one

	// copy over plain text

	// processStyleChanges ..

	return &pdfDoc, nil
}

func extractTxtSameStyles(pdfContentTxt []pdf.Text) []string {
	var pdfTxtSameStyles []string

	// Guard function .. what is it?

	// Happy path ..

	return pdfTxtSameStyles
}
