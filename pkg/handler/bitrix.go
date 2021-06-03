package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Orders struct {
	ID int
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (h *Handler) getAllBitrixOrders(c *gin.Context) {

	result := Orders{}
	_ = getJson("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json", &result)

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
