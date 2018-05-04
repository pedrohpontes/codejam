package main

type number []uint8 // a number by its digits

func (n number) String() string {
	start := 0
	if n[0] == 0 {
		start = 1 // skip possible leading zero
	}

	r := make([]rune, len(n)-start)
	for i := start; i < len(n); i++ {
		r[i-start] = rune(n[i] + '0')
	}

	return string(r)
}

// lastTidy returns the last tidy number <= n
func lastTidy(n number) number {
	tidy := make([]uint8, len(n))
	copy(tidy, n)

	// loop through the digits of n
	// if we find that the next digit is smaller than the current, we need to
	// tidy up, starting from the first digit equal to the current digit
	firstIndexEqualToCurrent := 0
	for i := 0; i < len(n)-1; i++ {
		if n[i] < n[i+1] {
			firstIndexEqualToCurrent = i + 1
		} else if n[i] > n[i+1] {
			return tidyUp(tidy, firstIndexEqualToCurrent)
		}
	}

	return tidy
}

// tidyUp makes n tidy by "borrowing" a digit at the index and changing the
// following digits to 9
func tidyUp(n number, index int) number {
	n[index]--

	for i := index + 1; i < len(n); i++ {
		n[i] = 9
	}

	return n
}
