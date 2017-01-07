package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "encoding/json"
    "fmt"
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

    fmt.Printf("%v\n", handler)
}
