package algorithm

/*
递归是种基础操作，这里主要理解递归法去调转链表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}


// 反转链表 - 递归法，这种事主要要讲的，我之前看混淆了
// 时间复杂度：O(n)，空间复杂度：O(n)（递归栈）
func ReverseListRecursive(head *ListNode) *ListNode {
	// 先直接钻到底，如果到nil返回4
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转剩余部分
	newHead := ReverseListRecursive(head.Next)

	// 反转当前节点和下一个节点的连接
	head.Next.Next = head
	head.Next = nil

	return newHead
}


// 反转链表 - 迭代法（推荐），这种迭代法是很好理解的
// 正常操作的时候，我们也会用for循环制定跳出机制，然后去迭代
// 时间复杂度：O(n)，空间复杂度：O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		// 保存下一个节点
		nextTemp := curr.Next
		// 反转当前节点的指向
		curr.Next = prev
		// 移动指针
		prev = curr
		curr = nextTemp
	}

	return prev
}



// 反转链表的前N个节点
func ReverseListFirstN(head *ListNode, n int) *ListNode {
	var successor *ListNode

	var reverseN func(*ListNode, int) *ListNode
	reverseN = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			// 记录第n+1个节点
			successor = head.Next
			return head
		}

		// 递归反转前n-1个节点
		newHead := reverseN(head.Next, n-1)

		// 连接第n个节点和第n+1个节点
		head.Next.Next = head
		head.Next = successor

		return newHead
	}

	return reverseN(head, n)
}

// 反转链表的第m到n个节点
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return ReverseListFirstN(head, n)
	}

	// 前进到第m个节点
	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}

// 创建链表的辅助函数（用于测试）
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// 打印链表的辅助函数（用于测试）
func PrintList(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// 链表长度
func listLength(head *ListNode) int {
	length := 0
	current := head

	for current != nil {
		length++
		current = current.Next
	}

	return length
}