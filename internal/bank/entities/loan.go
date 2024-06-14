package entities

type Loan struct {
	UserID                string
	Amount                float64
	AnnualInterestHard    float64
	TermYearHard          int
	AnnualInterestDynamic float64
	TermYearTotals        int
	PaymentPermonth       map[int]PaymentMonthInfo
}

type PaymentMonthInfo struct {
	PaymentMonth float64
	Interest     float64
	Principal    float64
	Debt         float64
}

func (l *Loan) Calculate() {
	amountPayPerMonth := l.Amount / float64(l.TermYearTotals*12)
	monthHard := l.TermYearHard * 12
	for i := 0; i < l.TermYearTotals*12; i++ {
		var interest float64
		if i <= monthHard {
			interest = l.Amount * float64(l.AnnualInterestHard/12/100)

		} else {
			interest = l.Amount * float64(l.AnnualInterestDynamic/12/100)
		}
		l.Amount -= (amountPayPerMonth)

		paymentInfo := PaymentMonthInfo{
			PaymentMonth: amountPayPerMonth + interest,
			Interest:     interest,
			Principal:    amountPayPerMonth,
			Debt:         l.Amount,
		}
		l.PaymentPermonth[i] = paymentInfo
	}
}
