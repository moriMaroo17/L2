package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	host    string
	port    string
	timeout time.Duration
}

var telnetConfig config

func main() {
	fullAddress := net.JoinHostPort(telnetConfig.host, telnetConfig.port)
	// using net.Dial here because it returns an net.Conn, which implements io.Reader interface
	connection, err := net.Dial("tcp", fullAddress)
	if err != nil {
		log.Fatal(err)
	}
	// init channels
	sigChan := make(chan os.Signal)
	errChan := make(chan error)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// init timer
	timer := time.NewTimer(telnetConfig.timeout)
	// close all before return
	defer timer.Stop()
	defer close(sigChan)
	defer close(errChan)
	defer connection.Close()

	go send(connection, sigChan, errChan)
	go read(connection, sigChan, errChan)
	// listne channels
	for {
		select {
		case <-errChan:
			fmt.Println(err)
		case <-sigChan:
			return
		case <-timer.C:
			return
		}
	}
}

// parse flags and init connection config
func init() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")

	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("too few arguments")
	}

	telnetConfig = config{host: flag.Args()[0], port: flag.Args()[1], timeout: timeout}
}

// send is used to send data from sdin to the socket
func send(connection net.Conn, sigChan chan os.Signal, errChan chan error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			sigChan <- syscall.SIGQUIT
			return
		}
		if err != nil {
			errChan <- err
		}
		_, err = connection.Write(line)
		if err != nil {
			errChan <- err
		}
	}

}

// read is used to read data from the socket
func read(connection net.Conn, sigChan chan os.Signal, errChan chan error) {
	reader := bufio.NewReader(connection)

	for {
		message, err := reader.ReadBytes('\n')
		if err == io.EOF {
			sigChan <- syscall.SIGQUIT
			return
		}
		if err != nil {
			errChan <- err
		}
		fmt.Print("got message:", string(message))
	}
}
