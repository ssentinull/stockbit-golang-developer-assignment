package utils

import "strings"

func SplitStringByComma(str string) []string {
	strSpaceRemoved := strings.ReplaceAll(str, ", ", ",")
	strSplits := strings.Split(strSpaceRemoved, ",")

	return strSplits
}
