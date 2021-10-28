package sort

import (
	"fmt"
	"testing"
)

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		pre := i - 1
		cur := nums[i]
		for pre >= 0 && nums[pre] > cur {
			nums[pre+1] = nums[pre]
			pre--
		}
		nums[pre+1] = cur
	}
	return nums
}

func TestInsertion(t *testing.T) {
	a := []int{3, 5, 5, 3, 1, 7, 5, 0, 8, 6, 9, 2}
	fmt.Println(a)
	fmt.Println(insertionSort(a))
}
