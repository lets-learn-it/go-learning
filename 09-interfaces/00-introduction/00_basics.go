package main

import "fmt"

type Tax interface {
	calculateTax() float64
}

type Salary struct {
	amount int64
}

type Profit struct {
	stock      float64
	isLongTerm bool
}

func (salary Salary) calculateTax() float64 {
	if salary.amount > 1000000 {
		return float64(salary.amount) * 0.3
	} else if salary.amount > 500000 {
		return float64(salary.amount) * 0.2
	} else {
		return float64(0)
	}
}

func (profit Profit) calculateTax() float64 {
	if profit.isLongTerm && profit.stock > 100000 {
		return float64(profit.stock) * 0.15
	} else if profit.stock > 100000 {
		return float64(profit.stock) * 0.2
	} else {
		return float64(0)
	}
}

func totalTax(tax Tax) float64 {
	return tax.calculateTax()
}

func main() {
	// interfaces in go are similar to java interfaces
	// but in go implementing methods is not compulsary

	salary := Salary{amount: 1000000}
	profit := Profit{stock: 49000, isLongTerm: true}

	fmt.Println("Total Tax =", totalTax(salary)+totalTax(profit))
}
