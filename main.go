package main

import (
	"fmt"
	"os"
)

func jaroWinkler(s1 string, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	if len(s1) > len(s2) {
		var temp string
		temp = s2
		s2 = s1
		s1 = temp
	}

	var s1r = []rune(s1)
	var s2r = []rune(s2)
	isCommonCharInS2 := make([]bool, len(s2))
	// (1) find the number of characters the two strings have in common.
	// note that matching characters can only be half the length of the
	// longer string apart.
	var c int = 0 // count of common characters
	var t int = 0 // count of transpositions
	var halfLength int = len(s2) / 2
	var prevPos int = -1
	for i := 0; i < halfLength; i++ {
		var end int = len(s2)
		if i+halfLength < end {
			end = i + halfLength
		}

		var start = 0
		if i-halfLength > 0 {
			start = i - halfLength
		}

		for ix2 := start; ix2 < end; ix2++ {
			if s1r[i] == s2r[ix2] && isCommonCharInS2[ix2] == false {
				c++
				isCommonCharInS2[ix2] = true
				if prevPos != -1 && ix2 < prevPos {
					t++
				}

				prevPos = ix2
				break
			}
		}
	}

	if c == 0 {
		return 0.0
	} else {
		// (2) common prefix modification
		var score float64 = float64(c)/float64(len(s1)) + float64(c)/float64(len(s2)) + (float64(c)-float64(t))/float64(c)/3.0

		var last int
		if len(s1) < 4 {
			last = len(s1)
		} else {
			last = 4
		}

		p := 0
		for i := 0; i < last; i++ {
			if s1r[i] == s2r[i] {
				p++
			} else {
				break
			}
		}

		score = score + ((float64(p) * (1.0 - float64(p))) / 10.0)
		return score
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: jarowinkler [inputword] [inputword]\n")
	os.Exit(2)
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	} else {
		var other string
		for i, word := range os.Args {
			if i > 0 {
				score := jaroWinkler(word, other)
				fmt.Print(word)
				fmt.Print(" ")
				fmt.Print(other)
				fmt.Print(" => ")
				fmt.Println(score)
			}

			other = word
		}
	}
}
