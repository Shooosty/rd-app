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

type getPhotoResponse struct {
	Data models.Photo `json:"data"`
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

func (h *Handler) getPhotoById(c *gin.Context) {
	id := c.Param("id")

	photo, err := h.services.Photos.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getPhotoResponse{
		Data: photo,
	})
}

func (h *Handler) getAllPhotosByPersonId(c *gin.Context) {
	id := c.Param("id")

	photos, err := h.services.Photos.GetAllPhotosByPersonId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPhotosResponse{
		Data: photos,
	})
}

func (h *Handler) getAllPhotosByOrderId(c *gin.Context) {
	id := c.Param("id")

	photos, err := h.services.Photos.GetAllPhotosByOrderId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPhotosResponse{
		Data: photos,
	})
}

func (h *Handler) createPhoto(c *gin.Context) {
	personId := c.Param("personId")
	orderId := c.Param("orderId")
	fileName := c.Param("fileName")

	var input models.Photo

	maxSize := int64(25000000) // 25mb max

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

	keyName, originalName, size, err := UploadPhotoToS3(s, file, fileName, header)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	keyNameResize, err := UploadResizedPhotoToS3(s, file, fileName, header)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Could not upload file")
		return
	}

	url := "https://rhinodesign.s3.eu-west-3.amazonaws.com/" + keyName
	urlResize := "https://rhinodesign.s3.eu-west-3.amazonaws.com/" + keyNameResize

	input.Name = originalName
	input.NameS3 = keyName
	input.Url = url
	input.UrlResize = urlResize
	input.Size = size / 1024
	input.PersonId = personId
	input.OrderId = orderId
	input.Type = "image/*"

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

func (h *Handler) deletePhotoOnS3(c *gin.Context) {
	fileName := c.Param("fileName")

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			"AKIAZ4EXIBF2T6T7UB64",
			"qqBiCHLMG7Nn9rGaIueZwnNxyBwiOGMw0AdK0UUn",
			""),
	})

	err = DeleteFileFromS3(s, fileName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAllItems(c *gin.Context) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			"AKIAZ4EXIBF2T6T7UB64",
			"qqBiCHLMG7Nn9rGaIueZwnNxyBwiOGMw0AdK0UUn",
			""),
	})

	err = DeleteAllItems(s)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
