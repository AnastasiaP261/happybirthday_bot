package formatting

import "strings"

func ToBold(str string) string {
	return "*" + str + "*"
}

func Format(str string) string {
	return strings.TrimSpace(strings.ToUpper(str))
}
