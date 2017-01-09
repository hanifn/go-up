# Go Upload Server

A File Upload server built on Go

### Dependencies
- Go (built and tested using Go 1.7.4)
- Sqlite3
- github.com/gorilla/mux
- github.com/mattn/go-sqlite3 
- github.com/nfnt/resize
- github.com/ventu-io/go-shortid
- github.com/mitchellh/goamz
- An Amazon AWS account with read/write access to an S3 bucket if you want to upload files to S3

### Installation
Run the standard `go get` command:
```
go get github.com/hanifn/go-up
```

This should install and build project to the `$GOPATH/bin` directory.
If not, then `cd` to the project root directory and run `go install`.

Or run the following from any directory:
```
go install github.com/hanifn/go-up
```

### Starting the server
Run the command:
```
go-up
```
The server will then start listening on port `8000`

### AWS S3 Support
If AWS S3 support is setup,  server will upload files to S3 bucket if
the `upsert` parameter is set in the `/upload` `POST` request. Authentication and
bucket name details must be set as environment variables.

**ENVIRONMENT VARIABLES REQUIRED**
- `AWS_ACCESS_KEY_ID` 
- `AWS_SECRET_ACCESS_KEY`
- `S3_BUCKET`

### REST API
- `POST /uploads`: Uploads a file

   Uploads a file as part of multi-part/form-data `POST` request.
   
   |parameter     |value               |
   |--------------|--------------------|
   |`file`        |file data           |
   |`description` |Optional description|
   |`resize`      |Optional. If set, server will try to resize image according to specified size. Format is `{width}x{height}`. E.g.: 120x120|
   |`upsert`      |Optional. If set to anything except an empty string, will upload the file to specified S3 bucket. See AWS S3 section for setup details.

- `GET /files`: Get a list of files
- `GET /files/:id`: Download specific file
- `DELETE /files/:id`: Delete specific file

### TODO
- Set server listening port as a flag argument
- An authentication system would be nice. Then we can limit uploads and
file listings to only authed user
- File compression?