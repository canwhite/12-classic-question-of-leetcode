package algorithm

func ClimbStairs(n int) int {
    if n <= 0 { return 0 }
    if n == 1 { return 1 }
    if n == 2 { return 2 }

    // dp[i] 表示爬到第i阶的方法数
    dp := make([]int, n+1)
    dp[0] = 0  // 0阶，0种方法
    dp[1] = 1  // 1阶，1种方法
    dp[2] = 2  // 2阶，2种方法

    // 从第3阶开始递推
    for i := 3; i <= n; i++ {
		//TODO:主要是dp[i]的时候和前边dp组合的关系，可以先猜，然后再验证
        dp[i] = dp[i-1] + dp[i-2]  // 斐波那契递推
    }

    return dp[n]
}