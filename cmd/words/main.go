package main

import (
	"fmt"
	"github.com/jskirchmeier/words"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
	"path/filepath"
	"time"
)

func main() {
	p := message.NewPrinter(language.English)

	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Various word analyses utilities",
		Long: `
		`,
	}

	letterDistribution := &cobra.Command{
		Use:   "letters [path to word list file]",
		Short: "letters [path to word list file]",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Determining letter distribution for : ", args[0])
			wordList, err := words.ReadFile(args[0])
			if err != nil {
				fmt.Println("Unable to process request : ", err)
				return
			}
			dist := words.RuneDistribution(wordList)
			p.Printf("Words                    : %15d\n", len(wordList))

			for _, l := range dist {
				p.Printf("%c   %6.3f  %4d  %10d %5d  %8d\n", l.Letter, l.Average, l.Max, l.Count, l.Over3, l.Over1)
			}
		},
	}
	signature := &cobra.Command{
		Use:   "signature [path to word list file]",
		Short: "signature [path to word list file]",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Calculating word signatures for : ", args[0])
			wordList, err := words.ReadFile(args[0])
			if err != nil {
				fmt.Println("Unable to process request : ", err)
				return
			}
			p.Printf("Words                    : %15d\n", len(wordList))
			for _, w := range wordList {
				w.Signature()
			}

		},
	}

	rootCmd.AddCommand(letterDistribution, signature)

	ts := time.Now()
	rootCmd.Execute()
	elapsed := time.Now().Sub(ts)
	fmt.Println("Procedure took : ", elapsed)

}
