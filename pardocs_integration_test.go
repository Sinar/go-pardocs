package pardocs_test

import (
	"testing"

	"github.com/Sinar/go-pardocs"
	"github.com/Sinar/go-pardocs/internal/hansard"
)

func TestParliamentDocs_Plan(t *testing.T) {
	tests := []struct {
		name string
		pd   *pardocs.ParliamentDocs
	}{
		{"test #1", &pardocs.ParliamentDocs{pardocs.Configuration{
			"par14sesi2", hansard.HANSARD_WRITTEN,
			".", "./raw/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf", pardocs.PLAN}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pd.Plan()
		})
	}
}

func TestParliamentDocs_Split(t *testing.T) {
	tests := []struct {
		name string
		pd   *pardocs.ParliamentDocs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pd.Split()
		})
	}
}

func TestParliamentDocs_Reset(t *testing.T) {
	tests := []struct {
		name string
		pd   *pardocs.ParliamentDocs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pd.Reset()
		})
	}
}
