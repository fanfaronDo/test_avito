package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		return
	}

	userUUID, err := h.service.Auth.GetUserId(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	defer c.Set(userIDCtx, userUUID)

	tenderid := c.DefaultQuery(tenderID, "")
	if tenderid == "" {
		return
	}

	userUUID, err = h.service.Auth.CheckUserCreatorTender(userUUID, tenderid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

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
