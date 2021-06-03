package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) getBitrixOrderByUserId(c *gin.Context) {
	id := c.Param("id")

	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/cno5mh8afndjmbcw/crm.deal.list.json?filter[CONTACT_ID]=" + id)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := body

	c.JSON(http.StatusOK, sb)
}

func (h *Handler) getBitrixUserByEmail(c *gin.Context) {
	email := c.Param("email")

	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/cno5mh8afndjmbcw/crm.contact.list.json?filter[EMAIL]=" + email)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := body

	c.JSON(http.StatusOK, sb)
}
