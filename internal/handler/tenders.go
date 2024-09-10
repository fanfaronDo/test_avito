package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) getTenders(c *gin.Context) {
	serviceTypes := c.DefaultQuery("service_type", "")
	page, _ := strconv.Atoi(c.DefaultQuery("offset", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	fmt.Println(serviceTypes, page, limit)
}
