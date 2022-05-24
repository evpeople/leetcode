package f24

import (
	"container/heap"
	"math"
	"sort"
)

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
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v { //已经存在的数据
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++ //这个字母是需要的，所以 ++
		}
		for check() && l <= r { //检查ori中的每个字母是否都有了足够的数据来覆盖
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]-- // -- 是因为，下一次循环要把指针再右移，这个数据就不重要了。
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
func buildHeap(nums []int) []int {
	heap.Init()
}

func findKthLargest(nums []int, k int) int {

}
