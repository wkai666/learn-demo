package main

import (
	"io"
	"time"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader(p []byte) (n int, err error)
	Writer(p []byte) (n int, err error)
}

type ReadWritCloser interface {
	Reader
	Writer
	Closer
}

var any interface{}

// any = true
// any = 12.234
// any = "hello"
// any = map[string]int{"one":1}
//any = new(bytes.Buffer)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}

func main() {

}
