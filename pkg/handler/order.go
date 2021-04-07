package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/models"
	"net/http"
)

type getAllOrdersResponse struct {
	Data []models.Order `json:"data"`
}

type getOrderResponse struct {
	Data models.Order `json:"data"`
}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.Orders.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) createOrder(c *gin.Context) {
	var input models.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Orders.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllForUserOrders(c *gin.Context) {
	id := c.Param("id")

	orders, err := h.services.Orders.GetAllForUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getAllForPhotographerOrders(c *gin.Context) {
	id := c.Param("photographer_id")

	orders, err := h.services.Orders.GetAllForPhotographer(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getAllForDesignerOrders(c *gin.Context) {
	id := c.Param("designer_id")

	orders, err := h.services.Orders.GetAllForDesigner(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getOrderById(c *gin.Context) {
	id := c.Param("id")

	order, err := h.services.Orders.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrderResponse{
		Data: order,
	})
}

func (h *Handler) updateOrder(c *gin.Context) {
	id := c.Param("id")

	var input models.UpdateOrderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Orders.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Orders.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
