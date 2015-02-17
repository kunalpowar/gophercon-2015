// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pair"
	"github.com/gdamore/mangos/transport/tcp"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/sensor/l3gd20"

	_ "github.com/kidoman/embd/host/all"
)

const url = "tcp://192.168.2.1:40899"

func main() {
	sock, err := pair.NewSocket()
	if err != nil {
		panic(err)
	}

	sock.AddTransport(tcp.NewTransport())
	if err := sock.Dial(url); err != nil {
		fmt.Errorf("could not dial %v", url)
		panic(err)
	}

	sock.SetOption(mangos.OptionRecvDeadline, 5*time.Second)
	fmt.Println("waiting for a message from app")
	msg, err := sock.Recv()
	if err != nil {
		fmt.Errorf("could not receive from %v", url)
		panic(err)
	}
	fmt.Println("received msg %v to start sending sensor data", msg)

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	gyro := l3gd20.New(bus, l3gd20.R2000DPS)
	defer gyro.Close()
	gyro.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	orientations, err := gyro.Orientations()
	if err != nil {
		panic(err)
	}

	sendUpdate := time.NewTicker(100 * time.Millisecond)
	defer sendUpdate.Stop()

	var orientation l3gd20.Orientation

	for {
		select {
		case orientation = <-orientations:
		case <-sendUpdate.C:
			msg := fmt.Sprintf("%.f,%.f,%.f", orientation.X, orientation.Y, orientation.Z)
			fmt.Printf("sending %q\n", msg)
			sock.Send([]byte(msg))
		case <-quit:
			return
		}
	}
}
