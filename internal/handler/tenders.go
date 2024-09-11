package handler

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getTenders(c *gin.Context) {
	serviceTypes := c.DefaultQuery("service_type", "")
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

	tenders, err := h.service.Tender.GetTenders(limit, offset, serviceTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func (h *Handler) createTender(c *gin.Context) {
	var tenderCreator domain.TenderCreator
	if err := c.BindJSON(&tenderCreator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userUUID, err := h.service.Auth.CheckUserExists(tenderCreator.CreatorUsername)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	_, err = h.service.Auth.CheckUserCharge(userUUID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

	tender, err := h.service.CreateTender(tenderCreator, userUUID)
	log.Println(tender)
	c.JSON(http.StatusOK, tender)
}

func (h *Handler) getUserTenders(c *gin.Context) {
	var tenders []domain.Tender

	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUnauthorizedError.Error()})
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
	userUUID, err := h.service.Auth.CheckUserExists(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	tenders, err = h.service.Tender.GetTendersByUserUUID(limit, offset, userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tenders)
}

func (h *Handler) getStatusTender(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": service.ErrUnauthorizedError.Error()})
		return
	}
	userUUID, err := h.service.Auth.CheckUserExists(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	_, err = h.service.Auth.CheckUserCharge(userUUID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}

	tenderid := c.Param("tenderid")
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	status, err := h.service.Tender.GetStatusTenderByTenderID(tenderid, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, status)
}
