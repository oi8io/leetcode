//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
//
//
// 示例 1：
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
//
// 示例 2：
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
//
// 👍 748 👎 0

package cn

import (
	//. "github.com/oi8io/lee/cn/449.serialize_and_deserialize_bst"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	answer := pathSum1(root, targetSum, make([]int, 0))
	return answer
}
func pathSum1(root *TreeNode, targetSum int, path []int) [][]int {
	if root == nil {
		return nil
	}
	nt := targetSum - root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if nt == 0 {
			return [][]int{path}
		} else {
			return nil
		}
	}
	var answer = make([][]int, 0)
	left := pathSum1(root.Left, nt, copyPath(path))
	answer = append(answer, left...)
	right := pathSum1(root.Right, nt, copyPath(path))
	answer = append(answer, right...)
	return answer
}

func copyPath(p []int) []int {
	np := make([]int, len(p))
	copy(np, p)
	return np
}

//leetcode submit region end(Prohibit modification and deletion)
