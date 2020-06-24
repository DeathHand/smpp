package smpp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

// Reader reads pdu buffer
type Reader struct {
	r *bufio.Reader
	p *[]byte
}

// NewReader reader constructor
func NewReader(r *bufio.Reader) *Reader {
	p := make([]byte, 4)
	return &Reader{
		r: r,
		p: &p,
	}
}

// Reader read pdu to buffer
func (r *Reader) Read(buffer *bytes.Buffer) error {
	n, err := r.r.Read(*r.p)
	if err != nil {
		return err
	}
	if n < len(*r.p) {
		return io.ErrUnexpectedEOF
	}
	commandLength := binary.BigEndian.Uint32(*r.p)
	b := make([]byte, commandLength-4)
	n, err = r.r.Read(b)
	if err != nil {
		return err
	}
	if n < len(b) {
		return io.ErrUnexpectedEOF
	}
	n, err = buffer.Write(*r.p)
	if err != nil {
		return err
	}
	if n < len(*r.p) {
		return io.ErrShortWrite
	}
	n, err = buffer.Write(b)
	if err != nil {
		return err
	}
	if n < len(b) {
		return io.ErrShortWrite
	}
	return nil
}
