package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("A man, a plan, a canal: Panama") {
		t.Fail()
	}
}
