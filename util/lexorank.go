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
	next = longify(next, maxChar, len(prev)-len(next))

	mid := make([]byte, 0)
	for i := 0; i < len(prev); i++ {
		midChar := getMid(prev[i], next[i])
		mid = append(mid, midChar)
		if midChar != prev[i] {
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
