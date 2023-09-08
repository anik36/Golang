package regex

import (
	"fmt"
	"regexp"
)

// This fuction will remove the sysmbols from given string and
// return the non-symbolic string
func RemoveSymbols(str string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		fmt.Println(err)
	}
	str = re.ReplaceAllString(str, "")
	return str
}
