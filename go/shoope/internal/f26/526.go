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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	pL := root.Left
	pR := root.Right
	root.Left = invertTree(pR)
	root.Right = invertTree(pL)
	return root
}
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
func moveZeroes(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++ //左指针，
		}
		r++
	}
}

// func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil && l2 == nil {
// 		return nil
// 	}
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Next == nil && l2.Next != nil {
// 		t := addTwoNumbers2(l1, l2.Next)
// 		if t.Val > 10 {
// 			t.Val %= 10
// 			l2.Val++
// 		}
// 		l2.Next = t
// 		return l2
// 	}
// 	if l1.Next != nil && l2.Next == nil {
// 		t := addTwoNumbers2(l1.Next, l2)
// 		if t.Val > 10 {
// 			t.Val %= 10
// 			l1.Val++
// 		}
// 		l1.Next = t
// 		return l1
// 	}
// 	if l1.Next == nil && l2.Next == nil {
// 		return &ListNode{l1.Val + l2.Val, nil}
// 	}
// 	t := &ListNode{0, addTwoNumbers2(l1.Next, l2.Next)}
// 	return t.Next
// }
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表的两数之和，考虑栈的使用
	stack1 := []int{}
	stack2 := []int{}
	// 两个链表遍历的val加到栈中
	for l1 != nil || l2 != nil {
		if l1 != nil {
			stack1 = append(stack1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			stack2 = append(stack2, l2.Val)
			l2 = l2.Next
		}
	}
	carry := 0
	var head *ListNode
	var v1, v2 int
	for len(stack1) > 0 || len(stack2) > 0 {
		v1, v2 = 0, 0
		if len(stack1) > 0 {
			v1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			v2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		head = &ListNode{Val: sum, Next: head}

	}
	if carry > 0 {
		head = &ListNode{Val: carry, Next: head}
	}
	return head
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack) == 1 {
		this.minStack = append(this.minStack, val)
		return
	}
	if this.stack[len(this.stack)-1] < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
func MiStack() []int {
	t := Constructor()
	t.Push(-2)
	t.Push(0)
	t.Push(-3)
	tan := []int{}
	tan = append(tan, t.GetMin())
	t.Pop()
	tan = append(tan, t.Top())
	tan = append(tan, t.GetMin())
	return tan
}
