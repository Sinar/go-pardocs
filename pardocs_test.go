package pardocs_test

import (
	"testing"

	"github.com/Sinar/go-pardocs"
)

func TestParliamentDocs_Plan(t *testing.T) {
	tests := []struct {
		name string
		pd   *pardocs.ParliamentDocs
	}{
		// TODO: Add test cases.
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
