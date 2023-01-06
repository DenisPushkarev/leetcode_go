package leetcode

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	prev := prices[0]
	profit := 0
	for _, value := range prices {
		if prev < value {
			profit += value - prev
		}
		prev = value
	}
	return profit
}
