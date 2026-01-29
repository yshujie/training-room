package main

func main() {
	// new() 函数与 slice/map/channel 的语义与 nil 行为
	testNewSlice()
	testNewMap()
	testNewChannel()

	// make() 函数与 slice/map/channel 的语义与 nil 行为
	testMakeSlice()
	testMakeMap()
	testMakeChannel()
}
