package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	println("plain: ", plain, "shiftKey: ", shiftKey)
	result := ""
	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			shifted := ((r - 'a' + rune(shiftKey)) % 26) + 'a'
			result = result + string(shifted)

		} else if r >= 'A' && r <= 'Z' {
			shifted := ((r - 'A' + rune(shiftKey)) % 26) + 'A'
			result = result + string(shifted)
		} else {
			result = result + string(r)
		}
	}

	return result
}
