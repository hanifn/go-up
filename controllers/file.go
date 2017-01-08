package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "fmt"
    "os"
    "bufio"
    "github.com/gorilla/mux"
    "github.com/hanifn/go-up/controllers/utils"
    "strings"
    "strconv"
    "io"
)

type FileController struct {}

func NewFileController() FileController {
    return FileController{}
}

func (fc FileController) Index(w http.ResponseWriter, req *http.Request) {
    files, err := models.GetFiles()
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, files)
}

func (fc FileController) Create(w http.ResponseWriter, req *http.Request) {
    // prep memory for file
    req.ParseMultipartForm(32 << 20)

    file, handler, err := req.FormFile("file")
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, "Error uploading file", 400)
        return
    }
    defer file.Close()

    hash := models.GenerateId();
    fileModel := models.NewFile(
        handler.Filename,
        hash,
        "./storage/"+hash,
        handler.Header.Get("Content-Type"),
        req.FormValue("description"),
    )

    // check if user provided resize params
    resize := req.FormValue("resize")
    if resize != "" {
        // resize image
        dimensions := strings.Split(resize, "x")
        width, err := strconv.Atoi(dimensions[0])
        if err != nil {
            fmt.Printf("%v\n", err)
            utils.JsonError(w, "Error parsing resize dimensions", 400)
            return
        }
        height, err := strconv.Atoi(dimensions[1])
        if err != nil {
            fmt.Printf("%v\n", err)
            utils.JsonError(w, "Error parsing resize dimensions", 400)
            return
        }

        err = fileModel.ResizeImage(file, width, height)
        if err != nil {
            fmt.Printf("%v\n", err)
            utils.JsonError(w, "Error resizing file", 400)
            return
        }
    } else {
        // create new file on filesystem
        f, err := os.OpenFile("./storage/"+hash, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            utils.JsonError(w, err, 400)
            return
        }
        defer f.Close()

        // copy uploaded file to new file
        io.Copy(f, file)
    }

    err = fileModel.Save()
    if err != nil {
        // delete file since saving failed
        os.Remove(fileModel.Path)
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, fileModel)
}

func (fc FileController) Read(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    file, err := models.GetFile(hash)
    if err != nil {
        utils.JsonError(w, err, 404)
        return
    }

    f, err := os.Open(file.Path)
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }
    defer f.Close()

    w.Header().Add("Content-Type", file.Type)

    br := bufio.NewReader(f)
    br.WriteTo(w)
}

// delete method just to meet controller interface
func (fc FileController) Update(w http.ResponseWriter, req *http.Request) {}

func (fc FileController) Delete(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    err := models.DeleteFile(hash)
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, "File deleted")
}
