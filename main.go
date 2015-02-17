// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/push"
	"github.com/gdamore/mangos/transport/ipc"
	"github.com/gdamore/mangos/transport/tcp"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/sensor/l3gd20"

	_ "github.com/kidoman/embd/host/all"
)

const url = "tcp://127.0.0.1:40899"

func main() {
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	var sock mangos.Socket
	var err error
	if sock, err = push.NewSocket(); err != nil {
		panic(err)
	}
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Dial(url); err != nil {
		panic(err)
	}
	defer sock.Close()

	bus := embd.NewI2CBus(1)

	gyro := l3gd20.New(bus, l3gd20.R250DPS)
	defer gyro.Close()

	gyro.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	orientations, err := gyro.Orientations()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case orientation := <-orientations:
			msg := fmt.Sprintf("%.f,%.f,%.f", orientation.X, orientation.Y, orientation.Z)
			fmt.Printf("sending %q\n", msg)
			go sock.Send([]byte(msg))
		case <-quit:
			return
		}
	}
}
