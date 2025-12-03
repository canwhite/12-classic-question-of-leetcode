package algorithm

// 方法1：单调栈解法
// 思路：维护一个单调递减的栈，遇到高柱子时计算接水量
// 时间复杂度：O(n)，空间复杂度：O(n)
func Trap(height []int) int {
	var stack []int // 存储柱子索引的栈
	result := 0   // 累计接水量

	// 遍历每个柱子
	for i := 0; i < len(height); i++ {
		// 当前柱子比栈顶柱子高，说明可能形成凹槽
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]        // 凹槽底部
			stack = stack[:len(stack)-1]     // 弹出栈顶

			// 栈为空，无法形成凹槽
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]   // 凹槽左边
			width := i - left - 1        // 凹槽宽度
			water := min(height[i], height[left]) - height[top] // 凹槽深度
			result += width * water     // 累加接水量
		}

		//为什么当前柱子入栈在这个位置
		stack = append(stack, i) // 当前柱子入栈
	}

	return result
}

// 方法2：动态规划解法
// 思路：预处理每个位置左边和右边的最大高度，计算接水量
// 时间复杂度：O(n)，空间复杂度：O(n)
// 区别于状态即答案，这里的状态是辅助信息
func Trap2(height []int) int {
    n := len(height)
    if n == 0 { return 0 }

    // 第一维状态：左边最高
    dpLeft := make([]int, n)
    dpLeft[0] = height[0]
    for i := 1; i < n; i++ {
        dpLeft[i] = max(dpLeft[i-1], height[i])    // 状态转移
    }

    // 第二维状态：右边最高
    dpRight := make([]int, n)
    dpRight[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        dpRight[i] = max(dpRight[i+1], height[i])  // 状态转移
    }

    // 用状态计算答案
    ans := 0
    for i := 0; i < n; i++ {
        water := min(dpLeft[i], dpRight[i]) - height[i]
        if water > 0 {
            ans += water
        }
    }
    return ans
}

// 方法3：双指针解法
// 思路：左右指针移动，维护左右最大高度，避免空间数组
// 时间复杂度：O(n)，空间复杂度：O(1)
func Trap3(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	// 左右指针向中间移动
	for left <= right {
		if height[left] <= height[right] {
			// 处理左指针
			if height[left] >= leftMax {
				leftMax = height[left] // 更新左最大高度
			} else {
				result += leftMax - height[left] // 累加接水量
			}
			left++
		} else {
			// 处理右指针
			if height[right] >= rightMax {
				rightMax = height[right] // 更新右最大高度
			} else {
				result += rightMax - height[right] // 累加接水量
			}
			right--
		}
	}

	return result
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}