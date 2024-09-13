package handler

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createBid(c *gin.Context) {
	var bidsCreater domain.BidCreator
	if err := c.BindJSON(&bidsCreater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	_, err := h.service.Auth.CheckUserID(bidsCreater.AuthorID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	_, err = h.service.Auth.CheckUserChargeAffiliation(bidsCreater.AuthorID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

	bids, err := h.service.Bid.CreateBid(bidsCreater)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": bids})
}
