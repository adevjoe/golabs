package sort

import (
	"fmt"
	"testing"
)

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

func TestBubbleSort(t *testing.T) {
	a := []int{3, 5, 5, 3, 1, 7, 5, 0, 8, 6, 9, 2}
	fmt.Println(a)
	fmt.Println(bubbleSort(a))
}
