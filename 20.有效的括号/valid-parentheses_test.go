package leetcode

import "testing"

func Test_ValidParentheses(t *testing.T) {

	tests := map[string]bool{
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
	}

	for in, out := range tests {
		rtn := out == isValid(in)
		if !rtn {
			t.Errorf("[in]: %v    [out]: %v    [rtn]: %v", in, out, rtn)
		}
	}
}
