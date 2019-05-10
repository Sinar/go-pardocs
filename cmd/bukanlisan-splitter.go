package cmd

import (
	"fmt"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

func SplitBukanLisanPDFs() {
	fmt.Println("In SplitBukanLisanPDFs ...")
	// Break apart full document into a PDF struct for analysis
	// ./raw/BukanLisan/split/Pertanyaan Jawapan Bukan Lisan 22019_76-90.pdf
	// is a split from the modified original below:
	// $ file ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf
	// 		/Users/mleow/Downloads/Pertanyaan Jawapan Bukan Lisan 22019.pdf: PDF document, version 1.7
	//pdfDoc := hansard.PDFDocument{}
	pdfPath := "./raw/BukanLisan/split/Pertanyaan Jawapan Bukan Lisan 22019_76-90.pdf"
	pdfDoc, err := hansard.NewPDFDoc(pdfPath)
	if err != nil {
		panic(err)
	}
	// Extract the rest of the data; or should it just be built in?
	// TODO: Probably ..
	exerr := pdfDoc.ExtractPDF()
	if exerr != nil {
		panic(exerr)
	}

	// Start Hansard DOcument
	// TODO: Refactor into a factory perhaps?
	hansardDoc, _ := hansard.NewHansardDocument(pdfPath)
	// Start SplitterState; refactored inside HansardDocument; is implementation details
	//splitterState := hansard.SplitterState{}

	for _, p := range pdfDoc.Pages {
		fmt.Println("PAGE:", p.PageNo)
		// try to recognize value from top of content ..
		// DEBUG
		//for _, c := range p.PDFTxtSameLines {
		//	fmt.Println("FOR CONSIDERATION: ", c)
		//}

		// Detect question
		dterr := hansardDoc.ProcessLinesExcerpt(p.PDFTxtSameLines)
		if dterr != nil {
			panic(dterr)
		}
		// Looks for consecutive Soalan keywords; mark potential split
		// Detect when we have gone too far; start a new Question with the data ..

	}
	// Re-run for sanity check; point out missing numbers
	// Output structure for plan; can be manipulated; with fancy overlays :P
	hansardDoc.String()
	// Split based on the planned structure
	hansardDoc.SplitPDFByQuestions()

}
