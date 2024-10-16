package interval

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
)

type Parser struct {
	scanner *bufio.Scanner
}

// NewParser returns a new Parser that reads from the given reader
// Note that the parser does not read from the reader until Intervals() is called.
// So do not close the reader before having interated over all intervals.
func NewParser(reader io.Reader) *Parser {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitInterval)
	return &Parser{
		scanner: scanner,
	}
}

// Intervals returns an iterator over the intervals in the input
func (p *Parser) Intervals() iter.Seq[Interval] {
	return func(yield func(Interval) bool) {
		for p.scanner.Scan() {
			intervalBytes := p.scanner.Bytes()
			interval := Interval{}
			err := interval.UnmarshalText(intervalBytes)
			if err != nil {
				panic("invalid interval")
			}
			if !yield(interval) {
				return
			}
		}
	}
}

// splitInterval is a split function for a Scanner that returns each interval in the input as a token.
// The byte-slice is split before the next opening bracket '[' and after the closing bracket ']'.
func splitInterval(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if j := bytes.IndexByte(data, ']'); j >= 0 {
		i := bytes.IndexByte(data[:j], '[')
		if i == -1 {
			return 0, nil, fmt.Errorf("invalid interval")
		}
		// We have a full newline-terminated line.
		return j + 1, data[i : j+1], nil
	}

	// Request more data.
	return 0, nil, nil
}
