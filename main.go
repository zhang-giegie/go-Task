package main

import (
	"fmt"

	"github.com/zhang-giegie/go-Task/task1"
)

func main() {
	// print(task1.SingleNumber([]int{2, 2, 1}))

	// print(task1.IsPalindrome(121))

	// print(task1.IsValid("()[]{}"))

	// print(task1.LongestCommonPrefix([]string{"flower","flow","flight"}))

	// fmt.Print(task1.PlusOne([]int{1, 2, 3}))

	// fmt.Print(task1.RemoveDuplicates([]int{0,0,1,1,1,2,2,3,3,0}))

	// intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// fmt.Println(task1.Merge(intervals1))
	nums1 := []int{2, 7, 11, 15}
	fmt.Println(task1.TwoSum(nums1, 9))

}
