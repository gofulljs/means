package encode_utils

import (
	"io"
	"strings"
	"testing"
)

func TestMd5Str(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base",
			data: "hello,gopher",
			want: "44bb857707d28a7f55e27c2cffe5f2ba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Str(tt.data); got != tt.want {
				t.Errorf("Md5Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5Reader(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   string
	}{
		{
			name: "base",
			reader: func() io.Reader {
				return strings.NewReader("hello,gopher")
			}(),
			want: "44bb857707d28a7f55e27c2cffe5f2ba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Reader(tt.reader); got != tt.want {
				t.Errorf("Md5Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash256Str(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base",
			data: "hello,gopher",
			want: "201c197ac3a88f77c25964e54ed3deec6a1a542cadc0b18a53cd54e4a1bd4ef6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash256Str(tt.data); got != tt.want {
				t.Errorf("Hash256Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash256Reader(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   string
	}{
		{
			name: "base",
			reader: func() io.Reader {
				return strings.NewReader("hello,gopher")
			}(),
			want: "201c197ac3a88f77c25964e54ed3deec6a1a542cadc0b18a53cd54e4a1bd4ef6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash256Reader(tt.reader); got != tt.want {
				t.Errorf("Hash256Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash224Str(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base",
			data: "hello,gopher",
			want: "ab6d19e4b7f047983c72e79ea16763ce04eb62605a3e8caf56150fff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash224Str(tt.data); got != tt.want {
				t.Errorf("Hash256Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash224Reader(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   string
	}{
		{
			name: "base",
			reader: func() io.Reader {
				return strings.NewReader("hello,gopher")
			}(),
			want: "ab6d19e4b7f047983c72e79ea16763ce04eb62605a3e8caf56150fff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash224Reader(tt.reader); got != tt.want {
				t.Errorf("Hash256Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash512Str(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base",
			data: "hello,gopher",
			want: "328e82ba663aa0f5ac2c29a01b236d47fbfd4a81b5db514e9c591db45518d49fc91330defe9b628c7fec0f5b46aa5809cd6fcc23642d849838b8a4d08152efbe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash512Str(tt.data); got != tt.want {
				t.Errorf("Hash256Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash512Reader(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   string
	}{
		{
			name: "base",
			reader: func() io.Reader {
				return strings.NewReader("hello,gopher")
			}(),
			want: "328e82ba663aa0f5ac2c29a01b236d47fbfd4a81b5db514e9c591db45518d49fc91330defe9b628c7fec0f5b46aa5809cd6fcc23642d849838b8a4d08152efbe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash512Reader(tt.reader); got != tt.want {
				t.Errorf("Hash256Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash384Str(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "base",
			data: "hello,gopher",
			want: "f80222ea8805b0e1bd8cc6169f7fb09aa888a8c991957178e834407c90e65ed8f19f097659882a8585e56d7438929eea",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash384Str(tt.data); got != tt.want {
				t.Errorf("Hash256Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash384Reader(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   string
	}{
		{
			name: "base",
			reader: func() io.Reader {
				return strings.NewReader("hello,gopher")
			}(),
			want: "f80222ea8805b0e1bd8cc6169f7fb09aa888a8c991957178e834407c90e65ed8f19f097659882a8585e56d7438929eea",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash384Reader(tt.reader); got != tt.want {
				t.Errorf("Hash256Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}
