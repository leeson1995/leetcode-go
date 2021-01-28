package tree

import (
	"fmt"
	"testing"
)

func Test_traverse(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{9, &TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	(RecursionPreorderTraversal(root))

	fmt.Println(preorderTraversal(root))
}

// V3：通过非递归遍历
func preorderTraversal(root *TreeNode) []int {
	// 非递归
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}
func RecursionPreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	RecursionPreorderTraversal(root.Left)
	RecursionPreorderTraversal(root.Right)
}
