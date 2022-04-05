package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	a := maxProfit([]int{1, 2, 4})
	if a != 3 {
		t.Error(a)
	}
}
func TestReverseList(t *testing.T) {
	a := reverseList(&ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	for i := 0; i < 4; i++ {
		if a.Val == 4-i {
			t.Log("true")
			a = a.Next
		} else {
			t.Error(a.Val)
		}
	}

}
