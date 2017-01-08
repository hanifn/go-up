package main

import (
    "log"
    "net/http"
    "github.com/hanifn/go-up/routes"
    "github.com/hanifn/go-up/controllers"
)

func main() {
    router := routes.NewRouter(controllers.NewFileController())

    log.Fatal(http.ListenAndServe(":8000", router))
}

type Error struct {
    message string
}

