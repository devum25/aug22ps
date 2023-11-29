package greedy

// 122. Best Time to Buy and Sell Stock II
// Companies
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
// Find and return the maximum profit you can achieve.

// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.
// Example 2:
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Total profit is 4.
// Example 3:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
func BuyAndSellStock(prices []int) int {
	profit := 0
	buy := 0
	sell := 0

	buy = prices[0]
	isPurchased := true
	for i := 1; i < len(prices); i++ {

		if !isPurchased {
			isPurchased = true
			buy = prices[i]
			continue
		}

		if prices[i] > buy {
			if (i + 1) < len(prices) {
				if prices[i+1] < prices[i] {
					sell = prices[i]
					profit += sell - buy
					buy = 0
					sell = 0
					isPurchased = false
				}
			} else {
				sell = prices[i]
				profit += sell - buy
				buy = 0
				sell = 0
				isPurchased = false
			}

		} else if prices[i] < buy {
			buy = prices[i]
			isPurchased = true
		}
	}

	if profit == 0 && prices[len(prices)-1] > prices[0] {
		profit = prices[len(prices)-1] - prices[0]
	}

	return profit

}

////////////////////////////////////////////////////////////////////////////

// Clean Code

func MaxProfit(prices []int) int {

	profitFromPriceGain := 0

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {
			profitFromPriceGain += (prices[i+1] - prices[i])
		}
	}

	return profitFromPriceGain
}
