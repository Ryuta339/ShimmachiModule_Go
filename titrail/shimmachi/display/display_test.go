package display

import (
	"../railroad"
)

const (
	numTrack = 3
)

func ExampleCuiDisplay() {
	var display Displayable

	wiring := railroad.NewSimpleRailroadWiring()
	display = NewCuiDisplayWithStdout(wiring)
	for idx := 0; idx < numTrack; idx++ {
		track := railroad.NewSimpleTrack(idx + 1)
		track.AddListener(display)
		wiring.AddTrack(track)
	}

	// wiring.Tracks[0].NotifyListeners(railroad.NewNullEvent(nil))
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
