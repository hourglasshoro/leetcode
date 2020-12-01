package main

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{s: "42"}, 42},
		{"test2", args{s: "-42"}, -42},
		{"test3", args{s: "4193 with words"}, 4193},
		{"test4", args{s: "words and 987"}, 0},
		{"test5", args{s: "-91283472332"}, -2147483648},
		{"test6", args{s: "   -42"}, -42},
		{"test7", args{s: "+1"}, 1},
		{"test8", args{s: "+-12"}, 0},
		{"test8", args{s: "-+12"}, 0},
		{"test9", args{s: "00000-42a1234"}, 0},
		{"test10", args{s: "   +0 123"}, 0},
		{"test11", args{s: ""}, 0},
		{"test12", args{s: "20000000000000000000"}, 2147483647},
		{"test13", args{s: "  0000000000012345678"}, 12345678},
		{"test14", args{s: "-000000000000001"}, -1},
	}
	for _, tt := range tests {
		if got := myAtoi(tt.args.s); got != tt.want {
			t.Errorf("%q. myAtoi() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
