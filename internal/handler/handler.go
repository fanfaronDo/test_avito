package handler

import (
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userID   = "userid"
	tenderID = "tenderId"
	bidID    = "bidId"
	version  = "version"
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

		userTenders := api.Group("/tenders", h.userIdentity)
		{
			userTenders.GET("/my", h.getUserTenders)
		}

		editor := api.Group("/tenders", h.userAuthorisation)
		{
			editor.GET("/:tenderId/status", h.getStatusTender)
			editor.PUT("/:tenderId/status", h.setStatusTender)
			editor.PATCH("/:tenderId/edit", h.editTender)
			editor.PUT("/:tenderId/rollback/:version", h.rollbackTender)
		}

		bids := api.Group("/bids")
		{
			bids.POST("/new", h.createBid)
		}

		userBids := api.Group("/bids", h.userIdentityBids)
		{
			userBids.GET("/my", h.getBids)
			userBids.GET("/:tenderId/list", h.getBidByTenderId)
		}

		//editorBids := api.Group("/bids", h.userAuthorisationBids)
		//{
		//	editorBids.GET("/:tenderId/list", h.getBidByTenderId)
		//}
	}

	return router
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
