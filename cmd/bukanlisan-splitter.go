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

	//pdfPath := "./raw/BukanLisan/split/Pertanyaan Jawapan Bukan Lisan 22019_76-90.pdf"
	//pdfPath := "./raw/JawatanKuasa/rumusan-laporan-akhir-jawatankuasa-siasatan-tadbir-urus-perolehan-dan-kewangan-kerajaan-mengenai-projek-land-swap-di-bawah-kementerian-pertahanan.pdf"
	// Test some  BukanLisan
	//pdfPath := "./raw/BukanLisan"

	// Test some Lisan throughout the years ..
	pdfPath := "./raw/Lisan/"
	//pdfPath += "20140327__DR_JawabLisan.pdf" // <-- Good test case which causes panic in the pdf lib itself!

	//pdfPath += "20140327__DR_JawabLisan_clean.pdf" //  <-- A save and using OSX Preview resolves the problem!
	// Above has test cases:
	// QUESTION:  2013  START:  70  END:  70 <-- possibly sanity check that question is highly likely next running number
	// QUESTION:    START:  199  END:  201 <-- Finalizer cannot attach to the last remaining question

	//pdfPath += "JWP DR 031218.pdf" // <--  ALso have bad ; no codespace ..
	//pdfPath += "JDR25032019.pdf" // <-- Good test case for bad  rune; no codespace
	// Above has  test cases:
	//  Rune that fails regexp - "SOALAN� �NO.� �1� "
	// Edge case off by one question

	//pdfPath += "JDR12032019.pdf" // <-- Looks ok on first draft
	//pdfPath += "JWP DR 161018.pdf" // <-- Looks OK

	//pdfPath += "JWP DR 151018.pdf" // Another case  of partially bad PDF
	pdfPath += "JWP DR 151018_clean.pdf" // Clean up will solve the problem

	pdfDoc, err := hansard.NewPDFDoc(pdfPath)
	if err != nil {
		panic(err)
	}
	// Extract the rest of the data; or should it just be built in?
	// TODO: Probably ..
	//exerr := pdfDoc.ExtractPDF()
	//if exerr != nil {
	//	panic(exerr)
	//}

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
		dterr := hansardDoc.ProcessLinesExcerpt(p.PageNo, p.PDFTxtSameLines)
		if dterr != nil {
			panic(dterr)
		}
		// Looks for consecutive Soalan keywords; mark potential split
		// Detect when we have gone too far; start a new Question with the data ..

	}
	// Re-run for sanity check; point out missing numbers
	// Output structure for plan; can be manipulated; with fancy overlays :P
	//hansardDoc.String()

	hansardDoc.Finalize()
	// DEBUG:
	// hansardDoc.ShowState()
	fmt.Println("=== QUESTIONS ****************")
	//hansardDoc.ShowQuestions()

	// Split based on the planned structure
	hansardDoc.SplitPDFByQuestions()
	// TODO: Needs a wrap up state for the last state left ..
	// After the split; we should have the HansardQuestions
	//hansardDoc.String()

}
