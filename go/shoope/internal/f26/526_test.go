package f26

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test2", args{[]*ListNode{
			&ListNode{2, nil},
			nil,
			&ListNode{-1, nil},
		}}, &ListNode{1, nil}},
		{"test2", args{[]*ListNode{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}},
		}}, &ListNode{1, nil}},
		{"test1", args{[]*ListNode{
			&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{2, &ListNode{6, nil}},
		}}, &ListNode{1, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
