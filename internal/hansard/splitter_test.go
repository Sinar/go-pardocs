package hansard

import (
	"reflect"
	"testing"
)

func Test_detectSessionName(t *testing.T) {
	type args struct {
		hansardType        HansardType
		sourcePDFFileName  string
		contentofFirstPage []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectSessionName(tt.args.hansardType, tt.args.sourcePDFFileName, tt.args.contentofFirstPage); got != tt.want {
				t.Errorf("detectSessionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractCoverPage(t *testing.T) {
	type args struct {
		hansardType     HansardType
		originalPDFPath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractCoverPage(tt.args.hansardType, tt.args.originalPDFPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractCoverPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
