package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Contract struct {
	Url string `json:"url"`
}

func (h *Handler) createContract(c *gin.Context) {

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

	fileName, _, _, err := UploadFileToS3(s, file, header)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	url := "https://rhinodesign.s3.eu-west-3.amazonaws.com/" + fileName

	c.JSON(http.StatusOK, Contract{
		Url: url,
	})
}
