package controllers

import (
    "github.com/hanifn/go-up/models"
    "net/http"
    "fmt"
    "os"
    "bufio"
    "github.com/gorilla/mux"
    "github.com/hanifn/go-up/controllers/utils"
    "strings"
    "strconv"
    "io"
    "errors"
    "github.com/hanifn/go-up/services"
    "bytes"
    "io/ioutil"
)

const path = "./storage/goup.db"

type FileController struct {
    model models.FileModel
}

func NewFileController() FileController {
    conn := models.NewConnector(path)
    return FileController{models.NewFileModel(conn)}
}

func (fc FileController) Index(w http.ResponseWriter, req *http.Request) {
    files, err := fc.model.GetFiles()
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, files)
}

func (fc FileController) Create(w http.ResponseWriter, req *http.Request) {
    // prep memory for file
    req.ParseMultipartForm(32 << 20)

    file, handler, err := req.FormFile("file")
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, "Error uploading file", 400)
        return
    }
    defer file.Close()

    hash := models.GenerateId();
    fileModel := fc.model.NewFile(
        handler.Filename,
        hash,
        "./storage/"+hash,
        handler.Header.Get("Content-Type"),
        req.FormValue("description"),
    )

    // check if user provided resize params
    resize := req.FormValue("resize")
    if resize != "" {
        // resize image
        err := fc.resizeImage(file, fileModel, resize)
        if err != nil {
            utils.JsonError(w, err, 400)
            return
        }
    } else {
        // create new file on filesystem
        f, err := os.OpenFile("./storage/"+hash, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            utils.JsonError(w, err, 400)
            return
        }
        defer f.Close()

        // copy uploaded file to new file
        io.Copy(f, file)
    }

    // check if upload to s3
    upsertS3 := req.FormValue("upsert")
    if upsertS3 != "" {
        // open file as bytes
        data, err := ioutil.ReadFile(fileModel.Path)
        if err != nil {
            utils.JsonError(w, err, 400)
            return
        }

        err = services.UpsertToS3(fileModel.HashId, data, fileModel.Type)
        if err != nil {
            utils.JsonError(w, err, 400)
            return
        }

        // set s3 flag
        fileModel.AwsS3 = true
        os.Remove(fileModel.Path)
    }

    err = fileModel.Save()
    if err != nil {
        // delete file since saving failed
        os.Remove(fileModel.Path)
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, fileModel)
}

func (fc FileController) resizeImage(file io.Reader, fileModel models.File, resize string) error {
    dimensions := strings.Split(resize, "x")
    width, err := strconv.Atoi(dimensions[0])
    if err != nil {
        fmt.Printf("%v\n", err)
        return errors.New("Error parsing resize dimensions")
    }
    height, err := strconv.Atoi(dimensions[1])
    if err != nil {
        fmt.Printf("%v\n", err)
        return errors.New("Error parsing resize dimensions")
    }

    err = fileModel.ResizeImage(file, width, height)
    if err != nil {
        fmt.Printf("%v\n", err)
        return errors.New("Error resizing file")
    }

    return nil
}

func (fc FileController) Read(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    file, err := fc.model.GetFile(hash)
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, err, 404)
        return
    }

    w.Header().Add("Content-Type", file.Type)

    var br *bufio.Reader
    if file.AwsS3 {
        // get from S3
        data, err := services.DownloadFromS3(file.HashId)
        if err != nil {
            fmt.Printf("%v\n", err)
            utils.JsonError(w, err, 400)
            return
        }

        f := bytes.NewBuffer(data)
        br = bufio.NewReader(f)
    } else {
        // get from file
        f, err := os.Open(file.Path)
        if err != nil {
            fmt.Printf("%v\n", err)
            utils.JsonError(w, err, 400)
            return
        }
        defer f.Close()

        br = bufio.NewReader(f)
    }

    br.WriteTo(w)
}

// delete method just to meet controller interface
func (fc FileController) Update(w http.ResponseWriter, req *http.Request) {}

func (fc FileController) Delete(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    hash := vars["id"]

    file, err := fc.model.GetFile(hash)
    if err != nil {
        fmt.Printf("%v\n", err)
        utils.JsonError(w, err, 404)
        return
    }

    if file.AwsS3 {
        err = services.RemoveFromS3(hash)
    }

    err = fc.model.DeleteFile(hash)
    if err != nil {
        utils.JsonError(w, err, 400)
        return
    }

    utils.JsonResponse(w, "File deleted")
}
