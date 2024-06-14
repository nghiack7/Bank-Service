package main

import (
	"nghiack7/bank-it/internal/bank/deliveries/http"
	"nghiack7/bank-it/internal/bank/services"

	"github.com/gin-gonic/gin"
)

func main() {
	loanService := services.NewLoanService()
	loanHandler := http.NewLoanHandler(loanService)
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	http.NewLoanRoutes(router.Group("/api/v1"), loanHandler)

	router.Run(":8080")
}
