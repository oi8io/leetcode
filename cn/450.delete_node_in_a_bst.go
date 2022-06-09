//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
//根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
//
//
//
//
// 示例 1:
//
//
//
//
//输入：root = [5,3,6,2,4,null,7], key = 3
//输出：[5,4,6,2,null,null,7]
//解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
//一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
//另一个正确答案是 [5,2,6,null,4,null,7]。
//
//
//
//
// 示例 2:
//
//
//输入: root = [5,3,6,2,4,null,7], key = 0
//输出: [5,3,6,2,4,null,7]
//解释: 二叉树不包含值为 0 的节点
//
//
// 示例 3:
//
//
//输入: root = [], key = 0
//输出: []
//
//
//
// 提示:
//
//
// 节点数的范围 [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// 节点值唯一
// root 是合法的二叉搜索树
// -10⁵ <= key <= 10⁵
//
//
//
//
// 进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
// 👍 766 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	root = deleteNode1(root, key)
	return root
}

func FindBSTMin(root *TreeNode) *TreeNode {
	for root != nil && root.Left != nil {
		root = root.Left
	}
	return root
}

func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val && root.Right != nil {
		root.Right = deleteNode1(root.Right, key)
		return root
	} else if key < root.Val && root.Left != nil {
		root.Left = deleteNode1(root.Left, key)
		return root
	} else if key == root.Val {
		if root.Right == nil { // 右子树为空 0/1个子树
			return root.Left
		}
		if root.Left == nil { // 左子树为空 0/1个子树
			return root.Right
		}
		//两个孩子的情况，找右子节点
		min := FindBSTMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode1(root.Right, min.Val)
		return root
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
