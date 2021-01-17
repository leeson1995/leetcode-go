package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

// 703 数据流中k大元素

func Test_kth_Largest(t *testing.T) {
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(&k.H, k.H)
}

type KthLargest struct {
	k int
	H *myHeap
}

func (l *KthLargest) Add(i int) int {
	if l.H.Len() < l.k {
		heap.Push(l.H, i)
	} else if (*l.H)[0] < i {
		heap.Pop(l.H)
		heap.Push(l.H, i)
	}
	return (*l.H)[0]
}

//小顶堆
type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]

}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func Constructor(k int, nums []int) KthLargest {
	h := &myHeap{}
	heap.Init(h)
	kth := KthLargest{k, h}
	for _, n := range nums {
		kth.Add(n)

	}
	return kth

}
