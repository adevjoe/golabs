package sort

import (
	"fmt"
	"testing"
)

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

func TestSelection(t *testing.T) {
	a := []int{3, 5, 5, 3, 1, 7, 5, 0, 8, 6, 9, 2}
	fmt.Println(a)
	fmt.Println(selectionSort(a))
}
