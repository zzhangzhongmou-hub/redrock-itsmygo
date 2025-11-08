package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) totalvalue() float64 {
	return p.Price * float64(p.Stock)
}

func (p Product) islnStock() bool {
	return p.Stock > 0
}
func (p Product) info() string {
	return fmt.Sprintf("商品：%s,单价:%.1f,库存：%d", p.Name, p.Price, p.Stock)
}
func (p Product) Restock(amount int) {
	p.Stock += amount
	fmt.Printf("进货%d本,当前库存%d本\n", amount, p.Stock)
}
func (p *Product) Sell(amount int) (bool, string) {
	if p.Stock < amount {
		return false, "库存不足"
	}
	p.Stock -= amount
	fmt.Printf("售卖成功，剩余库存%d\n", p.Stock)
	return true, ""
}
func main() {

	product := Product{
		Name:  "go编程书",
		Price: 89.5,
		Stock: 10,
	}
	product.Sell(5)
	product.Restock(20)
	product.Sell(30)

	fmt.Printf("商品信息：%s\n", product.info())
	fmt.Println("商品总价值:", product.totalvalue())
}
