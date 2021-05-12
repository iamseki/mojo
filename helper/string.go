package helper

import (
	"strings"
	"unicode"
)

// FormatStringFromAWS recieves a string like:
//
// PRE ffa2f82c-3de9-4a61-b7b4-dd471e62fda6/\n
// PRE baa2f82c-3de9-4a61-b7b4-dd471e62fs54/\n
// PRE cfa2f82c-3de9-4a61-a7b4-dd471e62fda6/\n
//
// returning str separate by comma:
// ffa2f82c-3de9-4a61-b7b4-dd471e62fda6,baa2f82c-3de9-4a61-b7b4-dd471e62fs54 ...
func SliceFromAWSUnformattedString(str string) []string {
	// implemented as a pipeline so in each step unnecessary character is removed
	str = strings.ReplaceAll(str, "\n", ",")
	str = strings.ReplaceAll(str, "PRE", "")
	str = strings.ReplaceAll(str, "/", "")
	// remove blank spaces
	var out strings.Builder
	out.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			out.WriteRune(ch)
		}
	}
	str = out.String()

	if strings.Index(str, ",") == 0 {
		// removing comma in first position
		str = str[1:]
	}
	if strings.LastIndex(str, ",") == len(str)-1 {
		last := len(str) - 1
		// removing comma in last position
		str = str[:last]
	}

	return strings.Split(str, ",")
}
