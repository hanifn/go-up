package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/hanifn/go-up/models"
)

func main() {
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
    file := models.NewFile("File.txt", "/path/to/File.txt", "text")

    json.NewEncoder(w).Encode(file)
}

type Error struct {
    message string
}
