package f27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			t := nums[nums[i]]
			if t == nums[i] {
				return t
			}
			nums[nums[i]] = nums[i]
			nums[i] = t
			i--
		}
	}
	return 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
func postorderTraversal(root *TreeNode) []int {
	var tmp []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		tmp = append(tmp, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		tmp = append(tmp, postorderTraversal(root.Right)...)
	}
	tmp = append(tmp, root.Val)
	return tmp
}
func combine(n int, k int) (ans [][]int) {
	tmp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			comb := make([]int, k)
			copy(comb, tmp)
			ans = append(ans, comb)
			return
		}
		tmp = append(tmp, cur)
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	cur := head
// 	tail := head
// 	for i := 0; i < n+1; i++ {
// 		tail = tail.Next
// 	}
// 	for tail != nil {
// 		tail = tail.Next
// 		cur = cur.Next
// 	}
// 	cur.Next = cur.Next.Next
// 	return head
// }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
