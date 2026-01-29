package main

import "fmt"

func testNewSlice() {
	// 使用 new() 函数创建 nil slice
	p := new([]int)

	// ❌ 对 nil slice 索引取值，会出现运行时 panic：index out of range [0] with length 0
	fmt.Println("nil slice item: ", (*p)[0])

	// ✅ 获取 nil slice 长度，返回长度为 0
	fmt.Println("len of nil slice: ", len(*p))

	// ✅ 获取 nil slice 容量，返回容量为 0
	fmt.Println("cap of nil slice: ", cap(*p))
}

func testNewMap() {
	// 使用 new() 创建 nil map
	pm := new(map[string]int)

	// 读取 nil map， 不会 panic，但是会返回类型零值
	fmt.Println("nil map: ", *pm)
	fmt.Println("nil map item: ", (*pm)["January"], "nil map item: ", (*pm)["February"])
}

func testNewChannel() {
	// 使用 new() 创建 nil channel
	pc := new(chan int)

	// 读取 nil channel， 不会 panic，但是会返回类型零值
	fmt.Println("nil channel: ", *pc)
	fmt.Println("nil channel item: ", <-*pc)
}
