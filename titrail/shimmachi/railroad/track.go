package railroad

import (
	"errors"
	"strconv"

	"../util"
	"./direction"
)

type TrackListenee interface {
	AddListener(TrackListener)
	RemoveListener(TrackListener) error
	NotifyListeners(Event)
	WillChange()
	Changed()
}

// 番線の基底インターフェース
// パワーユニットを実装したもの
type Track interface {
	TrackListenee
	util.Stringify

	ChangeSpeed(int)
	ChangeDirection(direction.DirectionState)
	GetIndex() int
}

// 番線の単純な実装構造体
type SimpleTrack struct {
	Index int

	// Has unexported fields
	speed     int
	state     direction.DirectionState
	listeners []TrackListener
	depth     int
}

// リスナーを追加
func (t *SimpleTrack) AddListener(listener TrackListener) {
	t.listeners = append(t.listeners, listener)
}

// リスナーを削除
func (t *SimpleTrack) RemoveListener(listener TrackListener) error {
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

// リスナーに通知
func (t *SimpleTrack) NotifyListeners(event Event) {
	if t.depth == 1 {
		for _, v := range t.listeners {
			v.UpdateTrack(event)
		}
	}
}

// 設定変更前
func (t *SimpleTrack) WillChange() {
	t.depth++
}

// 設定変更後
func (t *SimpleTrack) Changed() {
	t.depth--
}

// 速さを変更する
func (t *SimpleTrack) ChangeSpeed(newSpeed int) {
	t.WillChange()
	defer t.Changed()
	t.speed = newSpeed
	t.NotifyListeners(NewChangeSpeedEvent(t))
}

// 方向を変更する
// 方向変更時に一旦速さをゼロにする
func (t *SimpleTrack) ChangeDirection(newState direction.DirectionState) {
	t.WillChange()
	defer t.Changed()
	t.ChangeSpeed(0)
	t.state = newState
	t.NotifyListeners(NewChangeDirectionEvent(t))
}

// 通し番号を取得する
func (t *SimpleTrack) GetIndex() int {
	return t.Index
}

// 文字列化
func (t *SimpleTrack) ToString() string {
	m := make([]byte, 0, 128)
	m = append(m, "Track no. "...)
	m = append(m, strconv.Itoa(t.Index)...)
	m = append(m, "\n\tDirection: "...)
	m = append(m, t.state.ToString()...)
	m = append(m, "\n\tSpeed: "...)
	m = append(m, strconv.Itoa(t.speed)...)
	return string(m)
}

func NewSimpleTrack(index int) Track {
	return &SimpleTrack{
		Index:     index,
		speed:     0,
		state:     direction.GetStopInstance(),
		listeners: make([]TrackListener, 0, 10),
		depth:     0,
	}
}
