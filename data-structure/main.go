package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 5, 2, 8} // 参数列表

	res, _ := findSecondlargestItem(arr)
	fmt.Println("second largerst item : ", res)
}
