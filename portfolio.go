package main

import(
	"fmt"
	"strconv"
)




type portfolio struct {
	balances map[string]float64
}

func newPortfolio() *portfolio{
	var port portfolio

	port.balances = make(map[string]float64)
	port.balances["dollar"] = 500000.00
	port.balances["ETH"] = 100.00
	return &port


}

func (port portfolio) printBalance(currency string) {
	fmt.Println(currency + ":" + strconv.FormatFloat(port.balances[currency], 'f', 2, 64))
}

