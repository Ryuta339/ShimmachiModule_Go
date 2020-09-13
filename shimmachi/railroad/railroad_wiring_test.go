package railroad

import "fmt"

func ExampleRailroadWiring() {
	wiring := NewRailroadWiring()
	for idx := 0; idx < 3; idx++ {
		wiring.AddTrack(NewTrack(idx + 1))
	}
	err := wiring.ChangeSpeed(0, 50)
	if err != nil {
		panic(err)
	}
	err = wiring.ChangeSpeed(2, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", wiring.ToString())
	// Output:
	// Track no. 1
	// 	Speed: 50
	// Track no. 2
	// 	Speed: 0
	// Track no. 3
	// 	Speed: 100
}
