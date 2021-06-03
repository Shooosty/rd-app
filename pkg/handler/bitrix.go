package handler

import (
	"io"
	"log"
	"net/http"
)

func (h *Handler) getAllBitrixOrders() io.ReadCloser {
	resp, err := http.Get("https://rosfotoproekt.bitrix24.ru/rest/3872/l00jxlvjy0aamuom/crm.deal.list.json")
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.Body

	return body
}
