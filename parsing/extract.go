package parsing

import "regexp"

func Extract(file []byte, reg *regexp.Regexp) string {
	submatches := reg.FindSubmatch(file)
	if submatches == nil {
		return ""
	}

	names := reg.SubexpNames()
	for i, name := range names {
		if name == "result" {
			return string(submatches[i])
		}
	}

	return ""
}

func ExtractLast(file []byte, reg *regexp.Regexp) string {
	submatches := reg.FindAllSubmatch(file, -1)
	if submatches == nil {
		return ""
	}

	names := reg.SubexpNames()
	for i, name := range names {
		if name == "result" {
			return string(submatches[len(submatches)-1][i])
		}
	}

	return ""
}
