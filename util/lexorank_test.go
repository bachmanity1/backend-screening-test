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
			rank: "mmm",
		},
		{
			prev: "aaa",
			next: "",
			rank: "mmm",
		},
		{
			prev: "aab",
			next: "aac",
			rank: "aabm",
		},
	} {
		if rank := Rank(test.prev, test.next); rank != test.rank {
			t.Errorf("prev[%s], next[%s]  => got[%s], expected[%s]", test.prev, test.next, rank, test.rank)
		}
	}
}
