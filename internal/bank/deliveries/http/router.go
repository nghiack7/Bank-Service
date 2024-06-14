package http

import "github.com/gin-gonic/gin"

func NewLoanRoutes(r *gin.RouterGroup, loan LoanHandler) {
	r.POST("/loan", loan.CreateLoan)
}
