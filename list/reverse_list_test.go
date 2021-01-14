package list

import (
	"testing"
)

func Test_reverse(t *testing.T) {

	bianli(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for nil != cur {
		tmp := cur.Next
		// 逆置，后节点指向前节点
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

