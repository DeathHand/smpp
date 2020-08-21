package smpp

import (
	"bytes"
	"io"
	"net"
)

// Writer writes pdu buffer
type Writer struct {
	conn *net.TCPConn
}

// NewWriter writer constructor
func NewWriter(conn *net.TCPConn) *Writer {
	return &Writer{conn: conn}
}

// Write write pdu
func (w *Writer) Write(buffer *bytes.Buffer) error {
	l := buffer.Len()
	n, err := w.conn.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	if n < l {
		return io.ErrShortWrite
	}
	return nil
}
