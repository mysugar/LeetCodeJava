package n635


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//Example 1:
//Input:
//5
/// \
//3   6
/// \   \
//2   4   7
//
//Target = 9
//
//Output: True
//Example 2:
//Input:
//5
/// \
//3   6
/// \   \
//2   4   7
//
//Target = 28
//
//Output: False

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	     }
//func findTarget(root *TreeNode, k int) bool {
//	find := k-root.Val
//	found := false
//	if root.Left==nil {
//		return found
//	}
//	if root.Right==nil {
//		return found
//	}
//	if find<0 {//find left
//		found = goLefet(root.Left,find)
//
//	}else {//find right
//
//	}
//}