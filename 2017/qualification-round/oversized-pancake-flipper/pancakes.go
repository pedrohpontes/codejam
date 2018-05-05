package main

const (
	happyPancake = '+'
	blankPancake = '-'
	panicStr     = "pancakes should be happy or blank"
)

type pancake rune

func (p *pancake) flip() {
	if *p == happyPancake {
		*p = blankPancake
	} else if *p == blankPancake {
		*p = happyPancake
	} else {
		panic(panicStr)
	}
}

// flipFirstPancakes flips the first numToFlip pancakes in ps
func flipFirstPancakes(ps []pancake, numToFlip int) {
	for i := 0; i < numToFlip; i++ {
		ps[i].flip()
	}
}

// minPancakeFlips returns the minimum number of flips required to flip all
// pancakes ps to happy.
// Returns -1 if this is imposssible.
func minPancakeFlips(ps []pancake, flipperSize int) int {
	if len(ps) == 0 {
		return 0
	}
	if ps[0] == happyPancake {
		return minPancakeFlips(ps[1:], flipperSize)
	} else if ps[0] == blankPancake {
		if len(ps) < flipperSize {
			return -1
		}

		flipFirstPancakes(ps, flipperSize)
		minFlipsOrImpossible := minPancakeFlips(ps[1:], flipperSize)
		if minFlipsOrImpossible == -1 {
			return -1
		}

		return 1 + minFlipsOrImpossible
	}

	panic(panicStr)
}
