package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"studing/tasks/task_14/client/handler"
	"studing/tasks/task_14/client/scanner"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ts := scanner.NewTerminalScanner()
	h := handler.NewSender(conn, ts)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := h.Validate(ctx); err != nil {
					if errors.Is(err, io.EOF) {
						log.Printf("connection is closed by server: %v\n", err)

						return
					}

					log.Println(err)
				}
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	cancel()
}
