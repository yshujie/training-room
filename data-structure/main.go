package main

import "fmt"

func main() {
	arr := [6]int{4, 7, 2, 1, 3} // 参数列表

	fmt.Println("--source array: ", arr)
	reverseArray(&arr)
	fmt.Println("reversed array: ", arr)
}
