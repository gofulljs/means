package encode_utils

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func Md5Str(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func Md5Reader(reader io.Reader) string {
	hasher := md5.New()

	io.Copy(hasher, reader)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Hash256Str(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func Hash256Reader(reader io.Reader) string {
	hasher := sha256.New()

	io.Copy(hasher, reader)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Hash224Str(data string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(data)))
}

func Hash224Reader(reader io.Reader) string {
	hasher := sha256.New224()

	io.Copy(hasher, reader)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Hash512Str(data string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
}

func Hash512Reader(reader io.Reader) string {
	hasher := sha512.New()

	io.Copy(hasher, reader)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func Hash384Str(data string) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(data)))
}

func Hash384Reader(reader io.Reader) string {
	hasher := sha512.New384()

	io.Copy(hasher, reader)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}
