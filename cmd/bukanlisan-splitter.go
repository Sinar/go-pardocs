package cmd

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/davecgh/go-spew/spew"

	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu/validate"
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

}

// SplitBukanLisanPDFs breaks apart ~100 questions into
func SplitBukanLisanPDFs() {
	fmt.Println("Inside SplitBukanLisanPDFs .. ")
	// Break apart full document into a PDF struct for analysis
	iteratePDF("raw/BukanLisan/test_optimized.pdf")
	// Looks for consecutive Soalan keywords; mark potential split
	// Detect when we have gone too far
	// Re-run for sanity check; point out missing numbers
	// Output structure for plan; can be manipulated; with fancy overlays :P
	// Split based on the planned structure
}
