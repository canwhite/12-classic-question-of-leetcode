package algorithm

// 最大子数组和 - Kadane算法
// 动态规划解法
func MaxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // dp[i]表示以nums[i]结尾的最大子数组和
    maxCurrent := nums[0] // 当前最大子数组和
    maxGlobal := nums[0]  // 全局最大子数组和

    for i := 1; i < len(nums); i++ {
        // 关键决策：是否包含前面的子数组
        // 如果前面是负数，不如从当前重新开始
        maxCurrent = max(nums[i], maxCurrent+nums[i])

        // 更新全局最大值
        maxGlobal = max(maxGlobal, maxCurrent)
    }

    return maxGlobal
}

// 标准动态规划解法（更易理解）
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
        // 状态转移方程
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        maxSum = max(maxSum, dp[i])
    }

    return maxSum
}

// 分治法解法（O(n log n)）
func MaxSubArrayDivide(nums []int) int {
    return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {
    if left == right {
        return nums[left]
    }

    mid := left + (right-left)/2

    // 最大子数组在左半部分
    leftMax := maxSubArrayHelper(nums, left, mid)
    // 最大子数组在右半部分
    rightMax := maxSubArrayHelper(nums, mid+1, right)
    // 最大子数组跨越中间
    crossMax := maxCrossingSum(nums, left, mid, right)

    return max(leftMax, max(rightMax, crossMax)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
    // 从中间向左找最大和
    leftSum := nums[mid]
    maxLeftSum := leftSum
    for i := mid - 1; i >= left; i-- {
        leftSum += nums[i]
        maxLeftSum = max(maxLeftSum, leftSum)
    }

    // 从中间向右找最大和
    rightSum := nums[mid+1]
    maxRightSum := rightSum
    for i := mid + 2; i <= right; i++ {
        rightSum += nums[i]
        maxRightSum = max(maxRightSum, rightSum)
    }

    return maxLeftSum + maxRightSum
}

// 辅助函数
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}