package models

import (
    "fmt"
    "github.com/nfnt/resize"
    "image/jpeg"
    "image/png"
    "os"
    "errors"
    "github.com/ventu-io/go-shortid"
    "io"
)

type File struct {
    Id int `json:"-"`
    HashId string `json:"hash"`
    Name string `json:"name"`
    Path string `json:"-"`
    Type string `json:"-"`
    Description string `json:"description"`
}

type Files []File

func NewFile(name string, hash string, path string, fileType string, desc string) File {
    // default id to zero before saving to DB
    return File{
        Id: 0,
        HashId: hash,
        Name: name,
        Path: path,
        Type: fileType,
        Description: desc,
    }
}

func GetFile(hash string) (File, error) {
    conn := initDB()

    sql := `
    SELECT * FROM files
    WHERE hash = ?
    LIMIT 1
    `

    row := conn.QueryRow(sql, hash)

    var file File
    err := row.Scan(&file.Id, &file.HashId, &file.Name, &file.Path, &file.Type, &file.Description)
    if err != nil {
        return File{}, err
    }

    return file, nil
}


func GetFiles() ([]File, error) {
    conn := initDB()

    sql := `
    SELECT * FROM files
    `

    rows, err := conn.Query(sql)
    if err != nil {
        return []File{}, err
    }

    var files []File
    for rows.Next() {
        file := File{}
        err := rows.Scan(&file.Id, &file.HashId, &file.Name, &file.Path, &file.Type, &file.Description)
        if err != nil {
            return files, err
        }

        files = append(files, file)
    }

    return files, nil
}

func DeleteFile(hash string) error {
    conn := initDB()

    sql := `
    DELETE FROM files
    WHERE hash = ?
    `

    row := conn.QueryRow(sql, hash)
    err := row.Scan()
    if err == nil {
        fmt.Printf("Delete error: %v\n", err)
        return errors.New("No file deleted")
    }

    return nil
}

func GenerateId() string {
    sid, err := shortid.New(1, shortid.DefaultABC, 4532652356)
    if err != nil {
        panic(err)
    }

    return sid.MustGenerate()
}

func (f *File) Save() error {
    // If ID is 0 then call insert method to create new entry
    // else call update method to update existing entry
    if f.Id == 0 {
        err := f.insert()
        if err != nil {
            return err
        }

        fmt.Printf("%v\n", f)
    } else {
        err := f.update()
        if err != nil {
            return err
        }
    }

    return nil
}
func (f *File) update() error {
    conn := initDB()

    sql := `
    UPDATE files SET
        hash = ?,
        name = ?,
        path = ?,
        type = ?
    WHERE id = ?
    `

    stmt, err := conn.Prepare(sql)
    if err != nil {
        fmt.Printf("Update prepare error: %v\n", err)
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(f.HashId, f.Name, f.Path, f.Type, f.Id)
    if err != nil {
        fmt.Printf("Update exec error: %v\n", err)
        return err
    }

    return nil
}

func (f *File) insert() error {
    conn := initDB()

    sql := `
	INSERT INTO files(
	    hash,
		name,
		path,
		type,
		description
	) values(?, ?, ?, ?, ?)
	`

    stmt, err := conn.Prepare(sql)
    if err != nil {
        fmt.Printf("Insert prepare error: %v\n", err)
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(f.HashId, f.Name, f.Path, f.Type, f.Description)
    if err != nil {
        fmt.Printf("Insert exec error: %v\n", err)
        return err
    }

    lastId, err := result.LastInsertId()
    if err != nil {
        return err
    }

    f.Id = int(lastId)

    return nil
}

func (f *File) ResizeImage(file io.Reader, w int, h int) error {
    // make sure file is jpeg or png
    if f.Type != "image/jpeg" && f.Type != "image/png" {
        return errors.New("Not valid image file for resizing")
    }

    var err error
    switch f.Type {
    case "image/jpeg":
        err = resizeJpeg(file, f.Path, w, h, f.Type)
        break;
    case "image/png":
        err = resizePng(file, f.Path, w, h, f.Type)
        break;
    }
    if err != nil {
        return err
    }

    return nil
}

func resizeJpeg(file io.Reader, path string, w int, h int, contentType string) error {
    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

    // resize to width 1000 using Lanczos resampling
    // and preserve aspect ratio
    m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

    out, err := os.Create(path)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)

    return nil
}

func resizePng(file io.Reader, path string, w int, h int, contentType string) error {
    // decode jpeg into image.Image
    img, err := png.Decode(file)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

    // resize to width 1000 using Lanczos resampling
    // and preserve aspect ratio
    m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

    out, err := os.Create(path)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }
    defer out.Close()

    // write new image to file
    png.Encode(out, m)

    return nil
}
