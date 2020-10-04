package display

import (
	"fmt"
	"io"
	"os"

	"../railroad"
)

type Displayable interface {
	railroad.TrackListener
}

type CuiDisplay struct {
	writer io.Writer
	wiring railroad.RailroadWiring
}

func (cd *CuiDisplay) UpdateTrack(event railroad.Event) {
	fmt.Fprintln(cd.writer, cd.wiring.ToString())
}

func NewCuiDisplay(writer io.Writer, wiring railroad.RailroadWiring) *CuiDisplay {
	return &CuiDisplay{
		writer: writer,
		wiring: wiring,
	}
}

func NewCuiDisplayWithStdout(wiring railroad.RailroadWiring) *CuiDisplay {
	return &CuiDisplay{
		writer: os.Stdout,
		wiring: wiring,
	}
}
