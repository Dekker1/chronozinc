package parsing

import "regexp"

// Match tries regular expressions from the given map on the file; it returns
// the key of the first match
func Match(file []byte, dict map[string]*regexp.Regexp) string {
	for key, reg := range dict {
		if reg.Match(file) {
			return key
		}
	}
	return ""
}
