 package algorithm

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
    // 递归终止条件：空节点深度为0
    if root == nil {
        return 0
    }

    // 递归计算左右子树的深度
    leftDepth := MaxDepth(root.Left)
    rightDepth := MaxDepth(root.Right)

    // 当前节点的深度 = max(左子树深度, 右子树深度) + 1
    if leftDepth > rightDepth {
        return leftDepth + 1
    } else {
        return rightDepth + 1
    }
}

// 迭代解法：BFS层次遍历（可选实现）
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }

//     queue := []*TreeNode{root}
//     depth := 0

//     for len(queue) > 0 {
//         // 处理当前层的所有节点
//         levelSize := len(queue)
//         depth++

//         for i := 0; i < levelSize; i++ {
//             node := queue[0]
//             queue = queue[1:]

//             if node.Left != nil {
//                 queue = append(queue, node.Left)
//             }
//             if node.Right != nil {
//                 queue = append(queue, node.Right)
//             }
//         }
//     }

//     return depth
// }