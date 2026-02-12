package main

import "fmt"

// 计算函数签名
type calcFunc func(...float64) (float64, error)

// 算法类型
type algorithm string

var (
	Sum algorithm = "sum"
	Avg algorithm = "avg"
)

// 计算函数映射
var calcHandlerMap = map[algorithm]calcFunc{
	"sum": sum,
	"avg": average,
}

// 计算函数
func calculate(alg algorithm, params ...float64) (result float64, err error) {
	calcHandler := calcHandlerMap[alg]
	if calcHandler == nil {
		return 0, fmt.Errorf("algorithm not found")
	}

	// 执行计算
	return calcHandler(params...)
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
