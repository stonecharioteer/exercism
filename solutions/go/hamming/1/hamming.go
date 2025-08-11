package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, errors.New("The length of the 2 strands should be the same!")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance += 1
		}

	}
	return hammingDistance, nil
}
