package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func printTime() error {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", time)
	return nil
}

func main() {
	if err := printTime(); err != nil {
		l := log.New(os.Stderr, "", 1)
		l.Fatal(err)
	}
}
