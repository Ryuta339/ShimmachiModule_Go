package direction

// 各線路の進行方向を表すクラスの基底。
// State パターン で実装している。

type DirectionState interface {
	ToString() string
}

var leftInstance = newLeftDirection()
var rightInstance = newRightDirection()
var stopInstance = newStopDirection()

// 左向きに進む状態を表す。
type leftDirection struct {
}

func (ld *leftDirection) ToString() string {
	return "Left"
}

func GetLeftInstance() DirectionState {
	return leftInstance
}

func newLeftDirection() DirectionState {
	return &leftDirection{}
}

// 右向きに進む状態を表す。
type rightDirection struct {
}

func (rd *rightDirection) ToString() string {
	return "Right"
}

func GetRightInstance() DirectionState {
	return rightInstance
}

func newRightDirection() DirectionState {
	return &rightDirection{}
}

// 停止状態を表す。
type stopDirection struct {
}

func (rd *stopDirection) ToString() string {
	return "Stop"
}

func GetStopInstance() DirectionState {
	return stopInstance
}

func newStopDirection() DirectionState {
	return &stopDirection{}
}
