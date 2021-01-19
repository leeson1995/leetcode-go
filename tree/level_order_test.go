package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	//fmt.Println(levelOrder(root))
	fmt.Println(bfs(root))
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	if root == nil {
		return nil
	}

	dfs(root, 0)

	return res
}

//人用循环，神用递归
func dfs(root *TreeNode, level int) {

	if root == nil {
		return
	}

	//如果第一次入当前层，初始化
	if level == len(res) {
		res = append(res, []int{})
	}
	//否则，直接追加
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)

}

// 核心在于对queue的操作。每次循环把当前节点queue[0]的左右子树 都放到queue里（queue[1]，queue[2]）
// 然后每次循环都剔除第0个元素，相当于遍历所有节点。
func bfs(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
