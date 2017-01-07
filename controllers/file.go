package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "encoding/json"
    "fmt"
    "os"
    "io"
    "bufio"
    "github.com/gorilla/mux"
)

func GetFiles(w http.ResponseWriter, req *http.Request) {
    // set return headers
    w.Header().Add("Content-Type", "application/json")

    files, err := models.GetFiles()
    if err != nil {
        w.WriteHeader(400)
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode(files)
}

func Upload(w http.ResponseWriter, req *http.Request) {
    // set return headers
    w.Header().Add("Content-Type", "application/json")

    req.ParseMultipartForm(32 << 20)
    file, handler, err := req.FormFile("file")

    if err != nil {
        fmt.Printf("%v\n", err)
        w.WriteHeader(404)
        json.NewEncoder(w).Encode("Error uploading file")
        return
    }
    defer file.Close()

    // create new file on filesystem
    f, err := os.OpenFile("./storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(err)
        return
    }
    defer f.Close()

    // copy uploaded file to new file
    io.Copy(f, file)

    fileModel := models.NewFile(
        handler.Filename,
        "./storage/"+handler.Filename,
        handler.Header.Get("Content-Type"),
    )

    err = fileModel.Save()
    if err != nil {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode(fileModel)
}

func Download(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    file, err := models.GetFile(hash)
    if err != nil {
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(err)
        return
    }

    f, err := os.Open(file.Path)
    if err != nil {
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(err)
        return
    }
    defer f.Close()

    w.Header().Add("Content-Type", file.Type)

    br := bufio.NewReader(f)
    br.WriteTo(w)
}

func Delete(w http.ResponseWriter, req *http.Request) {
    // set return headers
    w.Header().Add("Content-Type", "application/json")

    vars := mux.Vars(req)

    hash := vars["id"]

    err := models.DeleteFile(hash)
    if err != nil {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode("File deleted")
}
