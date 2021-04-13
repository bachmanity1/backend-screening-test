package util

import "strings"

// code is taken from https://github.com/xissy/lexorank, with slight modifications

const (
	MinRank = "aaaa"
	MaxRank = "zzzz"
	minChar = byte('a')
	maxChar = byte('z')
)

// Rank returns a new rank string between prev and next.
func Rank(prev, next string) string {
	if prev == "" {
		prev = strings.Repeat("a", len(next))
	}
	mid := make([]byte, 0)
	for i := 0; i < len(prev) || i < len(next); i++ {
		prevChar := getChar(prev, i, minChar)
		nextChar := getChar(next, i, maxChar)
		midChar := getMid(prevChar, nextChar)
		mid = append(mid, midChar)
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
