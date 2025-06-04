package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Order represents a simplified mock order
// @Description Order model
// @Description Includes ID, status, and product key
// @Description Used in API responses
// @name Order
// @tags orders
// @produce json
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

// GetOrders godoc
// @Summary Get mock orders
// @Description Returns paginated mock orders by partnerKey
// @Tags orders
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param partnerKey query string true "Partner Key"
// @Success 200 {object} OrdersResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders [get]
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
	partnerKey := c.Query("partnerKey")
	if partnerKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "partnerKey is required"})
		return
	}

	// Simulate orders for different partnerKeys
	var orders []Order
	switch partnerKey {
	case "foobank":
		orders = []Order{
			{ID: "foo-" + strconv.Itoa(page*10+1), Status: "waiting-for-integration", ProductKey: "FOO-123"},
			{ID: "foo-" + strconv.Itoa(page*10+2), Status: "waiting-for-integration", ProductKey: "FOO-456"},
		}
	case "barbank":
		orders = []Order{
			{ID: "bar-" + strconv.Itoa(page*10+1), Status: "waiting-for-integration", ProductKey: "BAR-789"},
			{ID: "bar-" + strconv.Itoa(page*10+2), Status: "waiting-for-integration", ProductKey: "BAR-101"},
		}
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "partnerKey not recognized"})
		return
	}

	c.JSON(http.StatusOK, OrdersResponse{
		Data:       orders,
		TotalPages: 5, // sempre 5 páginas mockadas
	})
}
