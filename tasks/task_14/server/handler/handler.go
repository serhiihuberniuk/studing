package handler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Handler struct {
	conn    net.Conn
	service service
}

func NewHandler(conn net.Conn, s service) *Handler {
	return &Handler{
		conn:    conn,
		service: s,
	}
}

type service interface {
	CheckString(str string) (string, error)
}

func (h *Handler) CheckStringFromRequest() {
	defer h.conn.Close()

	for {
		requestMessage, err := bufio.NewReader(h.conn).ReadString('\n')
		switch err {
		case nil:
			responseMessage, err := h.service.CheckString(requestMessage)
			if err != nil {

				log.Println(fmt.Errorf("error while checking string: %w", err))
				h.writeResponse("Internal server error")
				continue
			}

			h.writeResponse(responseMessage)

		case io.EOF:
			log.Printf("connection with %v closed by client: ", h.conn.RemoteAddr())

			return

		default:
			log.Println(fmt.Errorf("error while reading message: %w", err))
			h.writeResponse("Internal server error")
		}
	}
}

func (h *Handler) writeResponse(message string) {
	if _, err := h.conn.Write([]byte(message + "\n")); err != nil {
		log.Println(fmt.Errorf("error while writing response: %w", err))
	}

}
