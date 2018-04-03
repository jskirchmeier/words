package main

import (
	"fmt"
	"github.com/jskirchmeier/stream"
	"github.com/jskirchmeier/words"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

var (
	fileList = []string{
		"C:/stuff/words/twist.txt",
		"C:/stuff/words/google-10000-english/google-10000-english-usa-no-swears.txt",
		"C:/stuff/words/list.txt",
		"C:/stuff/words/American/2of12inf.txt",
		"C:/stuff/words/American/2of12.txt",
	}
)

func a() {
	p := message.NewPrinter(language.English)

	list := words.CombineLists(fileList...)
	p.Printf("%10d words in list\n", len(list))

	signatures := make(map[uint64]int)
	source := stream.StaticSource(list...)
	s := &stream.Stream{}
	stream.Process(source, s.
		Filter(stream.MinLen(3)).
		Filter(stream.MaxLen(7)).
		Filter(func(s string) bool {
			return !strings.ContainsAny(s, "-_ %'`!#@")
		}).
		ActOn(func(s string) {
			if len(s) == 7 {
				sig := words.Signature(s)
				signatures[sig] = signatures[sig] + 1
			}
		}).
		Accumulate())

	p.Printf("%10d words in filtered list\n", len(s.Strings))

	lenDist := words.WordLengths(s.Strings)
	for l, c := range lenDist {
		p.Printf("%2d  : %10d\n", l, c)
	}

	countDist := make(map[int]int)

	count := 0
	count1 := 0
	count2 := 0
	max := 0
	for _, c := range signatures {
		countDist[c] = countDist[c] + 1
		if c == 1 {
			count1++
		}
		if c > 1 {
			count++
		}
		if c > 2 {
			count2++
		}
		if c > max {
			max = c
		}
	}
	p.Printf("%10d distinct signature\n", len(signatures))
	p.Printf("%10d signature with multiple seven char words\n", count)
	p.Printf("%10d signature with one seven char word\n", count1)
	p.Printf("%10d signature with two seven char words\n", count2)
	p.Printf("%10d max number of words with same signature\n", max)

	for i := 1; i < 6; i++ {
		p.Printf("%d  : %10d\n", i, countDist[i])

	}
}

func createWordList(path string) {
	p := message.NewPrinter(language.English)
	list := words.CombineLists(fileList...)
	items := make([]*words.Word, len(list))

	for idx, w := range list {
		items[idx] = words.NewWord(w)
	}

	p.Printf("%10d words in list\n", len(list))
	err := words.ToFile(path, items)
	fmt.Println("error : ", err)
}

func fromFile(path string) {
	items, err := words.FromFile(path)
	if err != nil {
		panic(err)
	}
	p := message.NewPrinter(language.English)
	p.Printf("%10d words in list\n", len(items))
}

func main() {
	//createWordList("C:/stuff/words/words.list")
	fromFile("C:/stuff/words/words.list")
}
