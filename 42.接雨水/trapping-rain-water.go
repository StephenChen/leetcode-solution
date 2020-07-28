/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (51.65%)
 * Likes:    1491
 * Dislikes: 0
 * Total Accepted:    125.8K
 * Total Submissions: 242.7K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */

package leetcode

// @lc code=start
func trap(height []int) int {
	// 暴力解法 O(n^2) O(1)
	//ans := 0
	//for i := 1; i < len(height)-1; i++ {
	//	maxLeft, maxRight := 0, 0
	//	for j := i; j >= 0; j-- {
	//		maxLeft = max(maxLeft, height[j])
	//	}
	//	for j := i; j < len(height); j++ {
	//		maxRight = max(maxRight, height[j])
	//	}
	//	ans += min(maxLeft, maxRight) - height[i]
	//}
	//
	//return ans

	// 动态编程，提前存储 O(n) o(n)
	hlen := len(height)
	if hlen < 3 {
		return 0
	}
	hLeft, hRight := make([]int, hlen), make([]int, hlen)

	hLeft[0], hRight[hlen-1] = 0, 0
	hLeft[1], hRight[hlen-2] = height[0], height[hlen-1]
	maxLeft, maxRight := height[0], height[hlen-1]

	for i := 2; i < hlen; i++ {
		j := hlen - i - 1
		maxLeft, maxRight = max(height[i-1], maxLeft), max(height[j+1], maxRight)
		hLeft[i], hRight[j] = maxLeft, maxRight
	}

	water := 0
	for i := 0; i < hlen; i++ {
		water += max(min(hLeft[i], hRight[i])-height[i], 0)
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
