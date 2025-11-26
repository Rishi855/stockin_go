package helper

import "strconv"

func StringIncrement(s string) string {
	n, _ := strconv.Atoi(s)
	n++
	return strconv.Itoa(n)
}
