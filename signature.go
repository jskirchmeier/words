package words

import "fmt"

type runeField struct {
	letter rune
	bits   int
	mask   int
	pos    uint
}

var fields = map[rune]*runeField{
	'e': {'e', 3, 7, 0},
	's': {'s', 3, 7, 3},
	'i': {'i', 3, 7, 6},
	'o': {'o', 3, 7, 9},
	'a': {'a', 3, 7, 12},
	't': {'t', 3, 7, 15},
	'n': {'n', 3, 7, 18},
	'l': {'l', 3, 7, 21},
	'r': {'r', 3, 7, 24},
	'd': {'d', 3, 7, 27},
	'g': {'g', 3, 7, 30},
	'f': {'f', 2, 3, 33},
	'c': {'c', 2, 3, 35},
	'p': {'p', 2, 3, 37},
	'z': {'z', 2, 3, 39},
	'm': {'m', 2, 3, 41},
	'b': {'b', 2, 3, 43},
	'u': {'u', 2, 3, 45},
	'h': {'h', 2, 3, 47},
	'k': {'k', 2, 3, 49},
	'v': {'v', 2, 3, 51},
	'y': {'y', 2, 3, 53},
	'w': {'w', 2, 3, 55},
	'j': {'j', 2, 3, 57},
	'x': {'x', 2, 3, 59},
	'q': {'q', 2, 3, 61},
}

// signature indicates how many of each letter is in a word
// bit fields are used to hold the number of each letter, 11 letters (e,s,i,o,a,t,b,l,r,d,g) can have up to 7 instances
// the rest can have only 3 instances.  while not all words can be represented here it does cover all the common words
// that we are interested in
//
// usage: if you have a collection of letters (think boggle or scrabble tiles) and take the signature of that collection (with the above limitations)
// then you can go through your word list and compare their signature with the collection's.  If the word's signature is
// a subset of the collection then that word can be made out of the collection
//
func Signature(word string) uint64 {
	var sig uint64
	runeCount := make(map[rune]int)
	for _, r := range word {
		runeCount[r] = runeCount[r] + 1
	}

	//fmt.Println("word : ", word)
	for r, c := range runeCount {
		if c > 0 {
			rf := fields[r]
			if rf == nil {
				fmt.Printf("no rune field for %c in word %s\n", r, word)
				continue
			}
			if c > rf.mask {
				fmt.Printf("cannot calculate signature too large a count (%d max : %d) for %c int %s\n", c, rf.mask, rf.letter, word)
				c = rf.mask
			}
			sig |= uint64((c & rf.mask) << rf.pos)
			//fmt.Printf("rune  %c  count %02x   bits  %02x  ==> %10x   sig  %10x\n", rf.letter, c, c&rf.mask, uint64((c&rf.mask)<<rf.pos), sig)
		}
	}
	return sig
}

// for use in a stream
