package handler

import (
	"net/http"

	"dataforge/ingest-go/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	svc *service.HealthService
}

func NewHealthHandler(svc *service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", h.health)
}

func (h *HealthHandler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": h.svc.Status()})
}
