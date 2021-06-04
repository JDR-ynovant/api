package engine

import "testing"

func TestRandomCoordinates(t *testing.T) {
	const (
		MAX_WIDTH  = 15
		MAX_HEIGHT = 15
	)

	for i := 0; i < 100; i++ {
		randGeneratedX, randGeneratedY := randomCoordinates(MAX_WIDTH, MAX_HEIGHT)
		if !(0 <= randGeneratedX && randGeneratedX <= MAX_WIDTH) || !(0 <= randGeneratedY && randGeneratedY <= MAX_HEIGHT) {
			t.Errorf("%d or %d is not in expected range", randGeneratedX, randGeneratedY)
		}
	}
}
