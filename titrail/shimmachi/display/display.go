package display

import (
	"fmt"
	"io"

	"../railroad"
)

type Displayable interface {
	railroad.RailroadWiringListener
}

type CuiDisplay struct {
	writer io.Writer
}

func (cd *CuiDisplay) Update(
	wiring *railroad.RailroadWiring,
	event railroad.Event,
) {
	fmt.Fprintln(cd.writer, wiring.ToString())
}

func NewCuiDisplay(writer io.Writer) *CuiDisplay {
	return &CuiDisplay{writer: writer}
}