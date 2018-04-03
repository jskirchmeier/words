package words

import (
	"bufio"
	"fmt"
	"github.com/jskirchmeier/stream"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CombineLists(paths ...string) []string {

	wordMap := make(map[string]bool)

	for _, p := range paths {
		source, err := stream.FileSource(p)

		if err != nil {
			fmt.Println("Unable to open file : ", err)
			continue
		}
		s := &stream.Stream{}
		stream.Process(source, s.
			Modify(strings.TrimSpace).
			Modify(strings.ToLower).
			Filter(func(s string) bool {
				return !strings.ContainsAny(s, "-_ %'`!#@")
			}).
			ActOn(func(s string) {
				wordMap[s] = true
			}))
	}
	wordList := make([]string, len(wordMap))
	i := 0
	for k := range wordMap {
		wordList[i] = k
		i++
	}
	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i] < wordList[j]
	})
	return wordList
}

func WordLengths(wordList []string) map[int]int {
	lengthCounts := make(map[int]int)
	for _, w := range wordList {
		l := len(w)
		lengthCounts[l] = lengthCounts[l] + 1
	}
	return lengthCounts
}

func ToFile(path string, words []*Word) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	defer f.Close()
	for _, word := range words {
		w.WriteString(word.ToString())
		w.WriteString("\n")
	}
	return w.Flush()
}

func FromFile(path string) ([]*Word, error) {

	var words []*Word
	source, err := stream.FileSource(path)

	if err != nil {
		return nil, err
	}
	s := &stream.Stream{}
	stream.Process(source, s.Modify(strings.TrimSpace).ActOn(func(s string) {
		fields := strings.Fields(s)
		sig, err := strconv.ParseUint(fields[1], 16, 64)
		if err != nil {
			panic("Unable to parse signature (" + s + ")  in file " + path)
		}
		words = append(words, &Word{fields[0], sig})
	}))
	return words, nil
}
