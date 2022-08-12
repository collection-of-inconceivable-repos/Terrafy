package util

import "unicode"

func IsBlank(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}

	return true
}
