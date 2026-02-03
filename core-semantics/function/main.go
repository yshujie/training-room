package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := T{}
	dumpMethodSet(t)

	prt := &T{}
	dumpMethodSet(prt)

	s := S{}
	dumpMethodSet(s)

	prs := &S{}
	dumpMethodSet(prs)
}

type T struct{}

func (t T) M1() {}
func (t T) M2() {}

func (t *T) M3() {}
func (t *T) M4() {}

type S T

// dumpMethodSet 打印参数类型的方法集
func dumpMethodSet(i interface{}) {
	// 获取参数类型的反射类型
	dynTyp := reflect.TypeOf(i)
	fmt.Printf("---- start: %v ---- \n", dynTyp)
	// 如果反射类型为 nil，则打印错误信息
	if dynTyp == nil {
		fmt.Println("There is no type to dump")
		return
	}

	// 获取反射类型的方法数量
	numMethod := dynTyp.NumMethod()
	if 0 == numMethod {
		fmt.Println("There is no method of dynTyp")
		return
	}

	// 打印方法列表
	for i := 0; i < numMethod; i++ {
		fmt.Printf("%v index method is %v \n", i, dynTyp.Method(i).Name)
	}

	fmt.Printf("---- end:%v ----\n", dynTyp)
}
