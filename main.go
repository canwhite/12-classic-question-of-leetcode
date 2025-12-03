package main

import (
	"fmt"
	"12-classic-question-of-leetcode/algorithm"
)

func main() {
	// 测试两数之和算法
	nums := []int{2, 7, 11, 15}
	target := 9


	fmt.Println("\n=== 两数之和测试 ===")
	result := algorithm.Sum(nums, target)
	fmt.Printf("两数之和结果: %v\n", result) // 应该输出 [0, 1]


	// 测试三数之和算法
	fmt.Println("\n=== 三数之和测试 ===")

	// 测试三数之和（和为0）
	threeNums := []int{-1, 0, 1, 2, -1, -4}
	threeResult := algorithm.SumThree(threeNums)
	fmt.Printf("三数之和为0的结果: %v\n", threeResult)

	// 测试三数之和最接近
	target = 1
	closest := algorithm.SumThreeClosest(threeNums, target)
	fmt.Printf("最接近%d的三数之和: %d\n", target, closest)

	// 测试三数之和小于target
	smallerCount := algorithm.SumThreeSmaller(threeNums, target)
	fmt.Printf("和小于%d的三元组数量: %d\n", target, smallerCount)



	// 测试反转链表算法
	fmt.Println("\n=== 反转链表测试 ===")

	// 创建链表: 1 -> 2 -> 3 -> 4 -> 5
	list := algorithm.CreateList([]int{1, 2, 3, 4, 5})
	fmt.Printf("原始链表: %v\n", algorithm.PrintList(list))

	// 反转整个链表
	reversed := algorithm.ReverseList(list)
	fmt.Printf("反转后链表: %v\n", algorithm.PrintList(reversed))

	// 测试反转前N个节点
	newList := algorithm.CreateList([]int{1, 2, 3, 4, 5})
	reversedFirst3 := algorithm.ReverseListFirstN(newList, 3)
	fmt.Printf("反转前3个节点: %v\n", algorithm.PrintList(reversedFirst3))

	// 测试反转第m到n个节点
	anotherList := algorithm.CreateList([]int{1, 2, 3, 4, 5})
	reversedBetween := algorithm.ReverseBetween(anotherList, 2, 4)
	fmt.Printf("反转第2到4个节点: %v\n", algorithm.PrintList(reversedBetween))

	// 测试有效括号算法
	fmt.Println("\n=== 有效括号测试 ===")

	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"[", false},
		{"]", false},
		{"{[()()]}", true},
		{"{[(])}", false},
	}

	fmt.Println("测试有效括号算法:")
	for i, tc := range testCases {
		result := algorithm.IsValid(tc.input)
		status := "✓"
		if !result {
			status = "✗"
		}
		fmt.Printf("Test %d: Input: %-12s Expected: %-5v Got: %-5v %s\n",
			i+1, fmt.Sprintf("\"%s\"", tc.input), tc.expected, result, status)
	}

	// 测试合并两个有序链表算法
	fmt.Println("\n=== 合并两个有序链表测试 ===")

	// 辅助函数：创建链表
	createList := func(nums []int) *algorithm.ListNode {
		if len(nums) == 0 {
			return nil
		}
		head := &algorithm.ListNode{Val: nums[0]}
		current := head
		for _, num := range nums[1:] {
			current.Next = &algorithm.ListNode{Val: num}
			current = current.Next
		}
		return head
	}

	// 辅助函数：打印链表
	printList := func(head *algorithm.ListNode) []int {
		var result []int
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	mergeTestCases := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "两个非空链表",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "一个空链表，一个非空",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
		{
			name:     "两个空链表",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "链表1全部小于链表2",
			list1:    []int{1, 2, 3},
			list2:    []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "链表2全部小于链表1",
			list1:    []int{4, 5, 6},
			list2:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "交替排列",
			list1:    []int{1, 3, 5},
			list2:    []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "长度不同的链表",
			list1:    []int{1, 3, 5, 7},
			list2:    []int{2, 4},
			expected: []int{1, 2, 3, 4, 5, 7},
		},
	}

	fmt.Println("测试合并两个有序链表算法:")
	for i, tc := range mergeTestCases {
		// 创建两个有序链表
		l1 := createList(tc.list1)
		l2 := createList(tc.list2)

		// 合并链表
		merged := algorithm.MergeTwoLists(l1, l2)
		result := printList(merged)

		// 检查结果
		status := "✓"
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", tc.expected) {
			status = "✗"
		}

		fmt.Printf("Test %d (%-25s): L1:%v L2:%v → %v %s\n",
			i+1, tc.name, tc.list1, tc.list2, result, status)
	}

	// 测试爬楼梯算法
	fmt.Println("\n=== 爬楼梯测试 ===")

	// 爬楼梯问题：每次可以爬1阶或2阶，问有多少种方法爬到n阶
	climbStairsTestCases := []struct {
		n        int
		expected int
		methods  []string
	}{
		{
			n:        0,
			expected: 0,
			methods:  []string{},
		},
		{
			n:        1,
			expected: 1,
			methods:  []string{"1"},
		},
		{
			n:        2,
			expected: 2,
			methods:  []string{"1+1", "2"},
		},
		{
			n:        3,
			expected: 3,
			methods:  []string{"1+1+1", "1+2", "2+1"},
		},
		{
			n:        4,
			expected: 5,
			methods:  []string{"1+1+1+1", "1+1+2", "1+2+1", "2+1+1", "2+2"},
		},
		{
			n:        5,
			expected: 8,
			methods:  []string{"1+1+1+1+1", "1+1+1+2", "1+1+2+1", "1+2+1+1", "1+2+2", "2+1+1+1", "2+1+2", "2+2+1"},
		},
		{
			n:        10,
			expected: 89,
			methods:  []string{"(太多方法，略)..."},
		},
	}

	fmt.Println("测试爬楼梯算法（动态规划版）:")
	for i, tc := range climbStairsTestCases {
		result := algorithm.ClimbStairs(tc.n)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}

		// 显示部分方法，避免输出太长
		methodsStr := ""
		if len(tc.methods) <= 3 {
			for j, method := range tc.methods {
				if j > 0 {
					methodsStr += ", "
				}
				methodsStr += method
			}
		} else {
			methodsStr = tc.methods[0] + ", " + tc.methods[1] + ", ... (共" + fmt.Sprintf("%d", len(tc.methods)) + "种)"
		}

		fmt.Printf("Test %d: n=%-2d → 方法数=%-3d %s (预期:%d) %s\n",
			i+1, tc.n, result, methodsStr, tc.expected, status)
	}

	fmt.Println("\n📊 斐波那契数列规律:")
	fmt.Printf("f(1)=1, f(2)=2, f(3)=3, f(4)=5, f(5)=8, f(6)=13, f(7)=21, f(8)=34, f(9)=55, f(10)=89\n")
	fmt.Println("每一项 = 前两项之和，这就是斐波那契数列！")

	fmt.Println("\n🎯 算法思想:")
	fmt.Println("1. 递归关系：f(n) = f(n-1) + f(n-2)")
	fmt.Println("2. 为什么要加？因为到第n阶有两种方式:")
	fmt.Println("   - 从第n-1阶爬1步")
	fmt.Println("   - 从第n-2阶爬2步")
	fmt.Println("3. 边界条件：f(0)=0, f(1)=1, f(2)=2")
	fmt.Println("4. 时间复杂度：O(n)，空间复杂度：O(1)")

	// 测试买卖股票最佳时机算法
	fmt.Println("\n=== 买卖股票最佳时机测试 ===")

	maxProfitTestCases := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "单调上涨",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "单调下跌",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "波动上涨",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "多次买卖机会",
			prices:   []int{1, 2, 3, 4, 5, 1, 2, 3, 4},
			expected: 9,
		},
		{
			name:     "单一价格",
			prices:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "空数组",
			prices:   []int{},
			expected: 0,
		},
	}

	fmt.Println("测试动态规划解法:")
	for i, tc := range maxProfitTestCases {
		result := algorithm.MaxProfit(tc.prices)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d (%-15s): Prices:%v → Profit:%d Expected:%d %s\n",
			i+1, tc.name, tc.prices, result, tc.expected, status)
	}

	fmt.Println("\n测试贪心解法（无限次交易）:")
	for i, tc := range maxProfitTestCases {
		result := algorithm.MaxProfitUnlimited(tc.prices)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d (%-15s): Prices:%v → Profit:%d Expected:%d %s\n",
			i+1, tc.name, tc.prices, result, tc.expected, status)
	}

	fmt.Println("\n测试贪心解法（含手续费）:")
	maxProfitFeeTestCases := []struct {
		name     string
		prices   []int
		fee      int
		expected int
	}{
		{
			name:     "简单案例",
			prices:   []int{1, 3, 2, 8, 4, 9},
			fee:      2,
			expected: 8,
		},
		{
			name:     "手续费过高",
			prices:   []int{1, 3, 7},
			fee:      10,
			expected: 0,
		},
		{
			name:     "多次交易",
			prices:   []int{1, 5, 3, 6, 4, 5, 2, 8},
			fee:      1,
			expected: 11,
		},
	}

	for i, tc := range maxProfitFeeTestCases {
		result := algorithm.MaxProfitGreedy(tc.prices, tc.fee)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d (%-15s): Prices:%v Fee:%d → Profit:%d Expected:%d %s\n",
			i+1, tc.name, tc.prices, tc.fee, result, tc.expected, status)
	}

	fmt.Println("\n🎯 买卖股票算法思想:")
	fmt.Println("1. 动态规划：记录持股和不持股两种状态")
	fmt.Println("2. 贪心算法：抓住每个盈利机会，设置最低买入价格")
	fmt.Println("3. 手续费版本：只有利润超过手续费才进行交易")
	fmt.Println("4. 时间复杂度：O(n)，空间复杂度：O(1)")

	// 测试二叉树最大深度算法
	fmt.Println("\n=== 二叉树最大深度测试 ===")

	// 辅助函数：创建二叉树
	createTree := func(vals []int) *algorithm.TreeNode {
		if len(vals) == 0 {
			return nil
		}

		nodes := make([]*algorithm.TreeNode, len(vals))
		for i, val := range vals {
			if val != -1 { // -1 表示空节点
				nodes[i] = &algorithm.TreeNode{Val: val}
			}
		}

		for i := 0; i < len(vals); i++ {
			if nodes[i] != nil {
				leftIndex := 2*i + 1
				rightIndex := 2*i + 2
				if leftIndex < len(vals) {
					nodes[i].Left = nodes[leftIndex]
				}
				if rightIndex < len(vals) {
					nodes[i].Right = nodes[rightIndex]
				}
			}
		}

		return nodes[0]
	}

	maxDepthTestCases := []struct {
		name     string
		treeVals []int // -1 表示空节点
		expected int
	}{
		{
			name:     "空树",
			treeVals: []int{},
			expected: 0,
		},
		{
			name:     "只有根节点",
			treeVals: []int{1},
			expected: 1,
		},
		{
			name:     "完全二叉树",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7},
			expected: 3,
		},
		{
			name:     "左子树较深",
			treeVals: []int{1, 2, -1, 3, -1, -1, -1, 4},
			expected: 4,
		},
		{
			name:     "右子树较深",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			expected: 3,
		},
		{
			name:     "平衡二叉树",
			treeVals: []int{3, 9, 20, -1, -1, 15, 7},
			expected: 3,
		},
		{
			name:     "单链表式树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			expected: 3,
		},
	}

	fmt.Println("测试二叉树最大深度算法:")
	for i, tc := range maxDepthTestCases {
		root := createTree(tc.treeVals)
		result := algorithm.MaxDepth(root)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}

		// 简化树的可视化显示
		treeStr := fmt.Sprintf("%v", tc.treeVals)
		fmt.Printf("Test %d (%-20s): Tree:%s → Depth:%d Expected:%d %s\n",
			i+1, tc.name, treeStr, result, tc.expected, status)
	}

	fmt.Println("\n🎯 二叉树最大深度算法思想:")
	fmt.Println("1. 递归思路：树的高度 = max(左子树高度, 右子树高度) + 1")
	fmt.Println("2. 终止条件：空节点高度为0")
	fmt.Println("3. 分治策略：将问题分解为左右子树的子问题")
	fmt.Println("4. 时间复杂度：O(n)，空间复杂度：O(h) h为树高")

}