package commands

import (
	"../railroad"
)

type Commands interface {
	// 操作する
	Operate()
}

// 線路に関するコマンド
type TrackCommands interface {
	Commands
}

// 速度コントロール
type SpeedControl struct {
	// Unexported fields
	wiring railroad.RailroadWiring
}
