package interval

import (
	"fmt"

	"github.com/gregorgebhardt/redblack"
)

type IntervalMerger struct {
	tree    redblack.Tree[Interval, Interval]
	Verbose bool
}

func NewIntervalMerger(intervals []Interval) (*IntervalMerger, error) {
	tree, err := redblack.NewTree[Interval, Interval](intervals, true)
	if err != nil {
		return nil, err
	}

	return &IntervalMerger{
		tree: *tree,
	}, nil
}

func (i *IntervalMerger) Add(interval Interval) {
	err := i.tree.Insert(interval)
	if i.Verbose {
		switch err {
		case nil:
			fmt.Println("Adding   " + interval.String())
		case redblack.KeyExistsError:
			fmt.Println("Skipping " + interval.String())
		default:
			fmt.Println("Error    " + interval.String() + " " + err.Error())
		}
	}
}

func (i *IntervalMerger) Merge() IntervalSlice {
	if i.Verbose {
		fmt.Println(i.tree.String())
	}
	mergedIntervals := make([]Interval, 0)
	for interval := range i.tree.Sorted() {
		if len(mergedIntervals) == 0 {
			mergedIntervals = append(mergedIntervals, interval)
			continue
		}

		if !mergedIntervals[len(mergedIntervals)-1].Merge(interval) {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}
	return mergedIntervals
}

func (i *IntervalMerger) TreeString() string {
	return i.tree.String()
}
