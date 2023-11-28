package color

type StyleFn func() string

func BoldFormat() string {
	return "1"
}

func OnlyBoldFormat() (left, right string) {
	left = "\x1b[1m"
	right = "\x1b[0m"

	return left, right
}

func ItalicFormat() string {
	return "3"
}

func OnlyItalicFormat() (left, right string) {
	left = "\x1b[3m"
	right = "\x1b[0m"

	return left, right
}
