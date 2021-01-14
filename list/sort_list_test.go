package list

import "testing"

func TestSortList(t *testing.T) {

	bianli(sortList(&ListNode{5, &ListNode{4, &ListNode{8,
		&ListNode{2, &ListNode{79, &ListNode{-3, nil}}}}}}))
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	var zhong *ListNode
	for fast != nil && fast.Next != nil {
		zhong = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//相当于left.next 指向空
	zhong.Next = nil
	left := sortList(head)
	right := sortList(slow)
	return mergeList(left, right)
}

func mergeList(l, r *ListNode) *ListNode {

	merge := &ListNode{0, nil}
	prev := merge

	for l != nil && r != nil {
		if l.Val < r.Val {
			prev.Next = l
			l = l.Next
		} else {
			prev.Next = r
			r = r.Next
		}
		prev = prev.Next
	}
	if l != nil {
		prev.Next = l
	}
	if r != nil {
		prev.Next = r
	}
	return merge.Next
}
func mergeList2(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
	l1.Next = mergeList(l2, l1.Next)
	return l1
}
