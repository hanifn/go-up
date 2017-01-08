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
		desc     string
	}
	tests := []struct {
		name string
		args args
		want File
	}{
		// TODO: Add test cases.
		{"test1", args{"file1", "path/to/file1", "text", "description"}, File{0, "", "file1", "path/to/file1", "text", "description"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFile(tt.args.name, tt.args.path, tt.args.fileType, tt.args.desc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFile(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    File
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFile(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFiles(t *testing.T) {
	tests := []struct {
		name    string
		want    []File
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	type args struct {
		hash string
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
			if err := DeleteFile(tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_generateId(t *testing.T) {
	type fields struct {
		Id          int
		HashId      string
		Name        string
		Path        string
		Type        string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				Id:          tt.fields.Id,
				HashId:      tt.fields.HashId,
				Name:        tt.fields.Name,
				Path:        tt.fields.Path,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
			}
			f.generateId()
		})
	}
}

func TestFile_Save(t *testing.T) {
	type fields struct {
		Id          int
		HashId      string
		Name        string
		Path        string
		Type        string
		Description string
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
			f := &File{
				Id:          tt.fields.Id,
				HashId:      tt.fields.HashId,
				Name:        tt.fields.Name,
				Path:        tt.fields.Path,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
			}
			if err := f.Save(); (err != nil) != tt.wantErr {
				t.Errorf("File.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_update(t *testing.T) {
	type fields struct {
		Id          int
		HashId      string
		Name        string
		Path        string
		Type        string
		Description string
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
			f := &File{
				Id:          tt.fields.Id,
				HashId:      tt.fields.HashId,
				Name:        tt.fields.Name,
				Path:        tt.fields.Path,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
			}
			if err := f.update(); (err != nil) != tt.wantErr {
				t.Errorf("File.update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_insert(t *testing.T) {
	type fields struct {
		Id          int
		HashId      string
		Name        string
		Path        string
		Type        string
		Description string
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
			f := &File{
				Id:          tt.fields.Id,
				HashId:      tt.fields.HashId,
				Name:        tt.fields.Name,
				Path:        tt.fields.Path,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
			}
			if err := f.insert(); (err != nil) != tt.wantErr {
				t.Errorf("File.insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
