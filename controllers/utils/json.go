package utils

import (
    "net/http"
    "encoding/json"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
    // set return headers
    w.Header().Add("Content-Type", "application/json")

    json.NewEncoder(w).Encode(data)
}

func JsonError(w http.ResponseWriter, err interface{}, errCode int) {
    // set return headers
    w.Header().Add("Content-Type", "application/json")

    w.WriteHeader(errCode)

    json.NewEncoder(w).Encode(err)
}
