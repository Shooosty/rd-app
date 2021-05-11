package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/pkg/repository"
	"net/http"
)

type MailOptions struct {
	Name  string `json:"name"`
	Email string `json:"string"`
}

func (h *Handler) sendNewOrderMessage(c *gin.Context) {
	var input MailOptions

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	subject := "Новый заказ"
	text := "У вас новый заказ на платформе lk.rhinodesign.ru"
	html := "<b>" + input.Name + "," + "</b>" + "<p>" + "У вас новый заказ на платформе lk.rhinodesign.ru" + "<p>" + "</br>"

	_ = repository.SendMail(subject, text, html, input.Name, input.Email)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
