package leetcode

import "testing"

func TestAtoi(t *testing.T) {
	var input string = "-1"
	var output, expected = myAtoi(input), -1
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "-+12"
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "a-0"
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "   "
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "   -42"
	output, expected = myAtoi(input), -42
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

	input = "-100000000000000"
	output, expected = myAtoi(input), -2147483648
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

	input = "21474836460"
	output, expected = myAtoi(input), 2147483647
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

	input = "4193 with words"
	output, expected = myAtoi(input), 4193
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "words and 987"
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}
	input = "-91283472332"
	output, expected = myAtoi(input), -2147483648
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

	input = "+-12"
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

	input = "00000-42a1234"
	output, expected = myAtoi(input), 0
	if output != expected {
		t.Logf("Atoi(%v) expect to be %d, but got %d\n", input, expected, output)
		t.Fail()
	}

}
