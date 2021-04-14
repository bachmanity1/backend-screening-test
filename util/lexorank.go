package util

import (
	"strings"
	"unicode"
)

// code is taken from https://github.com/xissy/lexorank, with slight modifications
const (
	minChar = byte('a')
	maxChar = byte('z')
)

// Rank returns a new rank string between prev and next.
func Rank(prev, next string) string {
	prev = longify(prev, minChar, len(next)-len(prev))

	mid := make([]byte, 0)
	for i := 0; i < len(prev) || i < len(next); i++ {
		prevChar := getChar(prev, i, minChar)
		nextChar := getChar(next, i, maxChar)
		midChar := getMid(prevChar, nextChar)
		mid = append(mid, midChar)
		if midChar != prevChar {
			break
		}
	}
	rank := string(mid)

	if rank == prev {
		return rank + "m"
	}
	return rank
}

func getMid(prev, next byte) byte {
	return (prev + next) / 2
}

func getChar(s string, i int, defaultChar byte) byte {
	if i >= len(s) {
		return defaultChar
	}
	return s[i]
}

func longify(prev string, ch byte, size int) string {
	if size <= 0 {
		return prev
	}
	return prev + strings.Repeat(string(ch), size)
}

func ParseRank(rank string) (string, bool) {
	rank = strings.ToLower(rank)
	for _, r := range rank {
		if !unicode.IsLetter(r) {
			return "", false
		}
	}
	return rank, true
}
