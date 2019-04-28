package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/davecgh/go-spew/spew"

	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu/validate"
	"github.com/ledongthuc/pdf"
)

// All extracted from pdfcpu .. da best!
func contentObjNrs(ctx *pdfcpu.Context, page int) ([]int, error) {

	objNrs := []int{}

	d, _, err := ctx.PageDict(page)
	if err != nil {
		return nil, err
	}

	o, found := d.Find("Contents")
	if !found || o == nil {
		return nil, nil
	}

	var objNr int

	ir, ok := o.(pdfcpu.IndirectRef)
	if ok {
		objNr = ir.ObjectNumber.Value()
	}

	o, err = ctx.Dereference(o)
	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	switch o := o.(type) {

	case pdfcpu.StreamDict:

		objNrs = append(objNrs, objNr)

	case pdfcpu.Array:

		for _, o := range o {

			ir, ok := o.(pdfcpu.IndirectRef)
			if !ok {
				return nil, errors.Errorf("missing indref for page tree dict content no page %d", page)
			}

			sd, err := ctx.DereferenceStreamDict(ir)
			if err != nil {
				return nil, err
			}

			if sd == nil {
				continue
			}

			objNrs = append(objNrs, ir.ObjectNumber.Value())

		}

	}

	return objNrs, nil
}

func exploreContent(ctx *pdfcpu.Context, selectedPages pdfcpu.IntSet) error {

	visited := pdfcpu.IntSet{}

	for p, v := range selectedPages {

		fmt.Println("Pages: ", p, v)
		// Page has been chosen for exploration ..
		if v {
			objNrs, err := contentObjNrs(ctx, p)
			if err != nil {
				log.Fatal("context_ERR for page:", p)
				return err
			}

			if objNrs == nil {
				log.Println("objNrs is NIL!!")
				continue
			}

			for _, objNr := range objNrs {

				if visited[objNr] {
					log.Println("VISITED BEFOREE:", objNr)
					continue
				}

				visited[objNr] = true

				b, err := pdfcpu.ExtractStreamData(ctx, objNr)
				if err != nil {
					log.Fatal("EXTRACT_ERR:", err)
					return err
				}

				if b == nil {
					log.Println("Nothing to do with: ", objNr)
					continue
				}

				log.Println(string(b[:]))
			}
		}
	}
	return nil
}

func iteratePDFFail(sourceFileName string) {

	pdfctx, readerr := pdfcpu.ReadFile(sourceFileName, pdfcpu.NewDefaultConfiguration())
	if readerr != nil {
		log.Fatal("ERR:", readerr)
	}

	// Needs to verify first otherwise page count is not in there ..

	valerr := validate.XRefTable(pdfctx.XRefTable)
	if valerr != nil {
		log.Fatal("val_ERR: ", valerr)
	}
	log.Println("Document has ", pdfctx.PageCount, " page(s)")
	pdfref, pgerr := pdfctx.Pages()
	if pgerr != nil {
		log.Fatal("ERR:", pgerr)
	}
	// DEBUG
	spew.Dump(pdfref)
	log.Println("Name:", pdfctx.Read.FileName, " Size:", pdfctx.Read.FileSize)
	// data, exerr := pdfcpu.ExtractContentData(pdfctx, 0)
	//spew.Println(api.ParsePageSelection("1-50"))
	// spew.Dump(pdfcpu.ExtractContentData(pdfctx, 1))

	pageSelection := pdfcpu.IntSet{}
	pageSelection[1] = true
	pageSelection[2] = true

	exerr := exploreContent(pdfctx, pageSelection)
	if exerr != nil {
		log.Fatal("explore_ERR: ", exerr)
	}

	//api.ExtractPages(&api.Command{})
}

