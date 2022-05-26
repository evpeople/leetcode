package f26

import (
	"container/heap"
	"sort"
)

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(Lremain int, rRemain int, path string)
	dfs = func(Lremain, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
		}
		if Lremain > 0 {
			dfs(Lremain-1, rRemain, path+"(")
		}
		if Lremain < rRemain {
			dfs(Lremain, rRemain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if lists == nil || lists[0] == nil {
// 		return nil
// 	}
// 	head := &ListNode{0, lists[0]}
// 	cur := head.Next
// 	for i := 1; i < len(lists); i++ {
// 		sec := lists[i]
// 		for sec != nil {
// 			tmp := sec.Next
// 			if cur.Next != nil {
// 				if cur.Val <= sec.Val && cur.Next.Val > sec.Val {
// 					sec.Next = cur.Next
// 					cur.Next = sec
// 					cur = cur.Next
// 				} else {
// 					cur = cur.Next
// 					tmp = sec
// 				}
// 			} else {
// 				cur.Next = sec
// 				break
// 			}
// 			sec = tmp
// 		}
// 		cur = head
// 	}
// 	return head.Next
// }

type ListNodes []*ListNode

func (l ListNodes) Len() int {
	return len(l)
}

func (l ListNodes) Less(i, j int) bool {
	return l[i].Val <= l[j].Val
}

func (l ListNodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	hp := &ListNodes{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(hp, lists[i])
		}
	}
	for len(*hp) > 0 {
		top := heap.Pop(hp).(*ListNode)
		pre.Next = top
		pre = pre.Next
		if top.Next != nil {
			heap.Push(hp, top.Next)
		}
	}
	return dummy.Next
}

func climbStairs(n int) int {
	// dp := make([]int, 0, 3)

	dp := [3]int{0, 1, 2}
	for i := 3; i < n+1; i++ {
		dp[i%3] = dp[(i+1)%3] + dp[(i+2)%3]
	}
	return dp[n%3]
}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := &ListNode{0, head}
	cur := ans
	for cur.Next != nil && cur.Next.Next != nil {
		if tmp := cur.Next.Val; tmp == cur.Next.Next.Val {
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return ans.Next
}

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	var mod int
	mod = 1e9 + 7
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}
func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}
