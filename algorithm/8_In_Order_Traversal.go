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

    var traverse func(*TreeNode)
    traverse = func(node *TreeNode) {
        // 递归终止条件
        if node == nil {
            return
        }

        // 中序遍历：左 → 根 → 右
        traverse(node.Left)       // 先遍历左子树
        result = append(result, node.Val)  // 再处理当前节点
        traverse(node.Right)      // 最后遍历右子树
    }
    
    traverse(root)
    return result
}

// 迭代解法：使用栈
func inorderTraversalIterative(root *TreeNode) []int {
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