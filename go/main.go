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
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hashMap[nums[i]]; ok {
			return []int{i, v}
		} else {
			hashMap[target-nums[i]] = i
		}
	}
	return nil
}
func firstUniqChar(s string) byte {
	var hMap map[byte]int
	tmp := []byte(s)
	hMap = map[byte]int{}
	for k, v := range tmp {
		if _, ok := hMap[v]; ok {
			hMap[v] = -1
			continue
		}
		hMap[v] = k
	}
	index := math.MaxInt32
	for _, v := range hMap {
		if v == -1 {
			continue
		}
		index = min(v, index)
	}
	if index == math.MaxInt32 {
		return ' '
	} else {
		return tmp[index]
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func isVaild(s string) bool {
	tmp := []byte(s)
	stack := make([]byte, 0)
	for _, v := range tmp {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			ans := true
			if len(stack) == 0 {
				return false
			}
			value := stack[len(stack)-1]

			switch value {
			case '(':
				ans = (v == ')')
			case '{':
				ans = (v == '}')
			case '[':
				ans = (v == ']')
			}
			if ans {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
func main() {

}
