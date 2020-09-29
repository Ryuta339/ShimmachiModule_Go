package railroad

import "strconv"

type EventType string

const (
	Nulle           = EventType("Null")
	ChangeSpeed     = EventType("ChangeSpeed")
	ChangeDirection = EventType("ChangeDirection")
)

type Event interface {
	GetSource() interface{} // 型安全じゃないんだけどどうしよう
	GetType() EventType
	ToString() string
}

// Event を実装したクラス
// 実際は何もしない
// デバッグ用
type NullEvent struct {
	source interface{}
}

func (ne *NullEvent) GetSource() interface{} {
	return ne.source
}

func (ne *NullEvent) GetType() EventType {
	return Nulle
}

func (ne *NullEvent) ToString() string {
	return "Null Event"
}

func NewNullEvent(source interface{}) Event {
	return &NullEvent{source: source}
}

// 番線に関するイベント
type TrackEvent struct {
	// Has unexported fields

	// イベントが起こった番線
	track Track
	// イベントの種類
	// 後で直す
	typ EventType
}

func (te *TrackEvent) GetSource() interface{} {
	return te.track
}

func (te *TrackEvent) GetType() EventType {
	return te.typ
}

func (te *TrackEvent) ToString() string {
	str := "Track no. "
	str += strconv.Itoa(te.track.GetIndex())
	str += ": "
	str += string(te.typ)
	return str
}

// 速さ変更のイベント
func NewChangeSpeedEvent(track Track) Event {
	return &TrackEvent{
		track: track,
		typ:   ChangeSpeed,
	}
}

// 方向変更のイベント
func NewChangeDirectionEvent(track Track) Event {
	return &TrackEvent{
		track: track,
		typ:   ChangeDirection,
	}
}
