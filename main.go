package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// HasChildrenSum reports whether every non-leaf node's value equals the sum
// of its children's values (a missing child counts as 0). Leaves pass
// trivially.
func HasChildrenSum(n *TreeNode) bool {
	if n == nil || (n.Left == nil && n.Right == nil) {
		return true
	}
	sum := 0
	if n.Left != nil {
		sum += n.Left.Val
	}
	if n.Right != nil {
		sum += n.Right.Val
	}
	return n.Val == sum && HasChildrenSum(n.Left) && HasChildrenSum(n.Right)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	ok := tn(10, tn(8, tn(3, nil, nil), tn(5, nil, nil)), tn(2, tn(2, nil, nil), nil))
	fmt.Println(HasChildrenSum(ok))
}
