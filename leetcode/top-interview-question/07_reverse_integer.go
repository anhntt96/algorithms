package leetcode

import "strconv"

//https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	max_num := "2147483647"
	min_num := "2147483648"
	negative_number := false
	s := strconv.Itoa(x)
	if s[0] == '-' {
		s = s[1:]
		negative_number = true
	}

	sRev := ""
	for _, v := range s {
		sRev = string(v) + sRev
	}
	for len(sRev) > 1 && sRev[0] == '0' {
		sRev = sRev[1:]
	}

	if negative_number {
		if len(sRev) < len(max_num) || sRev < min_num {
			sRev = "-" + sRev
			rs, err := strconv.Atoi(sRev)
			if err == nil {
				return rs
			}
		}
	} else {
		if len(sRev) < len(max_num) || sRev < max_num {
			rs, err := strconv.Atoi(sRev)
			if err == nil {
				return rs
			}
		}

	}
	return 0
}
