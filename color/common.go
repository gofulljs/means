package color

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Mgr struct {
	output  io.Writer
	formats []string
}

type Options func(mgr *Mgr)

func Color(opts ...Options) *Mgr {
	mgr := &Mgr{
		output: os.Stdout,
	}

	for _, opt := range opts {
		opt(mgr)
	}

	return mgr
}

func (mgr *Mgr) Format() string {
	return strings.Join(mgr.formats, ";")
}

func (mgr *Mgr) Output(message string) {
	fmt.Fprintf(mgr.output, "\x1b[%sm%s\x1b[0m", mgr.Format(), message)
}

func WithOutput(output io.Writer) Options {
	return func(mgr *Mgr) {
		mgr.output = output
	}
}

func WithRGB(r, g, b uint8) Options {
	return func(mgr *Mgr) {
		mgr.formats = append(mgr.formats, RGBFormat(RGB{r, g, b}))
	}
}

func WithHex(hex uint32) Options {
	return func(mgr *Mgr) {
		mgr.formats = append(mgr.formats, HexFormat(hex))
	}
}

func WithStyle(styles ...StyleFn) Options {
	return func(mgr *Mgr) {
		for _, style := range styles {
			mgr.formats = append(mgr.formats, style())
		}
	}
}
