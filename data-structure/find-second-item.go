package main

import (
	"fmt"
	"math"
)

/*
*
✅ 题目：查找数组中第二大的元素
实现一个函数，返回数组中第二大的数字（要求一次遍历完成）。

示例
输入：[3, 1, 4, 5, 2]
输出：4

考点
  - 一次遍历维护两个变量 max1、max2
  - 边界值判断（数组长度 < 2）
*/
func findSecondlargestItem(a []int) (int, error) {
	if len(a) < 2 {
		return 0, fmt.Errorf("find second largest item fail: lenght of params less than 2")
	}

	max1, max2 := math.MinInt, math.MinInt
	for _, v := range a {
		if v > max1 {
			max2 = max1
			max1 = v
		} else if v == max1 {
			continue
		} else if v > max2 {
			max2 = v
		}
	}

	if max2 == math.MinInt {
		return 0, fmt.Errorf("find second largest item fail: all item is same")
	}

	return max2, nil
}
