package main

// START READOMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

// END READOMIT

// START WRITEOMIT
type Writer interface {
	Write(p []byte) (n int, err error)
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// END WRITEOMIT
