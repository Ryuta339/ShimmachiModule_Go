package railroad

import "strconv"

type Track struct {
	Index int

	// Has unexported fields
	speed int
}

// 速さを変更する
func (t *Track) ChangeSpeed(newSpeed int) {
	t.speed = newSpeed
}

// 文字列化
func (t *Track) ToString() string {
	m := make([]byte, 0, 128)
	m = append(m, "Track no. "...)
	m = append(m, strconv.Itoa(t.Index)...)
	m = append(m, "\n\tSpeed: "...)
	m = append(m, strconv.Itoa(t.speed)...)
	return string(m)
}

func NewTrack(index int) *Track {
	return &Track{
		Index: index,
		speed: 0,
	}
}
