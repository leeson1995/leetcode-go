package tree

import "testing"

func Test_Balance(t *testing.T) {

}

//左边平衡 && 右边平衡 && 左右两边高度 <= 1
func isBalanced(root *TreeNode) bool {
	if maxDepth2(root) == -1 {
		return false
	}
	return true
}
func maxDepth2(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
