package main

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0 // j指针用于记录不重复元素的位置

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1 // 返回不重复元素的个数
}

// func main() {
// 	nums := []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9, 9, 10}
// 	length := RemoveDuplicates(nums)
// 	fmt.Println("不重复的元素个数:", length)
// 	fmt.Println("去重后的数组:", nums[:length])
// }
