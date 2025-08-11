package pangram

import "strings"

func IsPangram(input string) bool {
	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(strings.ToLower(input), i) {
			return false

		}
	}
	return true
}
