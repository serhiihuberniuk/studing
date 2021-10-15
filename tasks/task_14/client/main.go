package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"studing/tasks/task_14/client/getter"
	"studing/tasks/task_14/client/scanner"
	"studing/tasks/task_14/client/sender"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	s := scanner.NewScanner()
	sdr := sender.NewSender(conn, s)
	g := getter.NewGetter(conn)

	go func() {
		for {
			if err := sdr.SendStringToCheck(); err != nil {
				log.Printf("error while sending request: %v\n", err)
			}
			if err := g.GetResponse(); err != nil {
				log.Printf("error while getting response: %v\n", err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}
