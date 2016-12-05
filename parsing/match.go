package parsing

import (
	"fmt"
	"regexp"
)

func Match(file []byte, dict map[string]*regexp.Regexp) string {
	for key, reg := range dict {
		fmt.Printf("test %s", key)
		if reg.Match(file) {
			return key
		}
	}
	return ""
}
