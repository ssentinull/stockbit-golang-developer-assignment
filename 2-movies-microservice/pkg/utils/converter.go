package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

func IntToString(i int) string {
	return fmt.Sprint(i)
}
