package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hanifn/go-up/routes"
    "strings"
    "github.com/hanifn/go-up/models"
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
        {"New Filecontroller", FileController{models.NewFileModel(models.NewConnector("./storage/goup.db"))}},
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
        fm := models.NewFileModel(models.NewConnector("./testdata/test1.db"))
        fc := &FileController{fm}
        fc.Index(w, req)
        if status := w.Code; status != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusOK)
        }

        expected := `[{"hash":"4oks5FjCY","name":"IMG_1932.JPG","description":"Test description"},{"hash":"5hjcurjCY","name":"IMG_1931.JPG","description":"Test description"}]`
        if w.Body.String() != expected {
            t.Errorf("handler returned unexpected body: got %v want %v",
                w.Body.String(), expected)
        }
    })
}

func TestFileController_Create(t *testing.T) {
    reader = strings.NewReader("")

    w := httptest.NewRecorder()

    req, err := http.NewRequest("POST", url+"/upload", reader)
    if err != nil {
        t.Error(err)
    }

    t.Run("Test upload file", func(t *testing.T) {
        fc := &FileController{}
        fc.Create(w, req)

        if status := w.Code; status != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusOK)
        }
    })
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
