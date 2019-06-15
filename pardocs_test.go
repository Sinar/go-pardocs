package pardocs

import (
	"testing"
)

func TestParliamentDocs_Plan(t *testing.T) {
	type fields struct {
		Conf Configuration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &ParliamentDocs{
				Conf: tt.fields.Conf,
			}
			pd.Plan()
		})
	}
}

func TestParliamentDocs_Split(t *testing.T) {
	type fields struct {
		Conf Configuration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &ParliamentDocs{
				Conf: tt.fields.Conf,
			}
			pd.Split()
		})
	}
}

func TestParliamentDocs_Reset(t *testing.T) {
	type fields struct {
		Conf Configuration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &ParliamentDocs{
				Conf: tt.fields.Conf,
			}
			pd.Reset()
		})
	}
}
