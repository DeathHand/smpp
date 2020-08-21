package smpp

import (
	"bytes"
	"io"
	"net"
)

// Reader reads pdu buffer
type Reader struct {
	conn *net.TCPConn
}

// NewReader reader constructor
func NewReader(conn *net.TCPConn) *Reader {
	return &Reader{conn: conn}
}

// Reader read pdu to buffer
func (r *Reader) Read(buffer *bytes.Buffer) error {
	p := make([]byte, 4)
	n, err := io.ReadFull(r.conn, p)
	if err != nil {
		return err
	}
	if n < len(p) {
		return io.ErrUnexpectedEOF
	}
	commandLength := int32(p[3]) | int32(p[2])<<8 | int32(p[1])<<16 | int32(p[0])<<24
	b := make([]byte, commandLength-4)
	n, err = io.ReadFull(r.conn, b)
	if err != nil {
		return err
	}
	if n < len(b) {
		return io.ErrUnexpectedEOF
	}
	n, err = buffer.Write(p)
	if err != nil {
		return err
	}
	if n < len(p) {
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
