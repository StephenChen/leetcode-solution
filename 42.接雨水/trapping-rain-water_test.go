package leetcode

import "testing"

func Test_TrappingRainWater(t *testing.T) {

	tests := map[int][]int{
		6: {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}

	for out, in := range tests {
		rtn := out == trap(in)
		if !rtn {
			t.Errorf("[in]: %v    [out]: %v    [rtn]: %v", in, out, rtn)
		}
	}
}
