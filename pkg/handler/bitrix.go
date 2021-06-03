package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) getAllBitrixOrders(c *gin.Context) {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": body,
	})
}
