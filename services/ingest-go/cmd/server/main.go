package main

import (
	"log"

	"dataforge/ingest-go/internal/config"
	"dataforge/ingest-go/internal/handler"
	"dataforge/ingest-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	// Production-safe defaults
	gin.SetMode(gin.ReleaseMode)

	healthSvc := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthSvc)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Don’t trust all proxies by default (security best practice)
	_ = r.SetTrustedProxies(nil)

	healthHandler.RegisterRoutes(r)

	addr := ":" + cfg.Port
	log.Printf("ingest-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
