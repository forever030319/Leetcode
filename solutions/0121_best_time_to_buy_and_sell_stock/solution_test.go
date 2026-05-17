package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"basic case", []int{7, 1, 5, 3, 6, 4}, 5},
		{"descending prices", []int{7, 6, 4, 3, 1}, 0},
		{"buy and sell at last day", []int{1, 2, 3, 4, 5}, 4},
		{"single element", []int{5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
