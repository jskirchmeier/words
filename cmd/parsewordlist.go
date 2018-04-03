package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func addWords(in io.Reader, words map[string]bool) {
	scanner := bufio.NewScanner(in)

	nWords := len(words)
	count := 0
	tossed := 0
	for scanner.Scan() {

		word := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if len(word) < 3 {
			tossed++
			continue
		}

		//if words[word] {
		//	fmt.Println("Dup word : ", word)
		//}
		count++
		words[word] = true
	}
	fmt.Printf("Words in file : %10d\n", count+tossed)
	fmt.Printf("Words tossed  : %10d\n", tossed)
	fmt.Printf("Words kept    : %10d\n", count)
	fmt.Printf("Words added   : %10d\n", len(words)-nWords)
}

func parseFile(filePath string, words map[string]bool) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(filePath)
	addWords(file, words)
	fmt.Println()

}

func main() {
	wordList := make(map[string]bool)

	parseFile("c:/stuff/5000Words.txt", wordList)
	parseFile("c:/stuff/words/list.txt", wordList)
	parseFile("c:/stuff/words/twist.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/A Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/B Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/C Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/D Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/E Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/F Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/G Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/H Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/I Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/J Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/K Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/L Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/M Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/N Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/O Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/P Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/Q Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/R Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/S Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/T Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/U Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/V Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/W Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/X Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/Y Words.txt", wordList)
	parseFile("C:/stuff/words/EOWL-v1.1.2/LF Delimited Format/Z Words.txt", wordList)

	fmt.Println(len(wordList), "words")

	byLen := make(map[int]int)
	byLetter := make(map[string]int)

	for w := range wordList {
		l := len(w)
		byLen[l] = byLen[l] + 1

		s := w[:1]
		byLetter[s] = byLetter[s] + 1
	}

	fmt.Println("By length")
	for k, v := range byLen {
		fmt.Printf("%2d   %6d\n", k, v)
	}

	fmt.Println()
	fmt.Println("by starting letter")
	for l, c := range byLetter {
		fmt.Printf("%s   %6d\n", l, c)

	}

}
