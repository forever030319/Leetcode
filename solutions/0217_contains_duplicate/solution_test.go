package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"has duplicate", []int{1, 2, 3, 1}, true},
		{"no duplicate", []int{1, 2, 3, 4}, false},
		{"all same", []int{1, 1, 1, 1}, true},
		{"two elements duplicate", []int{1, 1}, true},
		{"two elements no duplicate", []int{1, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
