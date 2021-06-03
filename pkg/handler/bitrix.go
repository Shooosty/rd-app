package handler

import (
	"encoding/json"
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

	strBody, err := ioutil.ReadAll(resp.Body)

	data, err := json.Marshal(string(strBody))

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
