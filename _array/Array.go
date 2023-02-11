package _array

import (
	"strconv"
)

func InArray[T int | string](value T, array []T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func StringToInt(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
