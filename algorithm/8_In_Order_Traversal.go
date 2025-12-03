/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InorderTraversal(root *TreeNode) []int {
    var result []int

    //声明
    var traverse func(*TreeNode)
    //实现
    traverse = func(node *TreeNode) {
        // 递归终止条件
        if node == nil {
            return
        }
        traverse(node.Left)       // 递：向左递归
        result = append(result, node.Val)  // 归：左子树返回时执行
        traverse(node.Right)      // 在归的过程中，又开始对右节点递归
    }

    traverse(root)
    return result
}

// 迭代解法：使用栈
func InorderTraversalIterative(root *TreeNode) []int {
    var result []int
    var stack []*TreeNode
    current := root

    for current != nil || len(stack) > 0 {
        // 先一直往左走，把所有左节点入栈
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }

        // 弹出栈顶元素，处理当前节点
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)

        // 转向右子树
        current = current.Right
    }

    return result
}