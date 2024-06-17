package utils

import (
	"strconv"
	"strings"
)

func StringConvertToListOfInt(stringToBeConverted string) []int32 {
	var intList []int32
	strings := strings.Split(strings.TrimSpace(stringToBeConverted), ",")
	for _, string := range strings {
		intVal, err := strconv.Atoi(string)
		if err != nil {
			continue
		}
		intList = append(intList, int32(intVal))
	}
	return intList
}
