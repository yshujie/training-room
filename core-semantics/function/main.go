package main

import "fmt"

func main() {
	params := []float64{1, 2, 3, 4, 5} // 参数列表

	// 计算平均值
	fmt.Println("calculate Avg result: ")
	result, err := Calculate(Avg, params...)
	if err != nil {
		fmt.Println("calculate error: ", err)
	} else {
		fmt.Println("Avg result: ", result)
	}

	// 计算求和
	fmt.Println("calculate Sum result: ")
	result, err = Calculate(Sum, params...)
	if err != nil {
		fmt.Println("calculate error: ", err)
	} else {
		fmt.Println("Sum result: ", result)
	}
}
