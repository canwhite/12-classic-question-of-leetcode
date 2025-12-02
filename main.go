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

}