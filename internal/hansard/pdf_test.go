package hansard

import (
	"testing"

	"github.com/ledongthuc/pdf"
)

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
