package util

import "testing"

func TestRank(t *testing.T) {
	for _, test := range []struct {
		prev string
		next string
		rank string
	}{
		{
			prev: "aaa",
			next: "zzz",
			rank: "m",
		},
		{
			prev: "aaa",
			next: "",
			rank: "m",
		},
		{
			prev: "aab",
			next: "aac",
			rank: "aabm",
		},
		{
			prev: "",
			next: "",
			rank: "m",
		},
		{
			prev: "zzy",
			next: "",
			rank: "zzym",
		},
		{
			prev: "",
			next: "ab",
			rank: "aam",
		},
		{
			prev: "a",
			next: "ab",
			rank: "aam",
		},
		{
			prev: "ab",
			next: "b",
			rank: "an",
		},
	} {
		if rank := Rank(test.prev, test.next); rank != test.rank {
			t.Errorf("prev[%s], next[%s]  => got[%s], expected[%s]", test.prev, test.next, rank, test.rank)
		}
	}
}
