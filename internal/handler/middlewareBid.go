package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentityBids(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	userUUID, err := h.service.Auth.GetUserId(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUserNotFound.Error()})
		return
	}

	_, err = h.service.Auth.CheckUserChargeAffiliation(userUUID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUserNotFound.Error()})
		return
	}

	c.Set(userID, userUUID)
}

func (h *Handler) userAuthorisationBids(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	userUUID, err := h.service.Auth.GetUserId(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUserNotFound.Error()})
		return
	}
	_, err = h.service.Auth.CheckUserChargeAffiliation(userUUID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUserNotFound.Error()})
		return
	}

	bidid := c.Param(bidID)
	if bidid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	userUUID, err = h.service.Auth.CheckUserCreatorBids(userUUID, bidid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": ErrForbiddenRequest.Error()})
		return
	}

	c.Set(userID, userUUID)
}
