package debate

import (
	"testing"
)

func Test_detectSessionDate(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := detectSessionDate()
			if (err != nil) != tt.wantErr {
				t.Errorf("detectSessionDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("detectSessionDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detectTOCTopicPage(t *testing.T) {
	tests := []struct {
		name        string
		wantTopic   string
		wantPageNum string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTopic, gotPageNum, err := detectTOCTopicPage()
			if (err != nil) != tt.wantErr {
				t.Errorf("detectTOCTopicPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTopic != tt.wantTopic {
				t.Errorf("detectTOCTopicPage() gotTopic = %v, want %v", gotTopic, tt.wantTopic)
			}
			if gotPageNum != tt.wantPageNum {
				t.Errorf("detectTOCTopicPage() gotPageNum = %v, want %v", gotPageNum, tt.wantPageNum)
			}
		})
	}
}

func Test_extractTOC(t *testing.T) {
	type args struct {
		samplePages []PDFPage
		debateTOC   *DebateTOC
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
			if err := extractTOC(tt.args.samplePages, tt.args.debateTOC); (err != nil) != tt.wantErr {
				t.Errorf("extractTOC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hasSessionDateHeader(t *testing.T) {
	type args struct {
		rowContent string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSessionDateHeader(tt.args.rowContent); got != tt.want {
				t.Errorf("hasSessionDateHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeSessionDate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := normalizeSessionDate(); (err != nil) != tt.wantErr {
				t.Errorf("normalizeSessionDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
