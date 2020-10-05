package commands

import (
	"../railroad"
)

type Boundary interface {
	// 状態を表示する
	PrintStatus(railroad.RailroadWiring)
	// メッセージを表示する
	PrintMessage(string)
}

type TrackControlBoundary interface {
	Boundary

	// 線路を要求する
	ReqireTrack()
	// 線路を入力する
	InputTrack() int
}

type SpeedControlBoundary interface {
	// 新しい速さを要求する
	RequireNewSpeed()
	// 新しい速さを入力する
	InputNewSpeed() int
}
