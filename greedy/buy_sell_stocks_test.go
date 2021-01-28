package greedy

import (
	"fmt"
	"testing"
)

//leetcode: 122

func TestBuy_Sell_stocks(t *testing.T) {
	var price = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(price))

}

func maxProfit(price []int) int {
	profit := 0
	for k := 0; k < len(price)-1; k++ {
		if price[k] < price[k+1] {
			profit += price[k+1] - price[k]
		}
	}
	return profit

}
