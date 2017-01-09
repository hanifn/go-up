package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hanifn/go-up/routes"
    "strings"
)

var (
	server *httptest.Server
	reader io.Reader //Ignore this for now
    url string
)

func init() {
	server = httptest.NewServer(routes.NewRouter(NewFileController()))

    url = server.URL
}

func TestNewFileController(t *testing.T) {
	tests := []struct {
		name string
		want FileController
	}{
        {"New Filecontroller", FileController{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileController_Index(t *testing.T) {
    reader = strings.NewReader("")

    w := httptest.NewRecorder()

    req, err := http.NewRequest("GET", url+"/files", reader)
    if err != nil {
        t.Error(err)
    }

    t.Run("Test Get files", func(t *testing.T) {
        fc := &FileController{}
        fc.Index(w, req)
    })
}

func TestFileController_Create(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		fc   *FileController
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FileController{}
			fc.Create(tt.args.w, tt.args.req)
		})
	}
}

func TestFileController_Read(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		fc   *FileController
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FileController{}
			fc.Read(tt.args.w, tt.args.req)
		})
	}
}

func TestFileController_Update(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		fc   *FileController
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FileController{}
			fc.Update(tt.args.w, tt.args.req)
		})
	}
}

func TestFileController_Delete(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		fc   *FileController
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FileController{}
			fc.Delete(tt.args.w, tt.args.req)
		})
	}
}
