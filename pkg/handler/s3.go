package handler

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
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

func UploadResizedPhotoToS3(s *session.Session, file multipart.File, fileName string, fileHeader *multipart.FileHeader) (string, error) {

	size := fileHeader.Size
	originalName := fileHeader.Filename
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	resized := compressImageResource(buffer)

	keyName := fileName + "_compressed" + filepath.Ext(originalName)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(AWS_S3_BUCKET),
		Key:                aws.String(keyName),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(resized),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", err
	}

	return keyName, err
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

func compressImageResource(data []byte) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return data
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 40})
	if err != nil {
		return data
	}
	if buf.Len() > len(data) {
		return data
	}
	return buf.Bytes()
}
