package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()

		args := strings.Split(command, " ")

		switch args[0] {
		case `\quit`:
			return
		case "cd":
			if len(args) == 1 {
				home, _ := os.UserHomeDir()
				os.Chdir(home)
			} else {
				os.Chdir(args[1])
			}
		default:
			execute(args...)
		}
	}
}

func execute(commandAndArgs ...string) {
	command := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
