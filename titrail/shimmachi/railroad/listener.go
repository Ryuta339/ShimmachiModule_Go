package railroad

type RailroadWiringListener interface {
	Update(wiring *RailroadWiring)
}

type TrackListener interface {
	Update(track *Track)
}
