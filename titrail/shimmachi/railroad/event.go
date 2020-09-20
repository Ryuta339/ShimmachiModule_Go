package railroad

type Event interface {
	GetSource() interface{}
	GetType() int
	ToString() string
}

// Event を実装したクラス
// 実際は何もしない
type NullEvent struct {
	source interface{}
}

func (ne *NullEvent) GetSource() interface{} {
	return ne.source
}

func (ne *NullEvent) GetType() int {
	return -1
}

func (ne *NullEvent) ToString() string {
	return "Null Event"
}

func NewNullEvent(source interface{}) Event {
	return &NullEvent{source: source}
}
