package services

import (
    "github.com/mitchellh/goamz/aws"
    "github.com/mitchellh/goamz/s3"
    "os"
    "fmt"
)

var (
    bucketName string
)

func setup() (aws.Auth, error) {
    // The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
    auth, err := aws.EnvAuth()
    if err != nil {
        return aws.Auth{}, err
    }

    // get bucket from S3_BUCKET env var
    bucketName = os.Getenv("S3_BUCKET")

    return auth, nil
}

func UpsertToS3(name string, file []byte, fileType string) error {
    auth, err := setup()
    if err != nil {
        return err
    }

    // Open Bucket
    s := s3.New(auth, aws.APSoutheast)
    bucket := s.Bucket(bucketName)
    err = bucket.Put(name, file, fileType, s3.BucketOwnerFull)
    if err != nil {
        return err
    }

    return nil
}

func DownloadFromS3(name string) ([]byte, error) {
    auth, err := setup()
    if err != nil {
        return nil, err
    }

    // Open Bucket
    s := s3.New(auth, aws.APSoutheast)
    bucket := s.Bucket(bucketName)
    file, err := bucket.Get(name)
    if err != nil {
        fmt.Printf("%v\n", err)
        return nil, err
    }

    return file, nil
}
func RemoveFromS3(name string) error {
    auth, err := setup()
    if err != nil {
        return err
    }

    // Open Bucket
    s := s3.New(auth, aws.APSoutheast)
    bucket := s.Bucket(bucketName)
    err = bucket.Del(name)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

    return nil
}