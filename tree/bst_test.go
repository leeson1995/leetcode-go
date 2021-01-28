package tree

import "testing"

func Test_IS_BST(t *testing.T) {

}
func Test_insert_into_BST(t *testing.T) {
}

// DFS查找插入位置
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	//中序遍历，检查结果列表是否已经有序
	inOrder(root, &result)
	// check order
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

//  分治法,判断左 MAX < 根 < 右 MIN
type ResultType struct {
	IsValid bool
	// 记录左右两边最大最小值，和根节点进行比较
	Max *TreeNode
	Min *TreeNode
}

func isValidBST2(root *TreeNode) bool {
	result := helper(root)
	return result.IsValid
}
func helper(root *TreeNode) ResultType {
	result := ResultType{}
	// check
	if root == nil {
		result.IsValid = true
		return result
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}
	if left.Max != nil && left.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}

	result.IsValid = true
	// 如果左边还有更小的3，就用更小的节点，不用4
	//  5
	// / \
	// 1   4
	//      / \
	//     3   6
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}
	return result
}
