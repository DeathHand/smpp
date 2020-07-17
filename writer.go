package smpp

import (
	"bufio"
	"bytes"
	"io"
)

// Writer writes pdu buffer
type Writer struct {
	w *bufio.Writer
}

// NewWriter writer constructor
func NewWriter(w *bufio.Writer) *Writer {
	return &Writer{w: w}
}

// Write write pdu
func (w *Writer) Write(buffer *bytes.Buffer) error {
	l := buffer.Len()
	n, err := w.w.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	if n < l {
		return io.ErrShortWrite
	}
	return w.w.Flush()
}
