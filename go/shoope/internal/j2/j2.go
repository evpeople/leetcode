package j2

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	slow := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
