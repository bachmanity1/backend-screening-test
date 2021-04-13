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
			prev: "zzz",
			next: "",
			rank: "zzzm",
		},
		{
			prev: "",
			next: "ab",
			rank: "aam",
		},
	} {
		if rank := Rank(test.prev, test.next); rank != test.rank {
			t.Errorf("prev[%s], next[%s]  => got[%s], expected[%s]", test.prev, test.next, rank, test.rank)
		}
	}
}
