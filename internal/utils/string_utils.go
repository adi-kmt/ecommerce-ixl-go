package utils

import "strings"

func StringConvertToListOfInt(stringToBeConverted string) []string {
	return strings.Split(strings.TrimSpace(stringToBeConverted), ",")
}
