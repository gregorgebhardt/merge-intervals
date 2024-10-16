package interval

import (
	"testing"
)

func TestInterval_String(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     string
	}{
		{"empty", Interval{}, "[0,0]"},
		{"start 1 end 2", Interval{1, 2}, "[1,2]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.interval.String(); got != tt.want {
				t.Errorf("Interval.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    Interval
		wantErr bool
	}{
		{"empty", "", Interval{}, true},
		{"invalid", "[1,2", Interval{}, true},
		{"valid", "[1,2]", Interval{1, 2}, false},
		{"valid reverse", "[2,1]", Interval{1, 2}, false},
		{"valid with spaces", " [ 1 , 2 ] ", Interval{1, 2}, false},
		{"valid with tabs", "\t[\t1\t,\t2\t]\t", Interval{1, 2}, false},
		{"valid with negative numbers", "[-1,-2]", Interval{-2, -1}, false},
		{"valid with negative and positive numbers", "[-1,2]", Interval{-1, 2}, false},
		{"valid with negative and positive numbers reverse", "[2,-1]", Interval{-1, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interval := Interval{}
			if err := interval.UnmarshalText([]byte(tt.text)); (err != nil) != tt.wantErr {
				t.Errorf("Interval.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				if interval.Start != tt.want.Start || interval.End != tt.want.End {
					t.Errorf("Interval.UnmarshalText() = %v, want %v", interval, tt.want)
				}
			}
		})
	}
}

func TestInterval_Intersect(t *testing.T) {
	tests := []struct {
		name      string
		intervalA Interval
		intervalB Interval
		want      bool
	}{
		{"A | B", Interval{1, 2}, Interval{2, 3}, true},
		{"A / B", Interval{1, 2}, Interval{3, 4}, false},
		{"A = B", Interval{1, 2}, Interval{1, 2}, true},
		{"A < B", Interval{1, 2}, Interval{1, 3}, true},
		{"B | A", Interval{1, 2}, Interval{0, 1}, true},
		{"B / A", Interval{1, 2}, Interval{0, 0}, false},
		{"A | 0", Interval{1, 2}, Interval{2, 2}, true},
		{"A > B", Interval{3, 7}, Interval{2, 5}, true},
		{"A c B", Interval{0, 5}, Interval{1, 3}, true},
		{"B c A", Interval{1, 3}, Interval{0, 5}, true},
		{"0 | B", Interval{0, 0}, Interval{0, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.intervalA.Intersect(tt.intervalB); got != tt.want {
				t.Errorf("Interval.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Merge(t *testing.T) {
	tests := []struct {
		name      string
		intervalA Interval
		intervalB Interval
		want      bool
		expectedA Interval
	}{
		{"A | B", Interval{1, 2}, Interval{2, 3}, true, Interval{1, 3}},
		{"A / B", Interval{1, 2}, Interval{3, 4}, false, Interval{1, 2}},
		{"A = B", Interval{1, 2}, Interval{1, 2}, true, Interval{1, 2}},
		{"A < B", Interval{1, 2}, Interval{1, 3}, true, Interval{1, 3}},
		{"B | A", Interval{1, 2}, Interval{0, 1}, true, Interval{0, 2}},
		{"B / A", Interval{1, 2}, Interval{0, -3}, false, Interval{1, 2}},
		{"A | 0", Interval{1, 2}, Interval{2, 2}, true, Interval{1, 2}},
		{"A > B", Interval{3, 7}, Interval{2, 5}, true, Interval{2, 7}},
		{"A c B", Interval{0, 5}, Interval{1, 3}, true, Interval{0, 5}},
		{"B c A", Interval{1, 3}, Interval{0, 5}, true, Interval{0, 5}},
		{"0 | B", Interval{0, 0}, Interval{0, 1}, true, Interval{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.intervalA.Merge(tt.intervalB); got != tt.want {
				t.Errorf("Interval.Intersect() = %v, want %v", got, tt.want)
			}
			if tt.intervalA.CompareTo(tt.expectedA) != 0 {
				t.Errorf("Interval.Merge() = %v, want %v", tt.intervalA, tt.expectedA)
			}
		})
	}
}
