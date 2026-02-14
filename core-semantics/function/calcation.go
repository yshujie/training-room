package main

import "fmt"

// 算法类型
type algorithm string

// 定义算法类型
const (
	Sum algorithm = "sum"
	Avg algorithm = "avg"
)

// 计算函数签名
type calcFunc func(...float64) (float64, error)

// 计算函数映射
var calcHandlerMap = make(map[algorithm]calcFunc)

// 初始化函数，注册算法对应的计算函数
func init() {
	calcHandlerMap[Sum] = sum
	calcHandlerMap[Avg] = average
}

// 计算函数
func Calculate(alg algorithm, nums ...float64) (result float64, err error) {
	// 加载计算函数
	calcHandler, err := loadCalcHandler(alg)
	if err != nil {
		return 0, err
	}

	// 执行计算
	return calcHandler(nums...)
}

// 根据算法，加载计算函数
func loadCalcHandler(alg algorithm) (handler calcFunc, err error) {
	handler = calcHandlerMap[alg]
	if handler == nil {
		return nil, fmt.Errorf("load calc handler fail: don't set %v algorithm handler", alg)
	}

	return handler, nil
}

// 求和计算
func sum(addends ...float64) (total float64, err error) {
	if len(addends) == 0 {
		return 0, fmt.Errorf("sum failed: invalid parameter, please provide addends parameter")
	}

	for _, addend := range addends {
		total += addend
	}

	return total, nil
}

// 求平均值
func average(nums ...float64) (avg float64, err error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("average failed: invalid parameter, please provide nums parameter")
	}

	var total float64
	for _, num := range nums {
		total += num
	}

	return total / float64(len(nums)), nil
}
