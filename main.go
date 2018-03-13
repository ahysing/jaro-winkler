package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func jaroWinkler(s1 string, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	var lenS1 int = utf8.RuneCountInString(s1)
	var lenS2 int = utf8.RuneCountInString(s2)
	// ensure that s1 is shorter than or same length as s2
	if lenS1 > lenS2 {
		temp := s2
		s2 = s1
		s1 = temp
	}

	var s1r = []rune(s1)
	var s2r = []rune(s2)
	isCommonCharInS2 := make([]bool, lenS2)
	// (1) find the number of characters the two strings have in common.
	// note that matching characters can only be half the length of the
	// longer string apart.
	var c int = 0 // count of common characters
	var t int = 0 // count of transpositions

	var prevPos int = -1
	for ix := 0; ix < lenS1; ix++ {
		var ch rune = s1r[ix]
		var end int = lenS2
		if ix+(lenS2/2) < end {
			end = ix + (lenS2 / 2)
		}

		var start = 0
		if ix-(lenS2/2) > 0 {
			start = ix - (lenS2 / 2)
		}

		for ix2 := start; ix2 < end; ix2++ {
			if ch == s2r[ix2] && isCommonCharInS2[ix2] == false {
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
		var score float64 = float64(c)/float64(len(s1)) + float64(c)/float64(len(s2)) + (float64(c)-float64(t))/float64(c)/3.0
		// (2) common prefix modification
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
	fmt.Fprintf(os.Stderr, "usage: jarowinkler [inputword] [inputword2]\n")
	os.Exit(2)
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	} else {
		var previous string
		for i, word := range os.Args {
			if i > 1 {
				score := jaroWinkler(previous, word)
				fmt.Print(previous)
				fmt.Print(" ")
				fmt.Print(word)
				fmt.Print(" => ")
				fmt.Println(score)
			}

			previous = word
		}
	}
}
