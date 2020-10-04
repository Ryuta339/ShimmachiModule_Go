package railroad

type RailroadWiringListener interface {
	Update(wiring RailroadWiring, event Event)
}

type TrackListener interface {
	UpdateTrack(event Event)
}
