package main

import (
	"otto/vouchers-project/config"
	"otto/vouchers-project/internal/brand"
	"otto/vouchers-project/internal/transaction"
	"otto/vouchers-project/internal/voucher"
	"otto/vouchers-project/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database := db.ConnectDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	// Brand API
	brandRepo := brand.NewRepository(database)
	brandService := brand.NewService(brandRepo)
	brandHandler := brand.NewHandler(brandService)

	// Voucher API
	voucherRepo := voucher.NewRepository(database)
	voucherService := voucher.NewService(voucherRepo)
	voucherHandler := voucher.NewHandler(voucherService)

	// Transaction API
	transactionRepo := transaction.NewRepository(database)
	transactionService := transaction.NewService(transactionRepo, voucherRepo)
	transactionHandler := transaction.NewHandler(transactionService)

	r := gin.Default()

	v1 := r.Group("v1/api")
	{
		v1.POST("brand", brandHandler.CreateBrand)

		v1.POST("voucher", voucherHandler.CreateVoucher)
		v1.GET("voucher", voucherHandler.GetVoucherByID)
		v1.GET("voucher/brand", voucherHandler.GetAllVoucherByBrand)

		v1.POST("transaction/redemption", transactionHandler.CreateTransaction)
		v1.GET("transaction/redemption", transactionHandler.GetTransactionByID)
	}

	r.Run(":3000")
}