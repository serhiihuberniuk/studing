package handler

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"time"

	"studing/tasks/task_14/server/models"
)

type Handler struct {
	conn    net.Conn
	scanner scanner
}

func New(conn net.Conn, s scanner) *Handler {
	return &Handler{
		conn:    conn,
		scanner: s,
	}
}

type scanner interface {
	Scan(ctx context.Context) (string, error)
}

func (h *Handler) Handle(ctx context.Context) error {
	str, err := h.scanner.Scan(ctx)
	if err != nil {
		return fmt.Errorf("error while scanning: %w", err)
	}

	requestMessage := &models.RequestMessage{
		Text: str,
	}

	message, err := requestMessage.ToBytes()
	if err != nil {
		return fmt.Errorf("error while transforming message to bytes: %w", err)
	}

	if err = h.conn.SetWriteDeadline(time.Now().Add(time.Second * 1)); err != nil {
		return fmt.Errorf("error while setting timeout: %w", err)
	}

	_, err = h.conn.Write(message)
	if err != nil {
		return fmt.Errorf("error while writing request to server: %w", err)
	}

	if err = h.conn.SetReadDeadline(time.Now().Add(time.Second * 1)); err != nil {
		return fmt.Errorf("error while setting timeout: %w", err)
	}

	data, err := bufio.NewReader(h.conn).ReadBytes('}')
	if err != nil {
		return fmt.Errorf("error while reading responce: %w", err)
	}

	responseMessage := &models.ResponseMessage{}
	if err = responseMessage.FromBytes(data); err != nil {
		return fmt.Errorf("error while transforming message from bytes: %w", err)
	}

	if !responseMessage.Success {
		fmt.Println("Validation failed:", responseMessage.Text)

		return nil
	}

	fmt.Println("String is OK")
	return nil
}
