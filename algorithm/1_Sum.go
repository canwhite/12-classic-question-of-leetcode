package algorithm

/**
用空间换时间和目标思考的典范
*/
import (
	"sort"
)


func Sum(nums []int, target int) []int {
	// 使用map存储已经遍历过的数字及其索引
	// key: 数字, value: 索引, 这种事根据我们的目的，反向思考，放的value事索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 检查目标数字 - 当前数字是否存在于map中
		complement := target - num
		if j, exists := numMap[complement]; exists {
			// 找到了，返回两个数的索引
			return []int{j, i}
		}

		// 将当前数字存入map，供后续查找使用
		numMap[num] = i
	}

	// 没有找到
	return []int{}
}

// 三数之和 - 减少变量，固定一个，寻找另外两个
// 时间复杂度：O(n²)，空间复杂度：O(1)（不包括结果存储）
func SumThree(nums []int) [][]int {
	sort.Ints(nums) // 先排序，为双指针法做准备
	var result [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// 去重：跳过重复的第一个元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 如果当前数大于0，后面不可能有三数之和为0
		if nums[i] > 0 {
			break
		}

		// 双指针法
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 去重：跳过重复的第二个和第三个元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				// 和太小，移动左指针
				left++
			} else {
				// 和太大，移动右指针
				right--
			}
		}
	}

	return result
}


// 三数之和最接近 - 找出最接近target的三数之和
// 时间复杂度：O(n²)，空间复杂度：O(1)
func SumThreeClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)

	// 初始化最接近的和
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		// 避免重复计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			current := nums[i] + nums[left] + nums[right]

			// 更新最接近的和
			if abs(current-target) < abs(closest-target) {
				closest = current
			}

			// 移动指针
			if current < target {
				left++
			} else if current > target {
				right--
			} else {
				// 完全匹配，直接返回
				return current
			}
		}
	}

	return closest
}

// 三数之和小于 - 统计满足和小于target的三元组数量
// 时间复杂度：O(n²)，空间复杂度：O(1)
func SumThreeSmaller(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			current := nums[i] + nums[left] + nums[right]
			if current < target {
				// 因为数组已排序，如果当前和小于target，
				// 那么从left到right-1的所有组合都满足条件
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}