package binarytree

func MaxPathSum(root *TreeNode) int {
	// x := sum(root)
	return 0
}

type PathSumInfo struct {
	Max          int
	Sum          int
	LeftPathSum  int
	RightPathSum int
}

// func sum(root *TreeNode) *PathSumInfo {
// 	if root == nil {
// 		return &PathSumInfo{Max: math.MinInt}
// 	}

// 	left := sum(root.Left)
// 	right := sum(root.Right)
// 	var ans int

// 		ans = max(left.LeftPathSum+right.RightPathSum+root.Val, max1(left.LeftPathSum, right.RightPathSum)+root.Val, root.Val)
// 		ans = max2(left.Max, right.Max, ans)
// 		return &PathSumInfo{Max: ans, LeftPathSum: max1(left.LeftPathSum, right.RightPathSum) + root.Val, RightPathSum: max1(left.LeftPathSum, right.RightPathSum) + root.Val}
// 	}

// }
// func max2(a, b, c int) int {
// 	if (a > b && a > c) || (a >= b && a > c) || (a > b && a >= c) {
// 		return a
// 	} else if (b > a && b > c) || (b >= a && b > c) || (b > a && b >= c) {
// 		return b
// 	} else {
// 		return c
// 	}
// }
