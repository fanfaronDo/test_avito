package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userIDCtx = "userid"
	tenderID  = "tenderId"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/ping", ping)

		tenders := api.Group("/tenders")
		{
			tenders.GET("/", h.getTenders)
			tenders.POST("/new", h.createTender)
		}

		tendersWithAuth := api.Group("/tenders", h.userChargeIdentity)
		{
			tendersWithAuth.GET("/my", h.getUserTenders)
			tendersWithAuth.GET("/:tenderId/status", h.getStatusTender)
			tendersWithAuth.PUT("/:tenderId/status", h.setStatusTender)
			tenders.PATCH("/:tenderId/edit", h.editTender)
			tenders.PUT("/:tenderId/rollback/:version", h.rollbackTender)
		}
	}

	return router
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
