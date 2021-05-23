package handler

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const (
	AWS_S3_REGION = "eu-west-3"
	AWS_S3_BUCKET = "rhinodesign"
)

func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, string, int64, error) {

	size := fileHeader.Size
	originalName := fileHeader.Filename
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	tempFileName := bson.NewObjectId().Hex() + filepath.Ext(originalName)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(AWS_S3_BUCKET),
		Key:                aws.String(tempFileName),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", "", 0, err
	}

	return tempFileName, originalName, size, err
}

//func downloadFile(name Name) *os.File {
//
//	s, _ := session.NewSession(&aws.Config{
//		Region: aws.String(AWS_S3_REGION),
//		Credentials: credentials.NewStaticCredentials(
//			"AKIAZ4EXIBF2T6T7UB64",
//			"qqBiCHLMG7Nn9rGaIueZwnNxyBwiOGMw0AdK0UUn",
//			""),
//	})
//
//	file, _ := os.Create(name.FileName)
//
//	downloader := s3manager.NewDownloader(s)
//
//	_, _ = downloader.Download(file, &s3.GetObjectInput{
//		Bucket: aws.String(AWS_S3_BUCKET),
//		Key:    aws.String(name.FileName),
//	})
//
//	return file
//}
