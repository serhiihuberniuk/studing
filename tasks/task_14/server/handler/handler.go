package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"studing/tasks/task_14/server/models"
)

type Handler struct {
	conn      net.Conn
	validator validator
}

func NewHandler(conn net.Conn, v validator) *Handler {
	return &Handler{
		conn:      conn,
		validator: v,
	}
}

type validator interface {
	Validate(ctx context.Context, str string) error
}

func (h *Handler) Validate(ctx context.Context) error {

	if err := h.conn.SetReadDeadline(time.Now().Add(time.Minute)); err != nil {
		return fmt.Errorf("error while setting timeout: %w", err)
	}

	data, err := bufio.NewReader(h.conn).ReadBytes('}')
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			if err := h.writeResponse(false, "server time-out"); err != nil {
				return fmt.Errorf("error while writing response: %w", err)
			}
		}

		if err := h.writeResponse(false, "internal server error"); err != nil {
			return fmt.Errorf("error while writing response: %w", err)
		}

		return fmt.Errorf("error while reading message: %w", err)
	}

	requestMessage := &models.RequestMessage{}
	if err := json.Unmarshal(data, requestMessage); err != nil {
		if err := h.writeResponse(false, "invalid request"); err != nil {
			return fmt.Errorf("error while writing response: %w", err)
		}

		return fmt.Errorf("error while unmarshalling request: %w", err)
	}

	if err := h.validator.Validate(ctx, requestMessage.Text); err != nil {
		if err := h.writeResponse(false, err.Error()); err != nil {
			return fmt.Errorf("error while writing response: %w", err)
		}

		return fmt.Errorf("error while checking string: %w", err)
	}

	if err := h.writeResponse(true, "OK"); err != nil {
		return fmt.Errorf("error while writing response: %w", err)
	}

	return nil
}

func (h *Handler) writeResponse(success bool, text string) error {
	responseMessage := &models.ResponseMessage{
		Success: success,
		Text:    text,
	}

	message, err := json.Marshal(responseMessage)
	if err != nil {
		return fmt.Errorf("error while marshlling message: %w", err)
	}

	if err = h.conn.SetWriteDeadline(time.Now().Add(time.Second * 1)); err != nil {
		return fmt.Errorf("error while setting timeout: %w", err)
	}

	if _, err := h.conn.Write(message); err != nil {
		return fmt.Errorf("error while writing response: %w", err)
	}

	return nil
}
