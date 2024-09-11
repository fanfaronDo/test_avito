package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userIDCtx = "userid"
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

		tendersWithAuth := api.Group("/tenders", h.userAuth)
		{
			tendersWithAuth.GET("/my", h.getUserTenders)
			tendersWithAuth.GET("/:tenderid/status", h.getStatusTender)
			tendersWithAuth.PUT("/:tenderid/status", h.setStatusTender)
		}
	}

	return router
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
