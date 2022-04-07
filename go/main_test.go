package main

import (
	"reflect"
	"sort"
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
func TestTwoSum(t *testing.T) {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	if a[0] == 0 && a[1] == 1 {
		t.Log("True")
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// {"test1", args{"abaccdeff"}, 'b'},
		{"test2", args{"cc"}, ' '},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVaild(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"()"}, true},
		{"test2", args{"(]"}, false},
		{"test3", args{"([)]"}, false},
		{"test4", args{"("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVaild(tt.args.s); got != tt.want {
				t.Errorf("isVaild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test2", args{[]int{2, 0}, 1, []int{1}, 1}},
		{"test1", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !sort.IsSorted(sort.IntSlice(tt.args.nums1)) {
				t.Error(tt.args.nums1)
			}
		})
	}
}
