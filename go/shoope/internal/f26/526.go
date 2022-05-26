package f26

import "container/heap"

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
