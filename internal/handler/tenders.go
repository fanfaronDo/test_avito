package handler

import (
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
	userUUID, err := h.service.Auth.GetUserId(tenderCreator.CreatorUsername)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	_, err = h.service.Auth.CheckUserCharge(userUUID, tenderCreator.OrganizationID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"reason": err.Error()})
		return
	}
	var userChargeId string

	if !h.service.Auth.IsUserChargeExist(tenderCreator.CreatorUsername) {
		userChargeId, err = h.service.Auth.CreateUserCharge(userUUID, tenderCreator.CreatorUsername)
	} else {
		userChargeId, err = h.service.Auth.GetUserChargeId(tenderCreator.CreatorUsername)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	tender, err := h.service.CreateTender(tenderCreator, userChargeId)

	c.JSON(http.StatusOK, tender)
}

func (h *Handler) getUserTenders(c *gin.Context) {
	var tenders []domain.Tender

	userUUID, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
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

	tenders, err = h.service.Tender.GetTendersByUserUUID(limit, offset, userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenders)
}

func (h *Handler) getStatusTender(c *gin.Context) {
	userUUID, err := getUserId(c)
	if err != nil {
		return
	}

	tenderid := c.Param("tenderid")
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	status, err := h.service.Tender.GetStatusTenderByTenderID(tenderid, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}
	c.String(http.StatusOK, status)
}

func (h *Handler) setStatusTender(c *gin.Context) {
	userUUID, err := getUserId(c)
	if err != nil {
		return
	}
	tenderid := c.Param("tenderid")
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}
	status, isExist := c.GetQuery("status")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}
	tender, err := h.service.Tender.UpdateStatusTender(tenderid, status, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func (h *Handler) editTender(c *gin.Context) {
	var tenderEditor domain.TenderEditor
	userUUID, err := getUserId(c)
	if err != nil {
		return
	}
	tenderid := c.Param("tenderid")
	if tenderid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
		return
	}

	if err := c.BindJSON(&tenderEditor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	tender, err := h.service.Tender.EditTender(tenderid, userUUID, &tenderEditor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tender)
}

//func (h *Handler) rollbackTender(c *gin.Context) {
//	tenderid := c.Param("tenderid")
//	if tenderid == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
//		return
//	}
//	version, isExist := c.GetQuery("version")
//	if !isExist {
//		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
//		return
//	}
//	versionINT, err := strconv.Atoi(version)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"reason": ErrUnsupportedRequest.Error()})
//		return
//	}
//
//}
