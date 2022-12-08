package utils

import (
	"regexp"
	"strings"
)

func StringToList(str string) []string {
	withoutWhitespace := strings.ReplaceAll(str, " ", "")
	commaSeparated := strings.Split(withoutWhitespace, ",")

	return commaSeparated
}

func IsIdListStringAllowed(str string) bool {
	var regex, _ = regexp.Compile(`^\d+(,\d+)*$`)

	var isMatch = regex.MatchString(str)
	
	return isMatch
}