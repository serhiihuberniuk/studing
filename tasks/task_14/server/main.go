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

	"studing/tasks/task_14/server/handler"
	"studing/tasks/task_14/server/validator"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("Server is ready to accept connection:")

	s := validator.New()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("Connection with %v is accepted\n", conn.RemoteAddr())

				h := handler.NewHandler(conn, s)

				go func() {
					defer conn.Close()
					for {
						select {
						case <-ctx.Done():
							return
						default:
							if err := h.Validate(ctx); err != nil {
								if errors.Is(err, io.EOF) {
									log.Printf("connection %v with is closed by client\n", conn.RemoteAddr())

									return
								}
								if errors.Is(err, os.ErrDeadlineExceeded) {
									log.Println(err)

									return
								}
								log.Println(err)
							}
						}
					}

				}()
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
}
