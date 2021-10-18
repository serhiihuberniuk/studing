package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"studing/tasks/task_14/server/models"
)

type Handler struct {
	conn    net.Conn
	scanner scanner
}

func NewSender(conn net.Conn, s scanner) *Handler {
	return &Handler{
		conn:    conn,
		scanner: s,
	}
}

type scanner interface {
	Scan(ctx context.Context) (string, error)
}

func (h *Handler) Validate(ctx context.Context) error {
	str, err := h.scanner.Scan(ctx)
	if err != nil {
		return fmt.Errorf("error while scanning: %w", err)
	}

	requestMessage := &models.RequestMessage{
		Text: str,
	}

	message, err := json.Marshal(requestMessage)
	if err != nil {
		return fmt.Errorf("error while marshalling request: %w", err)
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
	if err = json.Unmarshal(data, responseMessage); err != nil {
		return fmt.Errorf("error while unmarshalling responce: %w", err)
	}

	if !responseMessage.Success {
		fmt.Println("Validation failed:", responseMessage.Text)

		return nil
	}

	fmt.Println("String is OK")
	return nil
}
