package models

import (
	"reflect"
	"testing"
)

func TestNewFile(t *testing.T) {
	type args struct {
		name     string
		path     string
		fileType string
	}
	tests := []struct {
		name string
		args args
		want File
	}{
	// TODO: Add test cases.
		{"test1", args{"file1", "path/to/file1", "text"}, File{"11jhg13", "file1", "path/to/file1", "text"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFile(tt.args.name, tt.args.path, tt.args.fileType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
