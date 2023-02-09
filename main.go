package main

import (
	"awesomeProject/part_2/chapter26"
	"bytes"
	"errors"
	"sync"
)

func test01() (int, error) {
	return 1, errors.New("error occurs")
}

type Reader interface {
	Read([]byte) (int, error)
}

type custom struct {
	name string
}

func ReadFull(r Reader, buf []byte) (int, error) {
	var (
		n   int
		err error
	)
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return n, err
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	chapter26.StartPrinter()
}
