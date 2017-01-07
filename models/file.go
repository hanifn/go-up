package models

type File struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Path string `json:"path"`
    Type string `json:"type"`
}

type Files []File

func NewFile(name string, path string, fileType string) File {
    return File{
        "11jhg13",
        name,
        path,
        fileType,
    }
}


