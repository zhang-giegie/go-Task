package main

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := fmt.Sprintf("%d", x) // 将整数转为字符串
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPalindrome(121))   // 输出：true
// 	fmt.Println(IsPalindrome(-121))  // 输出：false
// 	fmt.Println(IsPalindrome(10))    // 输出：false
// 	fmt.Println(IsPalindrome(12321)) // 输出：true
// }