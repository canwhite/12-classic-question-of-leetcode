package algorithm

//爬楼梯问题，斐波那契数列的应用
func ClimbStairs(n int) int {
	// 处理边界情况
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 使用动态规划，dp[i] = dp[i-1] + dp[i-2]
	// dp[i] 表示爬到第i阶台阶的方法数
	prev1, prev2 := 1, 2 // 分别对应第1阶和第2阶的方法数

	// 从第3阶开始计算
	for i := 3; i <= n; i++ {
		// 当前台阶的方法数 = 前两阶方法数之和
		current := prev1 + prev2
		prev1, prev2 = prev2, current
	}

	// 最终 prev2 包含了爬到第n阶的方法数
	return prev2
}