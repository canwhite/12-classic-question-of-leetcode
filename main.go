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

}