package handler

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/globalsign/mgo/bson"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
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

func UploadPhotoToS3(s *session.Session, file multipart.File, fileName string, fileHeader *multipart.FileHeader) (string, string, int64, error) {

	size := fileHeader.Size
	originalName := fileHeader.Filename
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	keyName := fileName + filepath.Ext(originalName)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(AWS_S3_BUCKET),
		Key:                aws.String(keyName),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", "", 0, err
	}

	return keyName, originalName, size, err
}

func CompressImageResource(file multipart.File) bytes.Buffer {
	var newImage image.Image

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	newImage = resize.Resize(1000, 800, img, resize.Lanczos3)

	// encode image to buffer
	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImage, nil)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	return buf
}

func UploadResizedPhotoToS3(s *session.Session, file multipart.File, fileName string, fileHeader *multipart.FileHeader) (string, error) {
	originalName := fileHeader.Filename

	buf := CompressImageResource(file)

	fileSize := buf.Len()
	fileBytes := buf.Bytes()

	keyNameResize := fileName + "_compressed" + filepath.Ext(originalName)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(AWS_S3_BUCKET),
		Key:                aws.String(keyNameResize),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(fileBytes),
		ContentLength:      aws.Int64(int64(fileSize)),
		ContentType:        aws.String(http.DetectContentType(fileBytes)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", err
	}

	return keyNameResize, err
}

func DeleteFileFromS3(s *session.Session, fileName string) error {
	_, err := s3.New(s).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return err
	}

	return err
}

func DeleteAllItems(s *session.Session) error {
	iter := s3manager.NewDeleteListIterator(s3.New(s), &s3.ListObjectsInput{
		Bucket: aws.String(AWS_S3_BUCKET),
	})

	err := s3manager.NewBatchDeleteWithClient(s3.New(s)).Delete(aws.BackgroundContext(), iter)

	if err != nil {
		return err
	}

	return err
}
