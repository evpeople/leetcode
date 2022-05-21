package shoope

import (
	"container/list"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	tmp1 := head
// 	tmp2 := head.Next
// 	tmp1.Next = nil
// 	for tmp2 != nil {
// 		t := tmp2.Next
// 		tmp2.Next = tmp1
// 		tmp1 = tmp2
// 		tmp2 = t
// 	}
// 	return tmp1
// }

type MyQueue struct {
	stack1 []int
	stack2 []int
}

// func Constructor() MyQueue {
// 	return MyQueue{
// 		stack1: make([]int, 0),
// 		stack2: make([]int, 0),
// 	}
// }
func (q *MyQueue) in2out() {
	for len(q.stack1) > 0 {
		q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
		q.stack1 = q.stack1[:len(q.stack1)-1]
	}
}

func (q *MyQueue) Push(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q *MyQueue) Pop() int {
	if len(q.stack2) == 0 {
		q.in2out()
	}
	x := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.stack2) == 0 {
		q.in2out()
	}
	return q.stack2[len(q.stack2)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}

/*
重点 O1 插入，O1，替换，O1查找
插入用的双向链表
查找用hashmap
也就是每次查找后，把这个节点放在ring的头部，

*/
type LRUCache struct {
	link     *list.List
	hashmap  map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		link:     list.New().Init(),
		hashmap:  make(map[int]*list.Element),
		capacity: capacity,
	}
}
func (lru *LRUCache) moveToHead(value list.Element) {
	lru.link.MoveToFront(&value)
}
func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.hashmap[key]; !ok {
		return -1
	} else {
		lru.moveToHead(*v)
		return v.Value.(int)
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if value2, ok := lru.hashmap[key]; ok {
		lru.link.Remove(value2)
		lru.link.PushFront(value2)
		lru.hashmap[key] = lru.link.Front()
	} else {
		lru.link.PushFront(value)
		lru.hashmap[key] = lru.link.Front()
		if lru.link.Len() > lru.capacity {
			t := lru.link.Back()
			delete(lru.hashmap, t.Value.(int))
			lru.link.Remove(t)
		}
	}

}

//a+c  b+c  a+c+b   b+c+a
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	if curA != nil && curB != nil {
		for curA != curB {

			if curA != nil {
				curA = curA.Next
			} else {
				curA = headB
			}
			if curB != nil {
				curB = curB.Next
			} else {
				curB = headA
			}
			if curA.Val == curB.Val {
				return curA
			}
		}
	}
	return nil
}
func lengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}
	n := len(s)
	l, r := -1, 0
	ans := 0
	for i := 0; i < n; i++ {
		if l > -1 {
			delete(m, s[l])
		}
		for r < n {
			if _, ok := m[s[r]]; !ok {
				m[s[r]] = struct{}{}
				r++
			} else {
				break
			}
		}
		ans = max(ans, r-l-1)
		l++
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// cur1 := list1
	// cur2 := list2
	// head := list1
	// for cur1 != nil && cur2 != nil {
	// 	if cur1.Val < cur2.Val {
	// 		head.Next = cur1
	// 		cur1 = cur1.Next
	// 		head = head.Next
	// 	}
	// 	if cur2.Val < cur1.Val {
	// 		head.Next = cur2
	// 		cur2 = cur2.Next
	// 		head = head.Next
	// 	}
	// }
	// if cur1 != nil {
	// 	head.Next = cur1
	// } else {
	// 	head.Next = cur2
	// }
	// return list1
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list2, list1.Next)
		return list1
	}
}
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for key, v := range nums {
		if a1, ok := m[v]; ok {
			return []int{a1, key}
		} else {
			m[target-v] = key
		}
	}
	return nil
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	cur := len(nums1) - 1
	m--
	n--
	//关键点为，是否都被填充了
	for n != -1 {
		if nums1[m] > nums2[n] {
			nums1[cur], nums1[m] = nums1[m], nums1[cur]
			m--
		} else if nums1[m] <= nums2[n] {
			nums1[cur] = nums2[n]
			n--
		}
		cur--
	}
}
func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head

}
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{Next: head}
	pre := node
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return node.Next
			}
		}
		nex := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return node.Next
}

// func sortArray(nums []int) []int {
// 	lens := len(nums) - 1
// 	for i := lens / 2; i >= 0; i-- {
// 		down(nums, i, lens)
// 	}
// 	for i := lens; i >= 1; i-- {
// 		nums[0], nums[1] = nums[j], nums[0]
// 		lens--
// 		down(nums, 0, lens)
// 	}
// 	return nums
// }
func down(nums []int, i, lens int) { //O(logn)大根堆，如果堆顶节点小于叶子，向下调整
	max := i
	if i<<1+1 <= lens && nums[i<<1+1] > nums[max] {
		max = i<<1 + 1
	}
	if i<<1+2 <= lens && nums[i<<1+2] > nums[max] {
		max = i<<1 + 2
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		down(nums, max, lens)
	}
}

func quickSort(q []int, l, r int) {
	if l > r {
		return
	}
	x := q[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for q[i] >= x {
			i++
		}
		for q[j] <= x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	ans := 0
// 	for k, v := range nums {
// 		dp[k] = max(dp[k-1]+v, v)
// 		ans = max(dp[k], ans)
// 	}
// 	return ans
// }
// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]+nums[i-1] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	return max
// }
func addStrings(num1 string, num2 string) string {
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	if t := l1 - l2; t >= 0 {
		for i := 0; i < t; i++ {
			num2 = "0" + num2
		}
		l2 = l1
	} else {
		for i := 0; i < -t; i++ {
			num1 = "0" + num1
		}
		l1 = l2
	}
	ans := ""
	add := false
	for l1 != -1 {
		tmp := num1[l1] - '0' + num2[l2] - '0'
		if add {
			tmp++
		}
		ans = strconv.Itoa(int(tmp%10)) + ans
		add = (tmp/10 == 1)
		l1--
		l2--
	}
	if add {
		return "1" + ans
	}
	return ans
}
func hasCycle(head *ListNode) bool {
	quick := head
	slow := &ListNode{Next: head}
	for quick != nil {
		if quick == slow {
			return true
		}
		slow = slow.Next
		if quick.Next == nil {
			return false
		} else {
			quick = quick.Next.Next
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	level := make([]int, 0)
	level = append(level, root.Val)
	ret = append(ret, level)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) != 0 {
		t := len(queue)
		le := make([]int, 0)
		for i := 0; i < t; i++ {
			tr := queue[i]
			if tr != nil {
				queue = append(queue, tr.Left)
				queue = append(queue, tr.Right)
				le = append(le, tr.Val)
			}
			queue = queue[2:]
		}
		ret = append(ret, le)
	}
	return ret
}
