package main

func addTen(num *int) {
	*num += 10
}

// 定义一个函数，接受一个整数切片的指针
func multiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2 // 将每个元素乘以 2
	}
}

// func main() {
// 	value := 5
// 	fmt.Println("原始值:", value)
// 	addTen(&value)
// 	fmt.Println("增加 10 后的值:", value)

//     numbers := []int{1, 2, 3, 4, 5}
//     fmt.Println("原始切片:", numbers)

//     // 调用函数并传递切片的地址
//     multiplyByTwo(&numbers)

//     // 输出修改后的切片
//     fmt.Println("修改后的切片:", numbers)
// }
