package f24

import (
	"math"
	"sort"
	"strconv"
	"strings"
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
func buildHeap(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- { //从第一个非叶子节点开始处理
		down(nums, i, n)
	}
}
func down(nums []int, i0, n int) {
	i := i0
	for {
		j1 := 2*i + 1 //对应的叶子节点
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		//j2 是左孩子
		if j2 := j1 + 1; j2 < n && nums[j2] > nums[j1] {
			j = j2 //构建大根堆，找到更大的那个。
		}
		if nums[i] > nums[j] {
			break //父节点大于子节点
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j //下次判断是否要继续下降
	}
}
func up(nums []int, j int) {
	for {
		i := (j - 1) / 2
		if i == j || nums[j] > nums[i] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		j = i
	}
}

func pop(nums *[]int) {
	n := len(*nums) - 1
	(*nums)[0], (*nums)[n] = (*nums)[n], (*nums)[0]
	down(*nums, 0, n)
	*nums = (*nums)[:n]

}

// func pop(nums []int) {
// 	n := len(nums) - 1
// 	(nums)[0], (nums)[n] = (nums)[n], (nums)[0]
// 	down(nums, 0, n)
// 	&nums = &nums[2:]

// }
func findKthLargest(nums []int, k int) int {
	buildHeap(nums)
	for i := 0; i < k-1; i++ {
		pop(&nums)
	}
	return nums[0]
}
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	ans = append(ans, root.Val)
	if root.Left != nil {
		ans = append(ans, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		ans = append(ans, preorderTraversal(root.Right)...)
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		head = head.Next
	}
	return head
}
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	begien := 0
	maxLen := 1
	for l := 2; l < n; l++ {
		for i := 0; i < n; i++ {
			j := l + i - 1
			if j > n {
				break
			}
			dp[i][j] = dp[i+1][j-1] && s[j] == s[i]
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begien = i
			}
		}

	}
	return s[begien : begien+maxLen]
}
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
