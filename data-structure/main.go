package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 参数列表

	fmt.Println("--source array: ", arr)
	rotationRight(arr, 3)
	fmt.Println("reversed array: ", arr)
}
