package models

import (
    "github.com/speps/go-hashids"
)

type File struct {
    Id int `json:"id"`
    HashId string `json:"hash"`
    Name string `json:"name"`
    Path string `json:"path"`
    Type string `json:"type"`
}

type Files []File

func NewFile(name string, path string, fileType string) File {
    // default id to zero before saving to DB
    return File{
        Id: 0,
        HashId: "",
        Name: name,
        Path: path,
        Type: fileType,
    }
}

func (f *File) generateId() {
    hd := hashids.NewData()
    hd.Salt = "$qz&vzp#rwLNe4LV6dr3>o39ei?#Rhud"
    hd.MinLength = 6
    h := hashids.NewWithData(hd)
    id, _ := h.Encode([]int{f.Id})

    f.HashId = id
}


