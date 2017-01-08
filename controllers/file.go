package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "fmt"
    "os"
    "io"
    "bufio"
    "github.com/gorilla/mux"
    "github.com/hanifn/go-up/controllers/utils"
)

func GetFiles(w http.ResponseWriter, req *http.Request) {
    files, err := models.GetFiles()
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, files)
}

func Upload(w http.ResponseWriter, req *http.Request) {
    // prep memory for file
    req.ParseMultipartForm(32 << 20)

    file, handler, err := req.FormFile("file")
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, "Error uploading file", 400)
        return
    }
    defer file.Close()

    // create new file on filesystem
    f, err := os.OpenFile("./storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }
    defer f.Close()

    // copy uploaded file to new file
    io.Copy(f, file)

    fileModel := models.NewFile(
        handler.Filename,
        "./storage/"+handler.Filename,
        handler.Header.Get("Content-Type"),
        req.FormValue("description"),
    )

    err = fileModel.Save()
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, fileModel)
}

func Download(w http.ResponseWriter, req *http.Request) {
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

func Delete(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    err := models.DeleteFile(hash)
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, "File deleted")
}
