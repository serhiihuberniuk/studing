package sender

import (
	"fmt"
	"net"
)

type Sender struct {
	conn    net.Conn
	scanner scanner
}

func NewSender(conn net.Conn, s scanner) *Sender {
	return &Sender{
		conn:    conn,
		scanner: s,
	}
}

type scanner interface {
	ScanTerminal() (string, error)
}

func (s *Sender) SendStringToCheck() error {
	str, err := s.scanner.ScanTerminal()
	if err != nil {
		return fmt.Errorf("error while scanning terminal")
	}

	_, err = s.conn.Write([]byte(str + "\n"))
	if err != nil {
		return fmt.Errorf("error while writing request to server: %w", err)
	}

	return nil
}
