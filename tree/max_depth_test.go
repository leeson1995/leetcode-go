package tree

import (
	"fmt"
	"testing"
)

func Test_Max_depth(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, &TreeNode{4, nil, nil}, nil},
			&TreeNode{7, nil, nil}}}

	//fmt.Println(levelOrder(root))
	fmt.Println(maxDepth(root))
}
func maxDepth(root *TreeNode) int {
	// 返回条件处理
	if root == nil {
		return 0
	}
	// 分治
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
