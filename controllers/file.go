package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "encoding/json"
    "fmt"
    "os"
    "io"
    "bufio"
)

func GetFile(w http.ResponseWriter, req *http.Request) {
    file := models.NewFile("Test", "/path/to/file", "text")

    json.NewEncoder(w).Encode(file)
}

func Upload(w http.ResponseWriter, req *http.Request) {
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

    json.NewEncoder(w).Encode(fileModel)
}

func Download(w http.ResponseWriter, req *http.Request) {
    f, err := os.Open("./storage/IMG_1928.JPG")

    if err != nil {
        json.NewEncoder(w).Encode(err)
        return
    }
    defer f.Close()

    w.Header().Add("Content-Type", "image/jpeg")

    br := bufio.NewReader(f)
    br.WriteTo(w)
}
