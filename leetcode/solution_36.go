package leetcode

func IsValidSudoku(board [][]byte) bool {
	return isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	var c byte
	for sc := 0; sc < 3; sc++ {
		for sr := 0; sr < 3; sr++ {
			var s = map[byte]bool{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					c = board[3*sc+i][3*sr+j]
					if c == '.' {
						continue
					}
					if s[c] == true {
						return false
					} else {
						s[c] = true
					}
				}
			}
		}
	}
	for i := 0; i < 9; i++ {
		var s = map[byte]bool{}
		var s2 = map[byte]bool{}
		for j := 0; j < 9; j++ {
			c = board[i][j]
			if c != '.' {

				if s[c] == true {
					return false
				} else {
					s[c] = true
				}
			}

			c = board[j][i]
			if c != '.' {
				if s2[c] == true {
					return false
				} else {
					s2[c] = true
				}
			}
		}
	}
	return true
}
