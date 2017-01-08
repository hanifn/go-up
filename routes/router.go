package routes

import (
    "github.com/gorilla/mux"
    "net/http"
)

type Controller interface {
    Index(w http.ResponseWriter, req *http.Request)
    Create(w http.ResponseWriter, req *http.Request)
    Read(w http.ResponseWriter, req *http.Request)
    Update(w http.ResponseWriter, req *http.Request)
    Delete(w http.ResponseWriter, req *http.Request)
}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

func NewRouter(controller Controller) *mux.Router {

    router := mux.NewRouter().StrictSlash(true)

    routes := getRoutes(controller)

    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

func getRoutes(controller Controller) []Route {

    return []Route{
        {
            "Index",
            "GET",
            "/",
            controller.Index,
        },
        {
            "GetFiles",
            "GET",
            "/files",
            controller.Index,
        },
        {
            "GetFile",
            "GET",
            "/files/{id}",
            controller.Read,
        },
        {
            "UploadFile",
            "POST",
            "/upload",
            controller.Create,
        },
        {
            "DeleteFile",
            "DELETE",
            "/files/{id}",
            controller.Delete,
        },
    }
}
