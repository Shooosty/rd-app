package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/models"
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
	var input models.Photo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
