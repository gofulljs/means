package color

import (
	"testing"
)

func TestMgr_Output(t *testing.T) {

	tests := []struct {
		name string
		opts []Options
	}{
		{
			name: "base",
			opts: []Options{
				WithRGB(255, 255, 66),
			},
		},
		{
			name: "base + bold",
			opts: []Options{
				WithRGB(255, 255, 66),
				WithStyle(BoldFormat),
			},
		},
		{
			name: "base + Italic",
			opts: []Options{
				WithRGB(255, 255, 66),
				WithStyle(ItalicFormat),
			},
		},
		{
			name: "base + bold + Italic",
			opts: []Options{
				WithRGB(255, 255, 66),
				WithStyle(BoldFormat),
				WithStyle(ItalicFormat),
			},
		},
		{
			name: "baseHex + bold + Italic",
			opts: []Options{
				WithHex(0xffff42),
				WithStyle(BoldFormat),
				WithStyle(ItalicFormat),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := Color(tt.opts...)
			mgr.Output("Hello World!")
		})
	}

}
