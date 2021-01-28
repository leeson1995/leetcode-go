package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var nums = []int{5, 2, 876, 1, 7, 34}
	fmt.Println(QuickSort(nums))
}
func QuickSort(nums []int) []int {
	// 把一个数组分为左右两段，左段小于右段，类似分治法没有合并过程
	quickSort(nums, 0, len(nums)-1)
	return nums

}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法：divide
		pivot := partition(nums, start, end)
		fmt.Println(nums, pivot)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	//如果遍历指针小于end指针，基准跟遍历指针交换
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			swap(nums, start, i)
			start++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, start, end)
	return start
}

//func swap(nums []int, i, j int) {
//	nums[i], nums[j] = nums[j], nums[i]
//}
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
