package interval

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%d,%d]", i.Start, i.End)
}

func (i *Interval) UnmarshalText(text []byte) error {
	values := [2]int{}
	err := json.Unmarshal(text, &values)
	if err != nil {
		return fmt.Errorf("invalid interval: %s", text)
	}
	i.Start = values[0]
	i.End = values[1]
	// Assume interval with Start > End is the same as [End, Start]
	if i.Start > i.End {
		i.Start, i.End = i.End, i.Start
	}
	return nil
}

// Implement the redblack.Comparable interface

// CompareTo returns -1 if i is less than other, 0 if they are equal, and 1 if i is greater than other
// The comparison is done by comparing the start values first and then the end values
func (i Interval) CompareTo(other Interval) int {
	if i.Start < other.Start {
		return -1
	}
	if i.Start > other.Start {
		return 1
	}
	if i.End < other.End {
		return -1
	}
	if i.End > other.End {
		return 1
	}
	return 0
}

func (i Interval) Value() Interval {
	return i
}

func (i Interval) Intersect(other Interval) bool {
	return i.Start <= other.End && i.End >= other.Start
}

func (i *Interval) Merge(other Interval) bool {
	if i.Intersect(other) {
		i.Start = min(i.Start, other.Start)
		i.End = max(i.End, other.End)
		return true
	}
	return false
}

// Convenience function for pretty printing interval slices

type IntervalSlice []Interval

func (s IntervalSlice) ToStrings() []string {
	strings := make([]string, len(s))
	for i, interval := range s {
		strings[i] = interval.String()
	}
	return strings
}

func (s IntervalSlice) String() string {
	return strings.Join(s.ToStrings(), " ")
}
