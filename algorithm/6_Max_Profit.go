package algorithm

//买卖股票的最佳时机含手续费，动态规划和贪心的结合应用


// 买卖股票的最佳时机含手续费，贪心解法（推荐）
func MaxProfitGreedy(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0
	minPrice := prices[0] // 记录最低买入价格

	for i := 1; i < len(prices); i++ {
		// 如果当前价格比之前最低价高，可以考虑卖出
		if prices[i] > minPrice {
			// 计算利润：当前价格 - 最低价格 - 手续费
			currentProfit := prices[i] - minPrice - fee

			// 如果有利润，就进行交易
			if currentProfit > 0 {
				profit += currentProfit
				// 更新最低价格（相当于重新买入）
				minPrice = prices[i]
			}
		} else {
			// 当前价格更低，更新最低价格
			minPrice = prices[i]
		}
	}

	return profit
}

// 买卖股票的最佳时机（无手续费限制），贪心解法
func MaxProfitUnlimited(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0] // 记录最低买入价格

	for i := 1; i < len(prices); i++ {
		// 如果当前价格比之前最低价高，计算卖出利润
		if prices[i] > minPrice {
			// 累加利润
			maxProfit += prices[i] - minPrice
		} else {
			// 更新最低价格
			minPrice = prices[i]
		}
	}

	return maxProfit
}

// 买卖股票的最佳时机（最多k次交易），动态规划解法
func MaxProfitK(prices []int, k int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	// DP[i][j]表示第i天完成j次交易能获得的最大利润
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 初始化第一天
	for j := 1; j <= k; j++ {
		dp[0][j] = -prices[0]
	}
	dp[0][0] = 0

	// 递推
	for i := 1; i < len(prices); i++ {
		// 第i天不交易
		dp[i][0] = 0
		for j := 1; j <= k; j++ {
			// 第i天完成j次交易：
			// 要么前一天完成j次交易今天不操作
			// 要么前一天完成j-1次交易，今天买入并完成第j次交易
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
		}
	}

	return dp[len(prices)-1][k]
}