package main

import (
	"fmt"
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
		language := determineLanguage(ctx)
		fmt.Printf("Language preference: %s\n", language)
		template := "index_en.html"
		// Render the appropriate template based on language
		if language == "vi" {
			template = "index_vi.html"
		}
		ctx.HTML(200, template, gin.H{})
	})

	http.NewLoanRoutes(router.Group("/api/v1"), loanHandler)

	router.Run(":8080")
}

// determineLanguage determines the language preference from request (for example, from headers or query parameters)
func determineLanguage(ctx *gin.Context) string {
	// For simplicity, let's assume a query parameter "lang" is used to determine language
	lang := ctx.Query("lang")
	if lang != "" {
		return lang
	}
	// Default to Vietnamese
	return "vi"
}
