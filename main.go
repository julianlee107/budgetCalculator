package main

import (
	"budgetBook/loan"
	"budgetBook/user"
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	l := loan.Loan{}
	l.SetLoan(1124000, 5.68, 360)
	m := loan.Mortgage{}
	m.SetMortgage(l, "MRP")
	budget := user.Budget{}
	budgetTotal := *budget.SetBudgetTotal("700000")
	if m.GetMortgageMode() == "MRPI" {
		i := 0
		monthlyConfession, _ := m.Loan.MRPI()
		for budgetTotal.GetBudgetTotal().Cmp(decimal.NewFromInt(0)) != -1 {
			i++
			budgetTotal.UpdateBudgetTotal(budgetTotal.GetBudgetTotal().Sub(monthlyConfession))
		}
		fmt.Printf("第%d个月消耗完毕", i)
	} else if m.GetMortgageMode() == "MRP" {
		interestList := m.Loan.MPR(0, m.GetPrincipleMonthly())
		for index, interest := range interestList {
			budgetTotal.UpdateBudgetTotal(budgetTotal.GetBudgetTotal().Sub(m.GetPrincipleMonthly().Add(interest)))
			if budgetTotal.GetBudgetTotal().Cmp(decimal.NewFromInt(0)) == -1 {
				fmt.Printf("第%d个月消耗完毕", index)
				break
			}
		}
	}
}
