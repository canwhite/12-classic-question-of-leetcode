package algorithm


// 标准动态规划解法（更易理解），注意是最大子数组和
func MaxSubArrayDP(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    // dp[i] = 以nums[i]结尾的最大子数组和
    dp := make([]int, n)
    dp[0] = nums[0]
    maxSum := dp[0]

    for i := 1; i < n; i++ {
        // 状态转移方程,这个很好理解，nums[i]本身，或者之前的最大子数组和加上当前
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        maxSum = max(maxSum, dp[i])
    }

    return maxSum
}