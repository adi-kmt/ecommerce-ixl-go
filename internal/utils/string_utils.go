package utils

import "strings"

func StringConvertToListOfString(stringToBeConverted string) []string {
	return strings.Split(strings.TrimSpace(stringToBeConverted), ",")
}
