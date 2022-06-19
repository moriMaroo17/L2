package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func printTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", time)
}

func main() {
	printTime()
}
