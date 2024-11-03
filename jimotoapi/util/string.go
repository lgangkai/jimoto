package util

import "strconv"

func Str2Num(n string) int {
	if n == "" {
		return 0
	}
	num, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	return num
}
