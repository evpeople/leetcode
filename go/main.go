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

/*
void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
    int i = nums1.size() - 1;
    m--;
    n--;
    while (n >= 0) {
        while (m >= 0 && nums1[m] > nums2[n]) {
            swap(nums1[i--], nums1[m--]);
        }
        swap(nums1[i--], nums2[n--]);
    }
}
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1) - 1
	// i := m - 1
	// j := n - 1
	m--
	n--
	for n != -1 {
		for m != -1 && nums1[m] > nums2[n] {
			nums1[index], nums1[m] = nums1[m], nums1[index]
			index--
			m--
		}
		nums1[index] = nums2[n]
		index--
		n--
	}
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headBP := headB
	headAP := headA
	if headA != nil && headB != nil {
		for headA != headB {
			if headA == nil {
				headA = headBP
			} else {
				headA = headA.Next
			}
			if headB == nil {
				headB = headAP
			} else {
				headB = headB.Next
			}
		}
		return headA
	} else {
		return nil
	}
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	i1 := &ListNode{}
	tmp := i1
	for list2 != nil && list1 != nil {
		if list1.Val > list2.Val {
			i1.Next = list2
			i1 = i1.Next
			list2 = list2.Next
		} else {
			i1.Next = list1
			i1 = i1.Next
			list1 = list1.Next
		}
	}
	if list1 != nil {
		i1.Next = list1
	}
	if list2 != nil {
		i1.Next = list2
	}
	return tmp.Next
}
func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}
func hasCycle(head *ListNode) bool {
	turtle := head
	if turtle == nil {
		return false
	}
	rabbit := head.Next
	for rabbit != turtle && rabbit != nil && turtle != nil {
		if rabbit.Next != nil {
			rabbit = rabbit.Next.Next
		} else {
			rabbit = rabbit.Next
		}
		turtle = turtle.Next
	}
	if rabbit == nil {
		return false
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	if root.Left != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}
func deleteDuplicates(head *ListNode) *ListNode {
	slow := head
	quick := head
	for quick != nil {
		for {
			if quick.Next != nil && quick.Val == quick.Next.Val {
				quick = quick.Next
			} else {
				break
			}
		}
		slow.Next = quick.Next
		slow = slow.Next
		quick = quick.Next
	}
	return head
}
func main() {

}
