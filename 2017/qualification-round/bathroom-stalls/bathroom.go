package main

import (
	"container/heap"
)

type bathroomSpaces struct {
	available    *maxHeap
	multiplicity map[uint64]uint64
}

// newBathroom returns a new bathroom of the given inital space
func newBathroom(space uint64) *bathroomSpaces {
	b := &bathroomSpaces{
		available: &maxHeap{space},
		multiplicity: map[uint64]uint64{
			space: 1,
		},
	}
	heap.Init(b.available)

	return b
}

// popSpace removes and returns the largest available space in the bathroom and
// the number of copies
func (b *bathroomSpaces) popSpace() (space uint64, mult uint64) {
	space = heap.Pop(b.available).(uint64)
	mult = b.multiplicity[space]
	delete(b.multiplicity, space)

	return space, mult
}

// addSpace adds a space of the given size to the bathroom, mult times
func (b *bathroomSpaces) addSpace(space, mult uint64) {
	b.multiplicity[space] += mult

	if b.multiplicity[space] == mult {
		heap.Push(b.available, space)
	}
}

// addPeople adds people to bathroom stalls and returns the min and max
// distances to the nearest people to the left and to the right and the amount
// of people that can be added with the same min and max results
func (b *bathroomSpaces) addPeople() (min uint64, max uint64, count uint64) {
	space, count := b.popSpace()

	min, max = divideSpace(space)
	b.addSpace(min, count)
	b.addSpace(max, count)

	return min, max, count
}

// divideSpace divides the given space along the middle, occupying the space
// closest to the middle, and returns the new min and max spaces
func divideSpace(space uint64) (min uint64, max uint64) {
	if space%2 == 0 {
		return (space / 2) - 1, (space / 2)
	}

	return (space - 1) / 2, (space - 1) / 2
}

// lastPersonDistances creates a new bathroom with the given number of stalls
// and places numPeople people in these stalls, according to the deterministic
// distribution in the exercise, and returns the min and max distances from
// the last person places to the two people nearest to it
func lastPersonDistances(numStalls, numPeople uint64) (min uint64, max uint64) {
	b := newBathroom(numStalls)

	var placed uint64
	for numPeople > placed {
		var count uint64
		min, max, count = b.addPeople()
		placed += count
	}

	return min, max
}
