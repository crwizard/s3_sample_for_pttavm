package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var minioClient *minio.Client

func init() {
	// bu bilgiler size ilettigim s3-access-key.env dosyasinin icinde mevcut
	endpoint := "storage.googleapis.com"
	accessKeyID := "ACCESS_KEY_ID"
	secretAccessKey := "ACCESS_KEY_SECRET"
	useSSL := true

	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func uploadFile(bucketName, localFilename, remoteFilename string) {
	contentType := "application/json"

	uploadInfo, err := minioClient.FPutObject(context.Background(), bucketName, remoteFilename, localFilename, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully uploaded", remoteFilename, "of size", uploadInfo.Size, "ETag:", uploadInfo.ETag)
}

func listFiles(bucketName string) {
	for object := range minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{}) {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		fmt.Println(object)
	}
}

func main() {
	bucketName := "pttavm"
	localFilename := "abc.txt"
	remoteFilename := "test.txt"

	uploadFile(bucketName, localFilename, remoteFilename)
	listFiles(bucketName)
}
