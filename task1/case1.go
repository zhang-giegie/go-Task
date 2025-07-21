package main

func SingleNumber(nums []int) int {
	countMap := make(map[int]int)

	// 记录每个数字出现的次数
	for _, num := range nums {
		countMap[num]++
	}

	// 找到只出现一次的数字
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return -1
}

// func main() {
// 	print(SingleNumber([]int{2, 2, 1}))
// }