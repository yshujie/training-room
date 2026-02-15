package main

/*
✅ 题目：找出切片中最大连续递增子序列的长度
编写函数 MaxIncreasingSubseq(nums []int) int，找出切片中最长的连续递增子序列长度。

输入示例：
nums := []int{1, 2, 2, 3, 4, 1, 2, 3, 4, 5}
fmt.Println(MaxIncreasingSubseq(nums))

输出：
5
*/

func MaxIncreasingSubseq(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	maxLen := 0
	winLen := 1
	for idx := 1; idx < len(nums); idx++ {
		// 子切片递增中，窗口持续增加
		if nums[idx] >= nums[idx-1] {
			winLen++
			continue
		}

		// 子切片递增中断
		// 判断当前窗口长度，超过之前则的保留
		if winLen > maxLen {
			maxLen = winLen
		}

		// 重置窗口
		winLen = 1
	}

	// 判断最后一个滑动窗口长度
	if winLen > maxLen {
		maxLen = winLen
	}

	return maxLen
}
