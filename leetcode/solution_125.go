package leetcode

func IsPalindrome(s string) bool {
	return isPalindrome(s)
}
func isPalindrome(s string) bool {
	var l, r int16 = 0, int16(len(s) - 1)
	var cr, cl byte
	for l < r {
		cl, cr = mayBeToLower(s[l]), mayBeToLower(s[r])
		if !isValidAlphanumeric(cl) {
			l++
		} else if !isValidAlphanumeric(cr) {
			r--
		} else {
			if cl != cr {
				return false
			}
			l++
			r--
		}
	}
	return true
}
func mayBeToLower(b byte) byte {
	if 64 < b && b < 91 {
		return b + 32
	} else {
		return b
	}
}

func isValidAlphanumeric(b byte) bool {
	return ((b >= '0' && b <= '9') || (b >= 'a' && b <= 'z'))
}
