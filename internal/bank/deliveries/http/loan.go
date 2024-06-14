package http

import (
	"nghiack7/bank-it/internal/bank/dtos"
	"nghiack7/bank-it/internal/bank/services"

	"github.com/gin-gonic/gin"
)

type LoanHandler interface {
	CreateLoan(c *gin.Context)
}

type loanHandler struct {
	loanSvc services.LoanService
}

func NewLoanHandler(loanSvc services.LoanService) LoanHandler {
	return &loanHandler{loanSvc: loanSvc}
}

func (h *loanHandler) CreateLoan(c *gin.Context) {
	var req dtos.LoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := req.Validate()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, httpCode, err := h.loanSvc.CreateLoan(c.Request.Context(), &req)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(httpCode, resp)
}
