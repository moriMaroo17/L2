package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func printTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", time)
}

func main() {
	printTime()
}
