package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	var output, result int
	var data []int
	data, output = []int{7, 1, 5, 3, 6, 4}, 7
	result = MaxProfit(data)
	if result != output {
		t.Logf("For %v expected result is %v, but got %v", data, output, result)
		t.Error()
	}

	data, output = []int{1, 2, 3, 4, 5}, 4
	result = MaxProfit(data)
	if result != output {
		t.Logf("For %v expected result is %v, but got %v", data, output, result)
		t.Error()
	}

	data, output = []int{7, 6, 4, 3, 1}, 0
	result = MaxProfit(data)
	if result != output {
		t.Logf("For %v expected result is %v, but got %v", data, output, result)
		t.Error()
	}

}
