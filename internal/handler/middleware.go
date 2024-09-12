package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUnauthorizedError.Error()})
		return
	}
	userUUID, err := h.service.Auth.GetUserId(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userIDCtx, userUUID)
}

func (h *Handler) userChargeIdentity(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUnauthorizedError.Error()})
		return
	}

	userUUID, err := h.service.Auth.GetUserId(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	userUUID, err = h.service.Auth.GetUserCharge(userUUID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userIDCtx, userUUID)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userIDCtx)
	if !ok {
		return "", service.ErrUserNotFound
	}
	idS, ok := id.(string)
	if !ok {
		return "", ErrUserIdInvalidType
	}
	return idS, nil
}
