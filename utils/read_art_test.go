package net_cat

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type MockConn struct {
	buf       strings.Builder
	readData  []byte
	readIndex int
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	if m.readIndex >= len(m.readData) {
		return 0, fmt.Errorf("EOF")
	}
	n = copy(b, m.readData[m.readIndex:])
	m.readIndex += n
	return n, nil
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.buf.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr                { return nil }
func (m *MockConn) RemoteAddr() net.Addr               { return nil }
func (m *MockConn) SetDeadline(t time.Time) error      { return nil }
func (m *MockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockConn) SetWriteDeadline(t time.Time) error { return nil }
