package dtos

import (
	"nghiack7/bank-it/internal/bank/entities"

	"github.com/go-playground/validator/v10"
)

type LoanRequest struct {
	UserID                string  `json:"user_id"`
	Amount                float64 `json:"amount" validate:"required"`
	AnnualInterestHard    float64 `json:"annual_interest_hard" validate:"required"`
	TermYearHard          int     `json:"term_year_hard" validate:"required"`
	AnnualInterestDynamic float64 `json:"annual_interest_dynamic" validate:"required"`
	TermYearTotals        int     `json:"term_year_totals" validate:"required"`
}

func (r *LoanRequest) Validate() error {
	validation := validator.New()
	return validation.Struct(r)
}

type LoanResponse struct {
	PaymentPermonth []Payment `json:"payment_permonth"`
}

type Payment struct {
	Month     int     `json:"month"`
	Total     float64 `json:"total"`
	Interest  float64 `json:"interest"`
	Principal float64 `json:"principal"`
	Debt      float64 `json:"debt"`
}

func ToEntity[T any, U any](data T, funcTrans func(T) U) U {
	return funcTrans(data)
}

func NewLoanEntity(data *LoanRequest) *entities.Loan {
	return &entities.Loan{
		UserID:                data.UserID,
		Amount:                data.Amount,
		AnnualInterestHard:    data.AnnualInterestHard,
		TermYearHard:          data.TermYearHard,
		AnnualInterestDynamic: data.AnnualInterestDynamic,
		TermYearTotals:        data.TermYearTotals,
		PaymentPermonth:       make(map[int]entities.PaymentMonthInfo),
	}
}

func CreateResponse(data *entities.Loan) LoanResponse {
	resp := LoanResponse{}
	for k, v := range data.PaymentPermonth {
		info := Payment{
			Month:     k,
			Total:     v.PaymentMonth,
			Interest:  v.Interest,
			Principal: v.Principal,
			Debt:      v.Debt,
		}
		resp.PaymentPermonth = append(resp.PaymentPermonth, info)
	}
	QuickSort(resp.PaymentPermonth, 0, len(resp.PaymentPermonth)-1)
	return resp
}

func QuickSort(payments []Payment, low, high int) {
	if low < high {
		p := partition(payments, low, high)
		QuickSort(payments, low, p-1)
		QuickSort(payments, p+1, high)
	}
}

func partition(payments []Payment, low, high int) int {
	pivot := payments[high]
	i := low - 1
	for j := low; j < high; j++ {
		if payments[j].Month < pivot.Month {
			i++
			payments[i], payments[j] = payments[j], payments[i]
		}
	}
	payments[i+1], payments[high] = payments[high], payments[i+1]
	return i + 1
}
