package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/models"
	"net/http"
)

type getAllPersonsResponse struct {
	Data []models.Person `json:"data"`
}

func (h *Handler) getAllPersons(c *gin.Context) {
	persons, err := h.services.Persons.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPersonsResponse{
		Data: persons,
	})
}

func (h *Handler) createPerson(c *gin.Context) {
	var input models.Person
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Persons.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updatePerson(c *gin.Context) {
	id := c.Param("id")

	var input models.UpdatePersonInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Persons.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deletePerson(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Persons.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
