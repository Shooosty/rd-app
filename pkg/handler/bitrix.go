package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type order struct {
	ID    int
	TITLE string
}

func (h *Handler) getAllBitrixOrders(c *gin.Context) {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	data := json.NewDecoder(resp.Body).Decode(&order{})

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
