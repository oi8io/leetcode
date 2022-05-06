/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (48.38%)
 * Likes:    1451
 * Dislikes: 0
 * Total Accepted:    431.4K
 * Total Submissions: 890.5K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

package problems

// @lc code=start
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	//将最后一个元素取出来，与之前的合并，如果合并成功，则删除最后一个元素
	for i := n - 1; i >= 0; {
		// fmt.Printf("%+v ", intervals[:n])
		if i < n-1 {
			intervals[i], intervals[n-1] = intervals[n-1], intervals[i]
		}
		x := intervals[n-1]
		// fmt.Println(n, x)
		m := n
		for j := 0; j < i; j++ {
			if n == 1 {
				break
			}
			if i == j {
				continue
			}
			y := intervals[j]
			//三种情况，
			//第一种 x 在 y 的左边/右边 不想交
			if x[1] < y[0] || x[0] > y[1] {
				continue
			}
			start, end := x[0], x[1]
			if x[0] > y[0] {
				start = y[0]
			}
			if x[1] < y[1] {
				end = y[1]
			}
			intervals[j][0] = start
			intervals[j][1] = end
			m = n - 1
			//fmt.Println(intervals[:n], "j:", i, j, n, x, y)
		}
		if n >= m {
			i--
		}
		n = m
	}
	//fmt.Println(intervals)
	return intervals[:n]
}

// @lc code=end
