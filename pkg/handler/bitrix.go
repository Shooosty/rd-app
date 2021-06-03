package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) getAllBitrixOrders(c *gin.Context) {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/cno5mh8afndjmbcw/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	c.JSON(http.StatusOK, gin.H{
		"data": sb,
	})
}

func (h *Handler) getBitrixUserByEmail(c *gin.Context) {
	email := c.Param("email")

	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/cno5mh8afndjmbcw?filter[EMAIL]=" + email)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	c.JSON(http.StatusOK, gin.H{
		"data": sb,
	})
}
