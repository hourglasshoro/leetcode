package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{"abcabcbb"}, want: 3},
		{name: "test2", args: args{"bbbbb"}, want: 1},
		{name: "test3", args: args{"pwwkew"}, want: 3},
		{name: "test4", args: args{""}, want: 0},
		{name: "test5", args: args{" "}, want: 1},
		{name: "test6", args: args{"dvdf"}, want: 3},
		{name: "test7", args: args{"abba"}, want: 2},
	}
	for _, tt := range tests {
		if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
			t.Errorf("%q. lengthOfLongestSubstring() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
