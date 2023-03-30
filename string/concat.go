package string

import "strings"

func Contact(a, b string) string {
	return a + "/" + b
}

func Builder(a, b string) string {
	bb := strings.Builder{}
	bb.Grow(len(a) + len(b) + 1)

	bb.WriteString(a)
	bb.WriteString("/")
	bb.WriteString(b)

	return bb.String()
}
