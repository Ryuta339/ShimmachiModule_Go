package railroad

type RailroadWiringListener interface {
	Update(wiring *RailroadWiring, event Event)
}

type TrackListener interface {
	Update(track *Track, event Event)
}
