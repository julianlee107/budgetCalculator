package user

import "github.com/shopspring/decimal"

type Budget struct {
	total decimal.Decimal
}

type user struct {
	username  string
	pwd       string
	lastLogin string
	budget    decimal.Decimal
}

func (b *Budget) SetBudgetTotal(total string) *Budget {
	to, err := decimal.NewFromString(total)
	if err != nil {
		panic(err)
	}
	b.total = to
	return b
}

func (b Budget) GetBudgetTotal() decimal.Decimal {
	return b.total
}

func (b *Budget) UpdateBudgetTotal(d decimal.Decimal) {
	b.total = d
}
