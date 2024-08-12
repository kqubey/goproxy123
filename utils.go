package main

import (
	"bytes"
	"sync"
)

var ReusableBuffer *bytes.Buffer
var ReusableBufferMux sync.Mutex

func init() {
	ReusableBuffer = bytes.NewBuffer(make([]byte, 0, 256))
}
