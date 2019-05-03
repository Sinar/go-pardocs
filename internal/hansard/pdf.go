package hansard

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/ledongthuc/pdf"
	"golang.org/x/xerrors"
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
	fmt.Println("In ExtractPDF ...")

	// Guard functions here ..

	// Now all c;ear, do the action
	// Start with the zero value
	var pdfDoc PDFDocument
	var pdfPages []PDFPage

	// Example form PR + comments --> https://github.com/rsc/pdf/pull/21/files?short_path=04c6e90#diff-04c6e90faac2675aa89e2176d2eec7d8
	f, r, err := pdf.Open(pdfPath)
	defer f.Close()
	if err != nil {
		return nil, xerrors.Errorf("Open failed: %s -  %w", pdfPath, err)
	}
	// iterate through all the pages one by one
	pdfDoc.NumPages = r.NumPage()

	for i := 1; i <= 5; i++ {
		// init
		pdfPage := PDFPage{}
		pdfPage.PageNo = i

		// Get details for the page
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}
		// copy over plain text; short form
		pt, pterr := p.GetPlainText(nil)
		if pterr != nil {
			return nil, xerrors.Errorf(" GetPlainText ERROR: %w", pt)
		}
		pdfPage.PDFPlainText = pt
		// processStyleChanges ..
		//extractTxtSameStyles()
		// DEBUG
		fmt.Println("LEN: ", p.V.Len())
		fmt.Println("KEYS", p.V.Keys())
		fmt.Println("KIND", p.V.Kind())
		extractTxtSameLine(p.Content().Text)
		pdfPages = append(pdfPages, pdfPage)
	}

	spew.Dump(pdfPages)
	//spew.Dump("BOB \n SUE \n MARY ....")

	return &pdfDoc, nil
}

func extractTxtSameLine(pdfContentTxt []pdf.Text) []string {
	var pdfTxtSameLine []string

	var numValidLineCounted int
	var currentLineNumber float64
	var currentContent string

	//spew.Dump(pdfContentTxt)

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?
		//if strings.TrimSpace(v.S) == "" {
		//	fmt.Println("Skipping blank line / content ..")
		//	continue
		//}

		if currentLineNumber == 0 {
			currentLineNumber = v.Y
			fmt.Println("Set first line to ", currentLineNumber)
			currentContent += v.S
			continue
		}

		// Happy path ..
		fmt.Println("Current CONTENT: ", currentContent, " X: ", v.X, " Y: ", v.Y)
		// number of valid line increase when new valid line ..
		if currentLineNumber != v.Y {
			fmt.Println("NEW Line ... collected: ", currentContent)
			pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
			numValidLineCounted++
			currentContent = v.S // reset .. after append
			currentLineNumber = v.Y
		} else {
			// If on the same line, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get 10 lines ..
		if numValidLineCounted > 10 {
			break
		}

		// Failsafe
		//if i > 50 {
		//	break
		//}
	}

	return pdfTxtSameLine
}

func extractTxtSameStyles(pdfContentTxt []pdf.Text) []string {
	var pdfTxtSameStyles []string

	// NOTE: Only get 10 lines ..

	// Guard function .. what is it?

	// Happy path ..

	return pdfTxtSameStyles
}
