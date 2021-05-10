package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/models"
	"net/http"
)

func (h *Handler) createPhoto(c *gin.Context) {
	uploadFile(c)
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
