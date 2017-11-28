package medium

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "(5) + (5) -> (0 -> 1)",
			args: args{l1: &ListNode{Val: 5}, l2: &ListNode{Val: 5}},
			want: &ListNode{Next: &ListNode{Val: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
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
		{"abcabcbb 's longest substring is abc", args{"abcabcbb"}, 3},
		{"bbbbb 's longest substring is b", args{"bbbbb"}, 1},
		{"pwwkew 's longest substring is wke", args{"pwwkew"}, 3},
		{"abcdefg 's longest substring is abcdefg", args{"abcdefg"}, 7},
		{"c 's longest substring is c", args{"c"}, 1},
		{"abba 's longest substring is ab", args{"abba"}, 2},
		{"dvdf 's longest substring is vdf", args{"dvdf"}, 3},
		{"anviaj 's longest substring is nviaj", args{"anviaj"}, 5},
		{"\"\" 's longest substring is \"\"", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
