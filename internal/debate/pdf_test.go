package debate

import (
	"reflect"
	"testing"

	"github.com/ledongthuc/pdf"
)

func TestNewPDFDoc(t *testing.T) {
	type args struct {
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *PDFDocument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPDFDoc(tt.args.sourcePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPDFDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPDFDoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPDFDocument_extractPDF(t *testing.T) {
	type fields struct {
		NumPages   int
		Pages      []PDFPage
		sourcePath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdfDoc := &PDFDocument{
				NumPages:   tt.fields.NumPages,
				Pages:      tt.fields.Pages,
				sourcePath: tt.fields.sourcePath,
			}
			if err := pdfDoc.extractPDF(); (err != nil) != tt.wantErr {
				t.Errorf("extractPDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_extractTxtSameLine(t *testing.T) {
	type args struct {
		ptrTxtSameLine *[]string
		pdfContentTxt  []pdf.Text
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := extractTxtSameLine(tt.args.ptrTxtSameLine, tt.args.pdfContentTxt); (err != nil) != tt.wantErr {
				t.Errorf("extractTxtSameLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_extractTxtSameStyles(t *testing.T) {
	type args struct {
		ptrTxtSameStyles *[]string
		pdfContentTxt    []pdf.Text
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := extractTxtSameStyles(tt.args.ptrTxtSameStyles, tt.args.pdfContentTxt); (err != nil) != tt.wantErr {
				t.Errorf("extractTxtSameStyles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
