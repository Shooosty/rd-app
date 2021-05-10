package handler

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const (
	AWS_S3_REGION = "eu-west-3"
	AWS_S3_BUCKET = "rhinodesign"
)

type File struct {
	Url string `json:"url"`
}

func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	size := fileHeader.Size
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	tempFileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(AWS_S3_BUCKET),
		Key:           aws.String(tempFileName),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

func uploadFile(c *gin.Context) {
	maxSize := int64(40000000) // 5mb max

	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Image too large")
	} else {
		logrus.Print("Image size is ok")
	}

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not get uploaded file")
		return
	}
	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			"AKIAZ4EXIBF2T6T7UB64",
			"qqBiCHLMG7Nn9rGaIueZwnNxyBwiOGMw0AdK0UUn",
			""),
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	fileName, err := UploadFileToS3(s, file, header)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	data := "https://renti-api-s3.s3.eu-west-3.amazonaws.com/" + fileName

	c.JSON(http.StatusOK, File{
		Url: data,
	})
}
