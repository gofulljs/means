package color

import (
	"bufio"
	"errors"
	"golang.org/x/xerrors"
	"io"
	"strings"
)

// 7 level colors
// logoText: logoText must only 7 lines
func ColorLogo(logoText string, writer io.Writer) (err error) {
	r := bufio.NewReader(strings.NewReader(logoText))

	// 1 line
	line, _, err := r.ReadLine()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	Color(WithOutput(writer), WithHex(0xff0000)).Outputn(string(line))

	// 2~6
	for i := 2; i <= 6; i++ {
		line, _, err = r.ReadLine()
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		n1 := len(line) / 3
		n2 := n1 * 2
		switch i {
		case 2:
			Color(WithOutput(writer), WithHex(0xff0000)).Output(string(line[:n1]))
			Color(WithOutput(writer), WithHex(0xff3b00)).Output(string(line[n1:n2]))
			Color(WithOutput(writer), WithHex(0xff7500)).Outputn(string(line[n2:]))
		case 3:
			Color(WithOutput(writer), WithHex(0xff7800)).Output(string(line[:n1]))
			Color(WithOutput(writer), WithHex(0xfd7b00)).Output(string(line[n1:n2]))
			Color(WithOutput(writer), WithHex(0xffad00)).Outputn(string(line[n2:]))
		case 4:
			Color(WithOutput(writer), WithHex(0xfeda00)).Output(string(line[:n1]))
			Color(WithOutput(writer), WithHex(0xddfd00)).Output(string(line[n1:n2]))
			Color(WithOutput(writer), WithHex(0x93ff00)).Outputn(string(line[n2:]))
		case 5:
			Color(WithOutput(writer), WithHex(0x80ff00)).Output(string(line[:n1]))
			Color(WithOutput(writer), WithHex(0x1aff00)).Output(string(line[n1:n2]))
			Color(WithOutput(writer), WithHex(0x00ff2e)).Outputn(string(line[n2:]))
		case 6:
			Color(WithOutput(writer), WithHex(0x00ff3b)).Output(string(line[:n1]))
			Color(WithOutput(writer), WithHex(0x00ffb1)).Output(string(line[n1:n2]))
			Color(WithOutput(writer), WithHex(0x00f2f9)).Outputn(string(line[n2:]))
		}
	}

	// 7 else

	for {
		line, _, err = r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return xerrors.Errorf("%w", err)
		}
		Color(WithOutput(writer), WithHex(0x00e0f9)).Outputn(string(line))
	}

	return nil
}
