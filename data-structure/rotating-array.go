package main

/**
✅ 题目：旋转数组（左旋转）
给定一个长度为 n 的数组 arr 和一个整数 k，将数组中的元素向左旋转 k 个位置（原地修改）。

示例
输入：[1, 2, 3, 4, 5, 6, 7, 8, 9], k = 3
输出：[4, 5, 6, 7, 8, 9, 1, 2, 3]

考点
	•	数组索引运算（取模）
	•	原地反转的三步法
	•	时间复杂度 O(n)、空间复杂度 O(1)
*/

func rotationLeft(a []int, k int) {
	if len(a) < 2 {
		return
	}

	// 防止 k 越界
	k = k % len(a)

	// 左反转，则旋转轴心点为 len(a) - k
	z := len(a) - k

	// 数组整体反转
	reverseSlice(a[:])

	// 前半段反转
	reverseSlice(a[0:z])

	// 后半段反转
	reverseSlice(a[z:])
}

/*
*
✅ 题目：旋转数组（右旋转）
给定一个长度为 n 的数组 arr 和一个整数 k，将数组中的元素向右旋转 k 个位置（原地修改）。

示例
输入：[1, 2, 3, 4, 5, 6, 7], k = 3
输出：[5, 6, 7, 1, 2, 3, 4]

考点
  - 数组索引运算（取模）
  - 原地反转的三步法
  - 时间复杂度 O(n)、空间复杂度 O(1)
*/
func rotationRight(a []int, k int) {
	if len(a) < 2 {
		return
	}

	// 防止 k 索引越界，对 k 取模
	k = k % len(a)

	// 右选择，旋转的轴心的为 k
	z := k

	// 三步反转法
	reverseSlice(a)
	reverseSlice(a[:z])
	reverseSlice(a[z:])

}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
