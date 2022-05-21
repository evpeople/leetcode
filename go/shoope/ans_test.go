package shoope

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		link     list.List
		hashmap  map[int]list.Element
		capacity int
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "te1",
			fields: fields{
				link:     *list.New(),
				hashmap:  make(map[int]list.Element),
				capacity: 2,
			},
			args: args{
				key:   1,
				value: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := Constructor(2)
			lru.Put(1, 1)
			lru.Put(2, 2)
			t1 := lru.Get(1)
			lru.Put(3, 3)
			t2 := lru.Get(2)
			lru.Put(4, 4)
			t3 := lru.Get(1)
			lru.Get(3)
			lru.Get(4)
			fmt.Println(t1, t2, t3)
		})
	}
}

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "te1",
			args: args{headA: &ListNode{4, &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}},
				headB: &ListNode{5, &ListNode{6, &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}}}},
			want: &ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{s: ""}, want: 0},
		{name: "t", args: args{s: "pwwkew"}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTwoSum(t *testing.T) {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	if a[0] == 0 && a[1] == 1 {
		t.Log("True")
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"test1", args{head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, k: 2}, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{num1: "456", num2: "77"}, "533"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"te", args{head: &ListNode{1, &ListNode{2, nil}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
