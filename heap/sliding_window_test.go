package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

// 239 滑动窗口最大值

// TODO 少个数 输出；[3,5,5,6,7] 预期：[3,3,5,5,6,7]
func TestSliding_window(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	max := &MaxHeap{}
	heap.Init(max)
	res := []int{}
	for _, n := range nums {

		if max.Len() < k {
			heap.Push(max, (n))

		} else {

			heap.Remove(max, max.Len()-1)
			heap.Push(max, (n))
			res = append(res, max.Get().(int))
			fmt.Println(max, max.Get())

		}
		//fmt.Println(max, max.Get())

	}
	return res
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, (x.(int)))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	l := len(old)
	*m = old[:l-1]
	v := old[l-1]
	return v
}
func (m *MaxHeap) Get() interface{} {
	return (*m)[0]
}
