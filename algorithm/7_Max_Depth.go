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

    // 递
    leftDepth := MaxDepth(root.Left)
    rightDepth := MaxDepth(root.Right)

    // 归
    if leftDepth > rightDepth {
        return leftDepth + 1
    } else {
        return rightDepth + 1
    }
}

// 迭代解法：BFS层次遍历（可选实现）
func MaxDepthBFS(root *TreeNode) int {
	//退出机制
    if root == nil {
        return 0
    }

    queue := []*TreeNode{root}
    depth := 0

	//开始便利
    for len(queue) > 0 {
		
		//为了便利
        levelSize := len(queue)
		//为了结果
        depth++

		//单层便利和往后进行
        for i := 0; i < levelSize; i++ {

		
            node := queue[0]
            queue = queue[1:]

			//拿新的往后进行
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return depth
}