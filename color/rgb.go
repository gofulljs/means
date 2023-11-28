package color

import (
	"fmt"
	"io"
)

// ref: https://juejin.cn/post/6920241597846126599
type RGB struct {
	r, g, b uint8
}

func RGBFormat(rgb RGB) string {
	return fmt.Sprintf("38;2;%v;%v;%v", rgb.r, rgb.g, rgb.b)
}

// HexFormat hex: 0x1a2b3c
func HexFormat(hex uint32) string {
	r := uint8(hex >> 16)
	g := uint8(hex >> 8)
	b := uint8(hex)
	return RGBFormat(RGB{r, g, b})
}

func OnlyRGBForamt(rgb RGB) (left, right string) {
	left = fmt.Sprintf("\x1b[3;1;38;2;%v;%v;%vm", rgb.r, rgb.g, rgb.b)
	right = "\x1b[0m"
	return left, right
}

func RGBOutput(writer io.Writer, rgb RGB, message string) {
	left, right := OnlyRGBForamt(rgb)
	fmt.Fprintf(writer, "%s%s%s", left, message, right)
}
