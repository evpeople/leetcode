package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp1 := head
	tmp2 := head.Next
	tmp1.Next = nil
	for tmp2 != nil {
		t := tmp2.Next
		tmp2.Next = tmp1
		tmp1 = tmp2
		tmp2 = t
	}
	return tmp1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	dp := make([]int, 2)
	//dp[i]=max(dp[i-1],prices[i]-min)
	min := math.MaxInt32
	for k, v := range prices {
		if v < min {
			min = v
		}
		if k != 0 {
			dp[k%2] = max(dp[(k-1)%2], v-min)
		} else {
			dp[k] = 0
		}
	}
	return max(dp[0], dp[1])
}
func main() {

}
