package main

import (
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode("Nothing to see here")
}

type Error struct {
    message string
}
