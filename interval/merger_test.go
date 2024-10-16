package interval

import (
	"reflect"
	"testing"

	"github.com/gregorgebhardt/redblack"
)

func TestIntervalMerger_Merge(t *testing.T) {
	type fields struct {
		tree redblack.Tree[Interval, Interval]
	}
	tests := []struct {
		name  string
		input IntervalSlice
		want  IntervalSlice
	}{
		// TODO: Add test cases.
		{"Empty", IntervalSlice{}, IntervalSlice{}},
		{"Single", IntervalSlice{{1, 2}}, IntervalSlice{{1, 2}}},
		{"Two Intersecting", IntervalSlice{{1, 2}, {2, 3}}, IntervalSlice{{1, 3}}},
		{"Two Non-Intersecting", IntervalSlice{{1, 2}, {3, 4}}, IntervalSlice{{1, 2}, {3, 4}}},
		{"Two Equal", IntervalSlice{{1, 2}, {1, 2}}, IntervalSlice{{1, 2}}},
		{"Two Intersecting Rev", IntervalSlice{{2, 3}, {1, 2}}, IntervalSlice{{1, 3}}},
		{"Two Non-Intersecting Rev", IntervalSlice{{3, 4}, {1, 2}}, IntervalSlice{{1, 2}, {3, 4}}},
		{"Three Intersecting", IntervalSlice{{1, 2}, {2, 3}, {3, 4}}, IntervalSlice{{1, 4}}},
		{"Three Non-Intersecting", IntervalSlice{{1, 2}, {3, 4}, {5, 6}}, IntervalSlice{{1, 2}, {3, 4}, {5, 6}}},
		{"Three Mixed", IntervalSlice{{1, 2}, {3, 4}, {2, 3}}, IntervalSlice{{1, 4}}},
		{"Three Mixed", IntervalSlice{{1, 3}, {2, 4}, {5, 9}}, IntervalSlice{{1, 4}, {5, 9}}},
		{"Four Mixed", IntervalSlice{{1, 3}, {2, 4}, {5, 9}, {10, 11}}, IntervalSlice{{1, 4}, {5, 9}, {10, 11}}},
		{"Negative", IntervalSlice{{-1, 3}, {2, 4}, {5, 12}, {10, 11}}, IntervalSlice{{-1, 4}, {5, 12}}},
		{"Decreasing2", IntervalSlice{{12, 16}, {10, 15}, {5, 12}, {4, 11}}, IntervalSlice{{4, 16}}},
		{"Duplicates", IntervalSlice{{1, 2}, {1, 2}, {1, 2}}, IntervalSlice{{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merger, err := NewIntervalMerger(tt.input)
			if err != nil {
				t.Errorf("IntervalMerger.Merge() error = %v", err)
			}

			if got := merger.Merge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntervalMerger.Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
