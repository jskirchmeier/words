package words

import (
	"fmt"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
)

type Word struct {
	word      string
	signature uint64
}

func (w *Word) Word() string {
	if w == nil {
		return ""
	}
	return w.word
}

func (w *Word) ToString() string {
	if w == nil {
		return ""
	}
	return fmt.Sprintf("%-15s  %15x", w.word, w.signature)
}

func (w *Word) Signature() uint64 {
	if w == nil {
		return 0
	}
	if w.signature == 0 {
		w.signature = Signature(w.word)
	}
	return w.signature
}

func NewWord(w string) *Word {
	return &Word{w, Signature(w)}
}

type RuneStats struct {
	Letter  rune
	Average float64
	Max     int
	Count   int
	Over3   int
	Over1   int

	wordCount int
}

func RuneDistribution(words []*Word) []*RuneStats {

	ld := make([]*RuneStats, 26)

	//for idx, l := range letters {
	//	ld[idx] = &RuneStats{Letter: l}
	//}
	////words, err := ReadFile(filePath)
	////if err != nil {
	////	return nil, 0, err
	////}
	//for _, w := range words {
	//	runeCount := w.RuneCount()
	//	for r, c := range runeCount {
	//		d := ld[r-'a']
	//		d.wordCount++
	//		d.Count += c
	//		if c > d.Max {
	//			d.Max = c
	//		}
	//		if c > 3 {
	//			d.Over3++
	//		}
	//		if c > 1 {
	//			d.Over1++
	//		}
	//	}
	//}
	//for _, l := range ld {
	//
	//	l.Average = float64(l.Count) / float64(l.wordCount)
	//}
	//sort.Slice(ld, func(i, j int) bool {
	//	return ld[i].Average > ld[j].Average
	//
	//})
	return ld
}
