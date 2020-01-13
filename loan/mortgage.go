package loan

import (
	"github.com/shopspring/decimal"
)

//matching the payment of principle and interest
//等额本息
func (l Loan) MRPI() (decimal.Decimal, decimal.Decimal) {
	principle := l.principle
	rate := l.rate
	duration := decimal.NewFromInt(int64(l.duration))
	numerator := principle.Mul(rate).Mul(rate.Add(decimal.NewFromInt(1)).Pow(duration))
	denominator := rate.Add(decimal.NewFromInt(1)).Pow(duration).Sub(decimal.NewFromInt(1))
	monthlyConfession := numerator.DivRound(denominator, 2)
	return monthlyConfession, principle
}

// matching the principle repayment
//等额本金
func (l *Loan) MPR(alreadyPay int64, principleMonthly decimal.Decimal) []decimal.Decimal {
	var interestLis []decimal.Decimal
	principle := l.principle
	interest := principle.Mul(l.rate).DivRound(decimal.NewFromInt(1), 2)
	interestLis = append(interestLis, interest)
	for l.duration > 1 {
		l.duration = l.duration - 1
		l.principle = l.principle.Sub(principleMonthly)
		return append(interestLis, l.MPR(alreadyPay+1, principleMonthly)...)
	}
	return interestLis
}
