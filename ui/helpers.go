package ui

import (
	"fmt"
	"strconv"
)

func getUint8(s string) (val uint8, ok bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		ok = false
		return
	}
	if i < 0 || i > 127 {
		ok = false
		return
	}
	val = uint8(i)
	ok = true
	return
}

func printv(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
