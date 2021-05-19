package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

type getAllPhotosResponse struct {
	Data []models.Photo `json:"data"`
}

func (h *Handler) getAllPhotos(c *gin.Context) {
	photos, err := h.services.Photos.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPhotosResponse{
		Data: photos,
	})
}

func (h *Handler) getAllByPersonId(c *gin.Context) {
	id := c.Param("id")

	photos, err := h.services.Photos.GetAllByPersonId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPhotosResponse{
		Data: photos,
	})
}

func (h *Handler) createPhoto(c *gin.Context) {
	personId := c.Param("id")

	var input models.Photo

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

	fileName, originalName, size, err := UploadFileToS3(s, file, header)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	url := "https://rhinodesign.s3.eu-west-3.amazonaws.com/" + fileName

	input.Name = originalName
	input.Url = url
	input.Size = size / 1024
	input.PersonID = personId

	id, err := h.services.Photos.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deletePhoto(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Photos.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
