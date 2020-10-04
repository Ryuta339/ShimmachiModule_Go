package railroad

import (
	"./direction"
)

type TrackOperator interface {
	// 線路を追加
	AddTrack(Track)
	// 線路が配線略図にあるかを判定する
	IsIn(int) bool
	// 速さが非負かつ上限以下か判定する
	IsSpeedInRange(int) bool
	// 速さを変更する
	ChangeSpeed(int, int) error
	// 方向を変更する
	ChangeDirection(int, direction.DirectionState) error
}
