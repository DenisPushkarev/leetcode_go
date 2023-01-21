package leetcode

import "testing"

func TestReverseInt(t *testing.T) {
	var input, output int
	var result int
	input, output = 123, 321
	result = reverseInt(input)
	if result != output {
		t.Logf("For %v result expected to be %v, but got %v", input, output, result)
		t.Fail()
	}

	input, output = -123, -321
	result = reverseInt(input)
	if result != output {
		t.Logf("For %v result expected to be %v, but got %v", input, output, result)
		t.Fail()
	}

	input, output = -100, -1
	result = reverseInt(input)
	if result != output {
		t.Logf("For %v result expected to be %v, but got %v", input, output, result)
		t.Fail()
	}

	input, output = 1534236469, 0
	result = reverseInt(input)
	if result != output {
		t.Logf("For %v result expected to be %v, but got %v", input, output, result)
		t.Fail()
	}

	input, output = -1, -1
	result = reverseInt(input)
	if result != output {
		t.Logf("For %v result expected to be %v, but got %v", input, output, result)
		t.Fail()
	}
}
