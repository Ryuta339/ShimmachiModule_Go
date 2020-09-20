package railroad

import (
	"errors"
	"strconv"

	"./direction"
)

type Track struct {
	Index int

	// Has unexported fields
	speed     int
	state     direction.DirectionState
	listeners []TrackListener
}

// リスナーを追加
func (t *Track) AddListener(listener TrackListener) {
	t.listeners = append(t.listeners, listener)
}

// リスナーを削除
func (t *Track) RemoveListener(listener TrackListener) error {
	idx := -1
	// 該当するリスナーを探索
	for i, v := range t.listeners {
		if listener == v {
			idx = i
			break
		}
	}
	if idx == -1 {
		// 該当するリスナーが存在しない
		return errors.New("No such listener.")
	}

	// 該当要素を成功
	t.listeners[idx] = t.listeners[len(t.listeners)-1]
	t.listeners = t.listeners[:len(t.listeners)-1]

	return nil // 成功
}

// リスナーに追加
func (t *Track) NotifyListeners(event Event) {
	for _, v := range t.listeners {
		v.Update(t, event)
	}
}

// 速さを変更する
func (t *Track) ChangeSpeed(newSpeed int) {
	t.speed = newSpeed
}

// 方向を変更する
// 方向変更時に一旦速さをゼロにする
func (t *Track) ChangeDirection(newState direction.DirectionState) {
	t.ChangeSpeed(0)
	t.state = newState
}

// 文字列化
func (t *Track) ToString() string {
	m := make([]byte, 0, 128)
	m = append(m, "Track no. "...)
	m = append(m, strconv.Itoa(t.Index)...)
	m = append(m, "\n\tDirection: "...)
	m = append(m, t.state.ToString()...)
	m = append(m, "\n\tSpeed: "...)
	m = append(m, strconv.Itoa(t.speed)...)
	return string(m)
}

func NewTrack(index int) *Track {
	return &Track{
		Index: index,
		speed: 0,
		state: direction.GetStopInstance(),
	}
}
