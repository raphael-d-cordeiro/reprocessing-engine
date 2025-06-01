package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Order represents a simplified mock order
type Order struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	ProductKey string `json:"product_key"`
}

// OrdersResponse wraps the list of orders and pagination info
type OrdersResponse struct {
	Data       []Order `json:"data"`
	TotalPages int     `json:"total_pages"`
}

func GetOrders(c *gin.Context) {
	ctx := c.Request.Context()
	select {
	case <-ctx.Done():
		log.Println("[handler] client cancelled the request")
		return
	default:
		// continue processing
	}
	pageParam := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	// Mock orders per page
	orders := []Order{
		{ID: "order-" + strconv.Itoa(page*10+1), Status: "waiting-for-integration", ProductKey: "PRD-001"},
		{ID: "order-" + strconv.Itoa(page*10+2), Status: "waiting-for-integration", ProductKey: "PRD-002"},
	}

	c.JSON(http.StatusOK, OrdersResponse{
		Data:       orders,
		TotalPages: 5, // sempre 5 páginas mockadas
	})
}
