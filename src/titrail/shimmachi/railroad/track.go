package railroad

import (
	"strconv"

	"./direction"
)

type Track struct {
	Index int

	// Has unexported fields
	speed int
	state direction.DirectionState
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
