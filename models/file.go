package models

import (
    "github.com/speps/go-hashids"
    "fmt"
    "image"
    "github.com/nfnt/resize"
    "image/jpeg"
    "image/png"
    "os"
    "errors"
    "bufio"
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

func NewFile(name string, path string, fileType string, desc string) File {
    // default id to zero before saving to DB
    return File{
        Id: 0,
        HashId: "",
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
        fmt.Println("Delete error: %v", err)
        return errors.New("No file deleted")
    }

    return nil
}

func (f *File) generateId() {
    hd := hashids.NewData()
    hd.Salt = "$qz&vzp#rwLNe4LV6dr3>o39ei?#Rhud"
    hd.MinLength = 6
    h := hashids.NewWithData(hd)
    id, _ := h.Encode([]int{f.Id})

    f.HashId = id
}

func (f *File) Save() error {
    // If ID is 0 then call insert method to create new entry
    // and generate new ID and saving it again.
    // else call update method to update existing entry
    if f.Id == 0 {
        err := f.insert()
        if err != nil {
            return err
        }
        f.generateId()

        f.Save()
        fmt.Println("%v", f)
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
        fmt.Println("Update prepare error: %v", err)
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(f.HashId, f.Name, f.Path, f.Type, f.Id)
    if err != nil {
        fmt.Println("Update exec error: %v", err)
        return err
    }

    return nil
}

func (f *File) insert() error {
    conn := initDB()

    sql := `
	INSERT INTO files(
		name,
		path,
		type,
		description
	) values(?, ?, ?, ?)
	`

    stmt, err := conn.Prepare(sql)
    if err != nil {
        fmt.Println("Insert prepare error: %v", err)
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(f.Name, f.Path, f.Type, f.Description)
    if err != nil {
        fmt.Println("Insert exec error: %v", err)
        return err
    }

    lastId, err := result.LastInsertId()
    if err != nil {
        return err
    }

    f.Id = int(lastId)

    return nil
}

func (f *File) ResizeImage(w int, h int) error {
    // make sure file is jpeg or png
    if f.Type != "image/jpeg" && f.Type != "image/png" {
        return errors.New("Not valid image file for resizing")
    }

    i, err := os.OpenFile(f.Path, os.O_RDWR, os.ModeAppend)
    image, _, err := image.Decode(bufio.NewReader(i))
    if err != nil {
        fmt.Printf("Image decode error: %v\n", err)
        return err
    }
    defer i.Close()

    resized := resize.Resize(uint(w), uint(h), image, resize.Lanczos3)

    // overwrite existing image
    writer := bufio.NewWriter(i)
    switch f.Type {
    case "image/jpeg":
        err = jpeg.Encode(writer, resized, nil)
        break;
    case "image/png":
        err = png.Encode(writer, resized)
        break;
    }
    if err != nil {
        return err
    }

    return nil
}
