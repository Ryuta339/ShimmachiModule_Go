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
}

type SimpleTrackListenee struct {
	listeners []TrackListener
}

// リスナーを追加
func (t *SimpleTrackListenee) AddListener(listener TrackListener) {
	t.listeners = append(t.listeners, listener)
}

// リスナーを削除
func (t *SimpleTrackListenee) RemoveListener(listener TrackListener) error {
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
func (t *SimpleTrackListenee) NotifyListeners(event Event) {
	for _, v := range t.listeners {
		v.UpdateTrack(event)
	}
}

func NewSimpleTrackListenee() *SimpleTrackListenee {
	return &SimpleTrackListenee{listeners: make([]TrackListener, 0, 10)}
}

// 番線の基底クラス
// パワーユニットを実装したもの
type Track interface {
	TrackListenee
	ChangeableSpeed
	ChangeableDirection
	util.Stringify

	GetIndex() int
}

type SimpleTrack struct {
	*SimpleTrackListenee

	Index int

	// Has unexported fields
	speed int
	state direction.DirectionState
}

// 速さを変更する
func (t *SimpleTrack) ChangeSpeed(newSpeed int) {
	t.speed = newSpeed
}

// 方向を変更する
// 方向変更時に一旦速さをゼロにする
func (t *SimpleTrack) ChangeDirection(newState direction.DirectionState) {
	t.ChangeSpeed(0)
	t.state = newState
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

func NewTrack(index int) Track {
	return &SimpleTrack{
		SimpleTrackListenee: NewSimpleTrackListenee(),
		Index:               index,
		speed:               0,
		state:               direction.GetStopInstance(),
	}
}
