package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"studing/tasks/task_14/server/handler"
	"studing/tasks/task_14/server/service"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("Server is ready to accept connection:")

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Connection with %v is accepted from\n", conn.RemoteAddr())

			s := service.NewService()
			h := handler.NewHandler(conn, s)

			go func() {
				h.CheckStringFromRequest()
			}()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
