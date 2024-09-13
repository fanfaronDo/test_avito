package handler

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusForbidden, gin.H{"reason": ErrForbiddenRequest})
		return
	}

	bids, err := h.service.Bid.CreateBid(bidsCreater)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *Handler) getBids(c *gin.Context) {
	userUUID, err := getUserId(c)
	if err != nil {
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	bids, err := h.service.Bid.GetBids(limit, offset, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *Handler) getBidByTenderId(c *gin.Context) {
	userUUID, err := getUserId(c)
	if err != nil {
		return
	}

	tenderid := c.Param(tenderID)
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	_, err = h.service.Auth.CheckUserCreatorBidsByTenderId(userUUID, tenderid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": ErrForbiddenRequest.Error()})
		return
	}

	bids, err := h.service.Bid.GetBidsByTenderId(limit, offset, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, bids)
}
