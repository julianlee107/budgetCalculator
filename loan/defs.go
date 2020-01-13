package loan

import (
	"github.com/shopspring/decimal"
)

type Loan struct {
	principle decimal.Decimal
	rate      decimal.Decimal
	duration  int64
}
type Mortgage struct {
	Loan
	mode string
}

func (l Loan) GetLoan() Loan {
	return l
}
func (l *Loan) SetLoan(p float64, r float64, duration int64) *Loan {
	l.duration = duration
	l.rate = decimal.NewFromFloat(r / (12 * 100))
	l.principle = decimal.NewFromFloat(p)
	return l
}

func (m *Mortgage) SetMortgage(l Loan, mode string) *Mortgage {
	m.Loan = l
	m.mode = mode
	return m
}

func (m Mortgage) GetMortgageMode() string {
	return m.mode
}

func (m Mortgage) GetPrincipleMonthly() decimal.Decimal {
	return m.Loan.principle.DivRound(decimal.NewFromInt(m.Loan.duration), 2)
}
