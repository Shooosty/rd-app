package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAllBitrixOrders(c *gin.Context) {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	data := resp.Body

	c.JSON(200, gin.H{
		"data":  data,
		"data2": resp,
	})
}
