package formatting

import "strings"

func ToBold(str string) string {
	return "*" + str + "*"
}

func Normalize(str string) string {
	return strings.TrimSpace(strings.ToUpper(str))
}
