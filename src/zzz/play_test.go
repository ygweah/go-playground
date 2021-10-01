package zzz

import (
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	s := `ABBBAAA`

	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if r[i] == 'A' {
			r[i] = 'B'
		} else {
			r[i] = 'A'
		}
	}
	t.Log(r)

	t.Log(string(r))
}

func Test_getHitProbability(t *testing.T) {
	{
		var R int32 = 2
		var C int32 = 3
		G := [][]int32{
			{0, 0, 1},
			{1, 0, 1},
		}
		t.Log(getHitProbability(R, C, G))
	}
	{
		var R int32 = 2
		var C int32 = 2
		G := [][]int32{
			{1, 1},
			{1, 1},
		}
		t.Log(getHitProbability(R, C, G))
	}
}

func getHitProbability(R int32, C int32, G [][]int32) float64 {
	// Write your code here
	var count int32 = 0
	for _, r := range G {
		for _, c := range r {
			if c != 0 {
				count++
			}
		}
	}
	return float64(count) / float64(R*C)
}

func Test_getMaxAdditionalDinersCount(t *testing.T) {
	{
		var N int64 = 10
		var K int64 = 1
		var M int32 = 2
		S := []int64{2, 6}

		t.Log(getMaxAdditionalDinersCount(N, K, M, S))
	}
	{
		var N int64 = 15
		var K int64 = 2
		var M int32 = 3
		S := []int64{11, 6, 14}

		t.Log(getMaxAdditionalDinersCount(N, K, M, S))
	}
}

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {
	// Write your code here
	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })

	calculate := func(begin int64, end int64) int64 {
		size := end - begin + 1
		if begin > 1 && end < N {
			size -= K
		}
		count := size / (K + 1)
		if count < 0 {
			count = 0
		}
		return count
	}

	var count int64 = 0
	var begin int64 = 1
	for _, s := range S {
		count += calculate(begin, s-1)
		begin = s + 1
	}

	count += calculate(begin, N)
	return count
}
