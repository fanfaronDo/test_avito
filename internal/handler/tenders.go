package handler

import (
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getTenders(c *gin.Context) {
	serviceTypes := c.DefaultQuery("service_type", "")
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	fmt.Println(serviceTypes, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}

	tenders, err := h.service.Tender.GetTenders(limit, offset, serviceTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}
	fmt.Println(tenders)
	c.JSON(http.StatusOK, tenders)
}

func (h *Handler) createTender(c *gin.Context) {
	var tenderCreator domain.TenderCreator
	if err := c.BindJSON(&tenderCreator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}

}
