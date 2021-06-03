package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAllBitrixOrders(c *gin.Context) {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	data := json.NewDecoder(resp.Body).Decode(&result)

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
