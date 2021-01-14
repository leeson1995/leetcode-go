package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func bianli(head *ListNode) {
	if head != nil {
		fmt.Println(head.Val)
		bianli(head.Next)
	}
}
