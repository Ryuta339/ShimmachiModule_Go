package railroad

import (
	"errors"
	"fmt"

	"../util"
	"./direction"
)

const (
	MAX_SPEED = 100
)

type RailroadWiringListenee interface {
	AddListener(RailroadWiringListener)
	RemoveListener(RailroadWiringListener) error
	NotifyListeners(Event)
}

// 操作盤の基底インターフェース
type RailroadWiring interface {
	RailroadWiringListenee
	TrackOperator
	util.Stringify
}

type SimpleRailroadWiring struct {
	Tracks []Track

	// Has unexported fields
	listeners []RailroadWiringListener
}

// リスナーを追加
func (rw *SimpleRailroadWiring) AddListener(listener RailroadWiringListener) {
	rw.listeners = append(rw.listeners, listener)
}

// リスナーを削除
func (rw *SimpleRailroadWiring) RemoveListener(listener RailroadWiringListener) error {
	idx := -1
	// 該当するリスナーを探索
	for i, v := range rw.listeners {
		if listener == v {
			idx = i
			break
		}
	}
	if idx == -1 {
		// 該当するリスナーが存在しない
		return errors.New("No such listener.")
	}

	// 該当要素を削除
	rw.listeners[idx] = rw.listeners[len(rw.listeners)-1]
	rw.listeners = rw.listeners[:len(rw.listeners)-1]

	return nil // 成功
}

// リスナーに通知
func (rw *SimpleRailroadWiring) NotifyListeners(event Event) {
	for _, v := range rw.listeners {
		v.Update(rw, event)
	}
}

// 線路を追加
func (rw *SimpleRailroadWiring) AddTrack(track Track) {
	rw.Tracks = append(rw.Tracks, track)
}

// 線路が配線略図にあるかを判定する
func (rw *SimpleRailroadWiring) IsIn(trackNumber int) bool {
	return 0 <= trackNumber && trackNumber < len(rw.Tracks)
}

// 速さが非負かつ上限以下か判定する
func (rw *SimpleRailroadWiring) IsSpeedInRange(speed int) bool {
	return 0 <= speed && speed <= MAX_SPEED
}

// 速さを変更する
func (rw *SimpleRailroadWiring) ChangeSpeed(trackNumber, newSpeed int) error {
	// エラー処理
	if !rw.IsIn(trackNumber) {
		errmsg := fmt.Sprintf("Illegal track number: %d.", trackNumber)
		return errors.New(errmsg)
	}
	if !rw.IsSpeedInRange(newSpeed) {
		errmsg := fmt.Sprintf("Illegal speed: %d.", newSpeed)
		return errors.New(errmsg)
	}

	rw.Tracks[trackNumber].ChangeSpeed(newSpeed)
	return nil
}

// 方向を変更する
func (rw *SimpleRailroadWiring) ChangeDirection(
	trackNumber int,
	newState direction.DirectionState,
) error {
	// エラー処理
	if !rw.IsIn(trackNumber) {
		errmsg := fmt.Sprintf("Illegal track number :%d.", trackNumber)
		return errors.New(errmsg)
	}

	rw.Tracks[trackNumber].ChangeDirection(newState)
	return nil
}

// 文字列化
func (rw *SimpleRailroadWiring) ToString() string {
	m := make([]byte, 0, 2048)

	m = append(m, "================================\n"...)
	for _, v := range rw.Tracks {
		m = append(m, v.ToString()...)
		m = append(m, '\n')
	}
	m = append(m, "================================\n"...)

	return string(m)
}

// コンストラクタ
func NewSimpleRailroadWiring() RailroadWiring {
	tracks := make([]Track, 0, 10)
	return &SimpleRailroadWiring{
		Tracks: tracks,
	}
}
