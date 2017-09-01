package gstuff

import "fmt"

func FnOne(ss ...string) string {
	return fnOne(ss)
}

func fnOne(ss []string) string {
	return fmt.Sprintf("FnOne:%s", ss[0])
}
