package display

import (
	"os"

	"../railroad"
)

const (
	numTrack = 3
)

func ExampleCuiDisplay() {
	var display Displayable

	wiring := railroad.NewRailroadWiring()
	for idx := 0; idx < numTrack; idx++ {
		wiring.AddTrack(railroad.NewTrack(idx + 1))
	}

	display = NewCuiDisplay(os.Stdout)
	wiring.AddListener(display)
	wiring.NotifyListeners(railroad.NewNullEvent(nil))
	// Output:
	// Track no. 1
	// 	Direction: Stop
	// 	Speed: 0
	// Track no. 2
	// 	Direction: Stop
	// 	Speed: 0
	// Track no. 3
	// 	Direction: Stop
	// 	Speed: 0
	//
}
