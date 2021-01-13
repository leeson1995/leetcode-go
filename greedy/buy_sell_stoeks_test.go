package greedy

import (
	"fmt"
	"testing"
)

//122

func TestBuy_Sell_stoeks(t *testing.T) {
	var price = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(price))

}

func maxProfit(price []int) int {
	var profit = 0
	len := len(price) - 1
	for k := 0; k < len; k++ {
		if price[k] < price[k+1] {
			profit += price[k+1] - price[k]
		}
	}
	return profit

}
