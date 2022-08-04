package iteration

import "strings"

func Iterate(statement string, count int) string {
	return strings.Repeat(statement, count)
}
