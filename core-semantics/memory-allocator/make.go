package main

import "fmt"

func testMakeSlice() {
	// 使用 make() 创建 slice
	p := make([]int, 0, 10)

	// ✅ 获取 slice 长度，返回长度为 0
	fmt.Println("len of make slice: ", len(p))

	// ✅ 获取 slice 容量，返回容量为 10
	fmt.Println("cap of make slice: ", cap(p))
}

func testMakeMap() {
	// 使用 make() 创建 map
	pm := make(map[string]int)

	// ✅ 获取 map 长度，返回长度为 0
	fmt.Println("len of make map: ", len(pm))
}

func testMakeChannel() {
	// 使用 make() 创建 channel
	pc := make(chan int)

	// ✅ 获取 channel 长度，返回长度为 0
	fmt.Println("len of make channel: ", len(pc))

	// ✅ 获取 channel 容量，返回容量为 0
	fmt.Println("cap of make channel: ", cap(pc))
}
