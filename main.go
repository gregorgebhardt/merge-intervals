package main

import (
	"bufio"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/gregorgebhardt/interval-merger/interval"
)

func main() {
	var (
		filename string
		verbose  bool
	)

	flag.StringVarP(&filename, "file", "f", "", "file to read intervals from")
	flag.BoolVarP(&verbose, "verbose", "v", false, "print debug information")
	flag.Parse()

	var file *os.File
	var err error
	if filename != "" {
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	buffReader := bufio.NewReader(file)

	intervalMerger := interval.IntervalMerger{
		Verbose: verbose,
	}

	parser := interval.NewParser(buffReader)
	for interval := range parser.Intervals() {
		intervalMerger.Add(interval)
	}

	mergedIntervals := intervalMerger.Merge()

	fmt.Println(mergedIntervals)
}
