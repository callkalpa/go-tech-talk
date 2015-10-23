package main
import "os"

type Reader interface {
	Read(f os.File) error
}

type Writer interface {
	Write(f os.File) error
}

type ReadWriter interface {
	Reader
	Writer
}