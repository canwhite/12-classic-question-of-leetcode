package algorithm

// ⚠️注意：迭代法是用循环的方式遍历
// MergeTwoListsIterative 合并两个有序链表 - 迭代法（空间O(1)）
// 迭代法：使用循环重复执行相同逻辑，通过循环变量控制过程
// 优点：空间复杂度最优，无栈溢出风险
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，简化边界处理
	dummy := &ListNode{}
	current := dummy

	// 比较两个链表的节点，选择较小的节点
	// 这是核心的迭代逻辑：循环处理直到其中一个链表为空
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 处理剩余的节点（一个链表已经遍历完）

	// 剩余的节点本身就是有序的，直接连接即可
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// 返回虚拟头节点的下一个节点，即合并后的链表头
	return dummy.Next
}

// MergeTwoListsRecursive 合并两个有序链表 - 递归法（空间O(n)）
// 递归法：函数调用自身来解决问题，将大问题分解成小问题
// 优点：代码简洁，逻辑直观
// 缺点：空间复杂度较高，可能栈溢出（长链表）
func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	// 基础情况：如果其中一个链表为空，直接返回另一个链表
	// 这是递归的终止条件
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 递归情况：比较当前节点，选择较小的作为合并结果的头节点
	if list1.Val <= list2.Val {
		// 选择 list1 的当前节点，然后递归处理 list1.Next 和 list2
		// list1 的 Next 指向递归合并的结果
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		// 选择 list2 的当前节点，然后递归处理 list1 和 list2.Next
		// list2 的 Next 指向递归合并的结果
		list2.Next = MergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}