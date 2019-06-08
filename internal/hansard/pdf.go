package hansard

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
	"golang.org/x/xerrors"
)

// All functions dealing with PDFs here; the output to be returned and fed into the metadata extraction in main
// hansard routines ..

type PDFPage struct {
	PageNo           int
	PDFPlainText     string
	PDFTxtSameLines  []string // combined content with same line .. proxy for changes
	PDFTxtSameStyles []string // combined content with same style .. proxy for changes
}

type PDFDocument struct {
	NumPages   int
	Pages      []PDFPage
	sourcePath string
}

const (
	MaxLineProcessed = 7
)

func NewPDFDoc(sourcePath string) (*PDFDocument, error) {
	// TODO: Guard checks to ensure file exists?? etc.
	// and is readbale??

	pdfDoc := PDFDocument{
		sourcePath: sourcePath,
	}

	exerr := pdfDoc.extractPDF()
	if exerr != nil {
		panic(exerr)
	}
	return &pdfDoc, nil

}
func (pdfDoc *PDFDocument) extractPDF() error {
	fmt.Println("In ExtractPDF ...")

	// Guard functions here ..

	// Now all c;ear, do the action
	// Start with the zero value
	//var pdfDoc PDFDocument
	var pdfPages []PDFPage

	// Example form PR + comments --> https://github.com/rsc/pdf/pull/21/files?short_path=04c6e90#diff-04c6e90faac2675aa89e2176d2eec7d8
	f, r, err := pdf.Open(pdfDoc.sourcePath)
	defer f.Close()
	if err != nil {
		return xerrors.Errorf("Open failed: %s -  %w", pdfDoc.sourcePath, err)
	}
	// iterate through all the pages one by one
	pdfDoc.NumPages = r.NumPage()
	// DEBUG
	//pdfDoc.NumPages = 7
	for i := 1; i <= pdfDoc.NumPages; i++ {
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
			return xerrors.Errorf(" GetPlainText ERROR: %w", pt)
		}
		pdfPage.PDFPlainText = pt
		// processStyleChanges ..
		//extractTxtSameStyles()
		// DEBUG
		//fmt.Println("LEN: ", p.V.Len())
		//fmt.Println("KEYS", p.V.Keys())
		//fmt.Println("KIND", p.V.Kind())
		// DEBUG
		//fmt.Println("== START CONTENT PAGE ", i)
		//spew.Dump(pt)
		// Top 10 lines for this page by line analysis
		//fmt.Println("== START ANALYZE by LINE")
		pdfPage.PDFTxtSameLines = make([]string, 0, 20)
		extractTxtSameLine(&pdfPage.PDFTxtSameLines, p.Content().Text)

		// Top 10
		//fmt.Println("== START ANALYZE by STYLE")
		pdfPage.PDFTxtSameStyles = make([]string, 0, 20)
		extractTxtSameStyles(&pdfPage.PDFTxtSameStyles, p.Content().Text)
		//fmt.Println("== END ANALYZE by STYLE")

		pdfPages = append(pdfPages, pdfPage)
	}

	//spew.Dump(pdfPages)
	//spew.Dump("BOB \n SUE \n MARY ....")
	pdfDoc.Pages = pdfPages

	return nil
}

func extractTxtSameLine(ptrTxtSameLine *[]string, pdfContentTxt []pdf.Text) error {

	var numValidLineCounted int
	var currentLineNumber float64
	var currentContent string

	var pdfTxtSameLine []string

	// DEBUG
	//spew.Dump(pdfContentTxt)

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?
		//if strings.TrimSpace(v.S) == "" {
		//	fmt.Println("Skipping blank line / content ..")
		//	continue
		//}

		if currentLineNumber == 0 {
			currentLineNumber = v.Y
			// DEBUG
			//fmt.Println("Set first line to ", currentLineNumber)
			currentContent += v.S
			continue
		}

		// Happy path ..
		// DEBUG
		//fmt.Println("Append CONTENT: ", currentContent, " X: ", v.X, " Y: ", v.Y)
		// number of valid line increase when new valid line ..
		if currentLineNumber != v.Y {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Line ... collected: ", currentContent)
				pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
				numValidLineCounted++
			}
			currentContent = v.S // reset .. after append
			currentLineNumber = v.Y
		} else {
			// If on the same line, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		if numValidLineCounted > MaxLineProcessed {
			break
		}

	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Line ... collected: ", currentContent)
		pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
	}

	*ptrTxtSameLine = pdfTxtSameLine
	//spew.Dump(ptrTxtSameLine)
	return nil
}

func extractTxtSameStyles(ptrTxtSameStyles *[]string, pdfContentTxt []pdf.Text) error {
	var numValidLineCounted int
	var currentFont string
	var currentContent string

	var pdfTxtSameStyles []string

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?

		if currentFont == "" {
			currentFont = v.Font
			// DEBUG
			//fmt.Println("Set first font to ", currentFont)
			currentContent += v.S
			continue
		}

		// Happy path ..
		if currentFont != v.Font {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Style ... collected: ", currentContent)
				pdfTxtSameStyles = append(pdfTxtSameStyles, currentContent)
				//fmt.Println("CURRENT ,...")
				//spew.Dump(pdfTxtSameStyles)
				numValidLineCounted++
			}
			// reset for next iteraton ..
			currentContent = v.S // reset .. after append
			currentFont = v.Font
		} else {
			// If with the same style, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		if numValidLineCounted > MaxLineProcessed {
			break
		}
	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Style ... collected: ", currentContent)
		pdfTxtSameStyles = append(pdfTxtSameStyles, currentContent)
	}

	*ptrTxtSameStyles = pdfTxtSameStyles
	//spew.Dump(ptrTxtSameStyles)

	return nil
}