func iteratePDF(sourceFileName string) {
	f, r, err := pdf.Open(sourceFileName)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	log.Println("Num pages: ", r.NumPage())
	log.Println("OUTLINE TITLE: ", r.Outline().Title)
	for _, v := range r.Outline().Child {
		spew.Dump(v)
	}
	log.Println("START PAGE 5 ================")
	spew.Dump(r.Page(5).GetPlainText(nil))
	log.Println("END PAGE 5 >>>>>>>>>>>>>>>>>")
	log.Println("TRAILER", r.Trailer().RawString())
	spew.Dump(r.Page(5).Resources().Keys())
	spew.Dump(r.Page(5).Content().Text)
	//_, rerr := r.GetPlainText()
	//if rerr != nil {
	//	panic(rerr)
	//}

	//log.Println(b)

}

func recognizeNewPage(currentPage *pdf.Page) {
	// Look out for keywords pattern of the following combination
	// SOALAN
	// <DIGIT>
	// PEMBERITAHUAN

	// Soalan Number will be within close range to the current page number

}

// SamplePDFPages will extract out source PDF, sample numberOfPages and save it in target location
func SamplePDFPages(sourcePDF string, numberOfPages int, targetPDF string) {

	pdfFile, werr := os.Create(targetPDF)
	defer pdfFile.Close()
	if werr != nil {
		panic(werr)
	}

	f, r, err := pdf.Open(sourcePDF)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//totalPage := r.NumPage()
	//totalPage := 1

	for pageIndex := 1; pageIndex <= numberOfPages; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		// Read all the pages
		//b, ierr := ioutil.ReadAll(p.V.Reader())
		//if ierr != nil {
		//	panic(ierr)
		//}
		//b := []byte(p.Content())
		//// Write it out ..
		//_, werr := pdfFile.Write(b)
		//if werr != nil {
		//	panic(werr)
		//}
	}
}

// SplitBukanLisanPDFs breaks apart ~100 questions into
func SplitBukanLisanPDFs() {
	fmt.Println("Inside SplitBukanLisanPDFs .. ")
	// In the new format, it is much more simplified; all questions during the session ..

	// Break apart full document into a PDF struct for analysis
	// Below is with clean up data below PDF7?
	//iteratePDF("raw/BukanLisan/test_optimized.pdf")
	iteratePDF("raw/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019.pdf")
	//readPdf2("raw/BukanLisan/fixture/test_1-15.pdf")
	//readPdf2(("raw/BukanLisan/clean_new.pdf"))
	//SamplePDFPages("raw/BukanLisan/fixture/test_1-15.pdf", 1, "/tmp/test.pdf")
	// Looks for consecutive Soalan keywords; mark potential split
	// Detect when we have gone too far
	// Re-run for sanity check; point out missing numbers
	// Output structure for plan; can be manipulated; with fancy overlays :P
	// Split based on the planned structure
}

// Example form PR + comments --> https://github.com/rsc/pdf/pull/21/files?short_path=04c6e90#diff-04c6e90faac2675aa89e2176d2eec7d8
func readPdf2(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	//totalPage := r.NumPage()
	totalPage := 5

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			// see if need to add more new line to make it look good
			//if isDifferentLine(text, lastTextStyle) {
			//	text.S = "\n" + text.S
			//}
			if text.S == "\n" {
				fmt.Println("NEWLINE!!!===============")
			} else if text.Y != lastTextStyle.Y {
				fmt.Println("CHange Y!!!")
			}

			if isSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
				lastTextStyle = text
			}
		}
	}

	spew.Dump(r.Page(5).GetPlainText(nil))
	spew.Dump(r.Page(6).GetPlainText(nil))

	log.Println("MULTI-PAGE================>")

	fmt.Println(r.Page(9).GetPlainText(nil))
	fmt.Println(r.Page(10).GetPlainText(nil))
	fmt.Println(r.Page(11).GetPlainText(nil))
	fmt.Println(r.Page(12).GetPlainText(nil))

	return "", nil
}

func isDifferentLine(t1, t2 pdf.Text) bool {
	// if Y axis changes new line else same line
	if t1.Y != t2.Y {
		return true
	}
	return false
}

func isSameSentence(t1, t2 pdf.Text) bool {
	if t1.Font == t2.Font && t1.FontSize == t2.FontSize {
		return true
	}
	return false
}
