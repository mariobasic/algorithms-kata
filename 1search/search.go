package main

import "math"

// linearSearch O(N)
func linearSearch(haystack []int, needle int) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false

}

// binarySearchListRecursion O(log N)
func binarySearchListRecursion(haystack []int, needle int) bool {
	length := len(haystack)

	if length == 1 && haystack[0] != needle {
		return false
	}
	if haystack[length/2] == needle {
		return true
	} else if haystack[length/2] > needle {
		return binarySearchListRecursion(haystack[:length/2], needle)
	} else if haystack[length/2] <= needle {
		return binarySearchListRecursion(haystack[length/2:], needle)
	}

	return false
}

// binarySearchList O(log N)
func binarySearchList(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)
	for lo < hi {
		m := (lo + hi) / 2
		//m := lo + (hi-lo)/2
		v := haystack[m]
		if v == needle {
			return true
		} else if v > needle {
			// if midpoint value is bigger than the needle
			// reduce the maximum array size from top to midpoint
			hi = m
		} else {
			// if midpoint value is lower than the needle
			// reduce maximum array size from bottom to midpoint
			lo = m + 1
		}
	}
	return false
}

// test Given two crystal balls that will break if dropped from high enough distance,
// determine the exact spot in which it will break in the most optimised way
//
// initial jump point is sqrt(N), check for breaks and jump sqrt(N), we jump back sqrt(N) and iterate
// sqrt(N)
func twoCristalBalls(breaks []bool) (index int) {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		// is first ball breaking after this jump
		if breaks[i] {
			break
		}
	}

	// iterate over last sqrt(len(breaks)) and find search false
	for j := i - jmpAmount; j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	// another way to do it
	/*	i -= jmpAmount
		for j := 0; j <= jmpAmount && i < len(breaks); j, i = j+1, i+1 {
			if breaks[i] {
				return i
			}
		}*/

	return -1
}
