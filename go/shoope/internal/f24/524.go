package f24

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		//需要和上次枚举的数据不同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//c 对应的指针指向数组的最右端
		third := n - 1
		target := -nums[first] //因为 a+b+c=0

		//枚举 b
		for second := first + 1; second < n; second++ {
			//需要和上次枚举的数据不同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//从第一个条件退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	right := lowestCommonAncestor(root.Right, p, q)
	left := lowestCommonAncestor(root.Left, p, q)
	if right != nil && left != nil {
		return root
	}
	if right == nil {
		return left
	}
	return right
}
