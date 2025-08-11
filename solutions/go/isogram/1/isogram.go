package isogram

import "strings"

func IsIsogram(word string) bool {
	charactersMap := map[rune]int{}
	for _, r := range strings.ToLower(word) {
		charactersMap[r] += 1
		if r != ' ' && r != '-' {
			if charactersMap[r] > 1 {
				return false
			}
		}

	}
	return true
}
