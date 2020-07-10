package main

func finalPrices(prices []int) []int {
	stack := make([]int, 0, len(prices))
	for i, price := range prices {
		for len(stack) != 0 && price <= prices[stack[len(stack)-1]] {
			prices[stack[len(stack)-1]] -= price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}
