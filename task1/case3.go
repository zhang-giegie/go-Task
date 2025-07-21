package main

func IsValid(s string) bool {
	// 创建一个栈
	stack := []rune{}

	// 映射右括号到左括号
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// 如果是右括号
		if leftBracket, exists := bracketMap[char]; exists {
			// 检查栈是否为空或栈顶元素是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != leftBracket {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈
			stack = append(stack, char)
		}
	}

	// 如果栈为空，则括号匹配有效
	return len(stack) == 0
}

// func main() {
// 	s := "()[]{}"
// 	result := IsValid(s)
// 	println(result)
// }
