package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
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

	c.Set(userID, userUUID)
}

func (h *Handler) userAuthorisation(c *gin.Context) {
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

	tenderid := c.Param(tenderID)
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	userUUID, err = h.service.Auth.CheckUserCreatorTender(userUUID, tenderid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userID, userUUID)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userID)
	if !ok {
		return "", service.ErrUserNotFound
	}
	idS, ok := id.(string)
	if !ok {
		return "", ErrUserIdInvalidType
	}
	return idS, nil
}

func getTenderId(c *gin.Context) (string, error) {
	id, ok := c.Get(tenderID)
	if !ok {
		return "", service.ErrTenderNotFound
	}
	idS, ok := id.(string)
	if !ok {
		return "", ErrTenderIdInvalidType
	}
	return idS, nil
}
