package color

import (
	"fmt"
	"os"
	"testing"
)

func TestRGBFormat(t *testing.T) {
	tests := []struct {
		name string
		args RGB
	}{
		{
			name: "base",
			args: RGB{
				r: 255,
				g: 255,
				b: 66,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := OnlyRGBForamt(tt.args)
			fmt.Printf("%sHello World!%s\n", gotLeft, gotRight)
		})
	}
}

func TestRGBOutput(t *testing.T) {
	type args struct {
		rgb     RGB
		message string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{
			name: "base",
			args: args{
				rgb: RGB{
					r: 255,
					g: 255,
					b: 66,
				},
				message: "Hello World!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RGBOutput(os.Stdout, tt.args.rgb, tt.args.message)
		})
	}
}
