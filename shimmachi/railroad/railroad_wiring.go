package railroad

import (
	"errors"
	"fmt"
)

const (
	MAX_SPEED = 100
)

type RailroadWiring struct {
	Tracks []*Track
}

// 線路を追加
func (rw *RailroadWiring) AddTrack(track *Track) {
	rw.Tracks = append(rw.Tracks, track)
}

// 線路が配線略図にあるかを判定する
func (rw *RailroadWiring) IsIn(trackNumber int) bool {
	return 0 <= trackNumber && trackNumber < len(rw.Tracks)
}

// 速さが非負かつ上限以下か判定する
func (rw *RailroadWiring) IsInRange(speed int) bool {
	return 0 <= speed && speed <= MAX_SPEED
}

// 速さを変更する
func (rw *RailroadWiring) ChangeSpeed(trackNumber, newSpeed int) error {
	// エラーメッセージ
	if !rw.IsIn(trackNumber) {
		errmsg := fmt.Sprintf("Illegal track number: %d.", trackNumber)
		return errors.New(errmsg)
	}
	if !rw.IsInRange(newSpeed) {
		errmsg := fmt.Sprintf("Illegal speed: %d.", newSpeed)
		return errors.New(errmsg)
	}

	rw.Tracks[trackNumber].ChangeSpeed(newSpeed)
	return nil
}

// 文字列化
func (rw *RailroadWiring) ToString() string {
	m := make([]byte, 0, 2048)

	for _, v := range rw.Tracks {
		m = append(m, v.ToString()...)
		m = append(m, '\n')
	}

	return string(m)
}

func NewRailroadWiring() *RailroadWiring {
	tracks := make([]*Track, 0, 10)
	return &RailroadWiring{
		Tracks: tracks,
	}
}
