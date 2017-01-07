package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/hanifn/go-up/controllers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GetFiles",
        "GET",
        "/files",
        controllers.GetFiles,
    },
    Route{
        "GetFile",
        "GET",
        "/files/{id}",
        controllers.Download,
    },
    Route{
        "UploadFile",
        "POST",
        "/upload",
        controllers.Upload,
    },
    Route{
        "DeleteFile",
        "DELETE",
        "/files/{id}",
        controllers.Delete,
    },
}
