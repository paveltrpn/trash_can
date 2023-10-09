package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigchnl := make(chan os.Signal, 1)
	caught := make(chan os.Signal, 1)

	signal.Notify(sigchnl)

	go func() {
		for {
			s := <-sigchnl
			caught <- s
		}
	}()

	for i := 0; ; i++ {
		select {
		case signal := <-caught:
			if signal == syscall.SIGINT {
				fmt.Printf("Got CTRL+C signal, closing!\n")
				os.Exit(0)
			}
		default:
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(i * i)
		}
	}
}
