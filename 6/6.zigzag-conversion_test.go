package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{s: "PAYPALISHIRING", numRows: 3}, "PAHNAPLSIIGYIR"},
		{"test2", args{s: "PAYPALISHIRING", numRows: 4}, "PINALSIGYAHRPI"},
		{"test3", args{s: "A", numRows: 1}, "A"},
	}
	for _, tt := range tests {
		if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
			t.Errorf("%q. convert() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
