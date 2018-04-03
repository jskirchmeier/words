package words

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestConstruct(t *testing.T) {
//	tests := []struct {
//		name      string
//		signature uint64
//		ratting   int
//	}{
//		{"one", 1, 1},
//		{"two", 1, 1},
//		{"three", 1, 1},
//		{"four", 1, 1},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			w := Word{test.name, test.signature, test.ratting}
//			assert.Equal(t, test.name, w.Word())
//			assert.Equal(t, test.signature, w.signature)
//			assert.Equal(t, test.ratting, w.Rating())
//		})
//	}
//}
//
//func TestRuneCount(t *testing.T) {
//
//	type counters struct {
//		r rune
//		c int
//	}
//	tests := []struct {
//		name   string
//		runes  int
//		counts []counters
//	}{
//		{"abc", 3, []counters{{'a', 1}, {'b', 1}, {'c', 1}}},
//		{"abcabcabc", 3, []counters{{'a', 3}, {'b', 3}, {'c', 3}}},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			w := Word{test.name, 0, 0}
//			counts := w.RuneCount()
//			assert.Equal(t, test.runes, len(counts))
//			for _, c := range test.counts {
//				assert.Equal(t, c.c, counts[c.r])
//			}
//
//		})
//	}
//}
//
//func TestSignature(t *testing.T) {
//	//assert.Equal(t, uint64(1), Signature("e"))
//	//assert.Equal(t, uint64(2), Signature("ee"))
//	//assert.Equal(t, uint64(3), Signature("eee"))
//	//assert.Equal(t, uint64(4), Signature("eeee"))
//	//assert.Equal(t, uint64(5), Signature("eeeee"))
//	//
//	//assert.Equal(t, uint64(9), Signature("es"))
//	//assert.Equal(t, uint64(9), Signature("se"))
//	//assert.Equal(t, uint64(18), Signature("eess"))
//	//assert.Equal(t, uint64(18), Signature("eses"))
//	//
//	//fmt.Printf("%-10s = %20x\n", "max size", math.MaxInt64)
//	//
//	//for _, w := range []string{"john", "zoo", "capitulate", "assert", "quiunxxhh"} {
//	//	fmt.Printf("%-10s = %20x\n", w, Signature(w))
//	//
//	//}
//
//	//Signature("qqqqqqqqq")
//
//}

func TestGenerateRuneMap(t *testing.T) {
	assert.True(t, true)
	// generate map creating code
	runes := [...]rune{'e', 's', 'i', 'o', 'a', 't', 'n', 'l', 'r', 'd', 'g', 'f', 'c', 'p', 'z', 'm', 'b', 'u', 'h', 'k', 'v', 'y', 'w', 'j', 'x', 'q'}
	pos := 0
	size := 3
	mask := 7
	count := 0
	// we only have enough bits for 11 words to have 3 bits, the rest have only 2, the runes are in frequency order
	for idx, r := range runes {
		if idx >= 11 {
			size = 2
			mask = 3
		}
		//type runeField struct {
		//	letter rune
		//	bits   int
		//	mask   int
		//	pos    uint
		//}

		count += size
		fmt.Printf("'%c' : {'%c', %d, %d, %d},\n", r, r, size, mask, pos)
		pos += size
	}
	fmt.Println(count, " bits used")

}
