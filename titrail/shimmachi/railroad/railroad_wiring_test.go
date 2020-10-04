package railroad

import (
	"fmt"

	"./direction"
)

func ExampleRailroadWiring() {
	wiring := NewSimpleRailroadWiring()
	for idx := 0; idx < 3; idx++ {
		wiring.AddTrack(NewTrack(idx + 1))
	}

	err := wiring.ChangeDirection(0, direction.GetLeftInstance())
	if err != nil {
		panic(err)
	}
	err = wiring.ChangeDirection(2, direction.GetRightInstance())
	if err != nil {
		panic(err)
	}

	err = wiring.ChangeSpeed(0, 50)
	if err != nil {
		panic(err)
	}
	err = wiring.ChangeSpeed(2, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", wiring.ToString())

	err = wiring.ChangeDirection(1, direction.GetLeftInstance())
	if err != nil {
		panic(err)
	}
	err = wiring.ChangeDirection(0, direction.GetStopInstance())
	if err != nil {
		panic(err)
	}

	err = wiring.ChangeSpeed(0, 0)
	if err != nil {
		panic(err)
	}
	err = wiring.ChangeSpeed(1, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", wiring.ToString())
	// Output:
	// Track no. 1
	// 	Direction: Left
	// 	Speed: 50
	// Track no. 2
	// 	Direction: Stop
	// 	Speed: 0
	// Track no. 3
	// 	Direction: Right
	// 	Speed: 100
	// Track no. 1
	// 	Direction: Stop
	// 	Speed: 0
	// Track no. 2
	// 	Direction: Left
	// 	Speed: 100
	// Track no. 3
	// 	Direction: Right
	// 	Speed: 100
}
