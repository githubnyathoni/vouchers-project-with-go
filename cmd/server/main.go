package main

import (
	"otto/vouchers-project/config"
	"otto/vouchers-project/internal/brand"
	"otto/vouchers-project/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database := db.ConnectDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	brandRepo := brand.NewRepository(database)
	brandService := brand.NewService(brandRepo)
	brandHandler := brand.NewHandler(brandService)

	r := gin.Default()

	v1 := r.Group("v1/api")
	{
		v1.POST("brand", brandHandler.CreateBrand)
	}

	r.Run(":3000")
}