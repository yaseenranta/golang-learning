package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tips  float64
	total float64
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price

	b.total = b.total + price
}

func (b *bill) updateTip(tip float64) {
	b.tips = tip
}

func (b *bill) format() string {

	fs := fmt.Sprintf("Biller Name %v \n", b.name)

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v $%0.2f\n", k+":", v)
	}

	fs += fmt.Sprintf("%-25v $%0.2f \n", "Tip:", b.tips)
	fs += fmt.Sprintf("%-25v $%0.2f \n", "Total:", b.total+b.tips)

	return fs
}

func main() {

	var b bill

	b.name = "Muhammad Yaseen"
	b.items = map[string]float64{}

	b.addItem("Cake", 15.1)
	b.addItem("Biscuts", 16.5)

	b.updateTip(1)

	fmt.Println(b.format())

}
