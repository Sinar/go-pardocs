package pardocs

import (
	"testing"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

func Test_getParliamentDocMetadata(t *testing.T) {
	type args struct {
		pdfPath string
		ht      hansard.HansardType
	}
	tests := []struct {
		name            string
		args            args
		wantSessionName string
		wantHansardType string
	}{
		{"test #1", args{"./raw/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf",
			hansard.HANSARD_WRITTEN},
			"Pertanyaan Jawapan Bukan Lisan 22019_new", "BukanLisan"},
		{"test #2", args{"/tmp/Parlimen_Lisan_22019.pdf",
			hansard.HANSARD_SPOKEN}, "Parlimen_Lisan_22019",
			"Lisan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSessionName, gotHansardType := getParliamentDocMetadata(tt.args.pdfPath, tt.args.ht)
			if gotSessionName != tt.wantSessionName {
				t.Errorf("getParliamentDocMetadata() gotSessionName = %v, want %v", gotSessionName, tt.wantSessionName)
			}
			if gotHansardType != tt.wantHansardType {
				t.Errorf("getParliamentDocMetadata() gotHansardType = %v, want %v", gotHansardType, tt.wantHansardType)
			}
		})
	}
}
