package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func addZeros(ctx context.Context, in chan string, out chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			s := <-in
			s = s + "000"
			out <- s
		}

	}
}
func convertToInt(ctx context.Context, in chan string, out chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			s := <-in
			i, err := strconv.Atoi(s)
			if err == nil {
				out <- i
			}
		}
	}
}

func convertToString(ctx context.Context, in chan int, out chan string) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			i := <-in
			s := strconv.Itoa(i)
			out <- s
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	setupGracefulShutdown(cancel)

	stringArray := []string{"90", "23", "30", "123", "34", "4452", "23", "123", "000", "12", "34"}
	ch := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan int)
	ch3 := make(chan string)

	go addZeros(ctx, ch, ch1)
	go convertToInt(ctx, ch1, ch2)
	go convertToString(ctx, ch2, ch3)

	go func() {
		for _, v := range stringArray {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- v
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(<-ch3)
			}
		}
	}()
	<-ctx.Done()
	close(ch)
	close(ch1)
	close(ch2)
	close(ch3)
}
func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		stop()
	}()
}
