package jenkins

import "testing"

func TestAdd(t *testing.T) {
	test := []struct {
		x      int
		y      int
		answer int
	}{
		{1, 1, 2},
		{0, 1, 1},
		{3, 3, 6},
	}

	for _, tt := range test {
		if add(tt.x, tt.y) != tt.answer {
			t.Errorf("%d + %d got %d, expected %d", tt.x, tt.y, tt.answer, tt.x+tt.y)
		}
	}
}
