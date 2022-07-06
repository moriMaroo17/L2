package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
		case "pwd":
			path, _ := filepath.Abs(".")
			fmt.Println(path)
		case "echo":
			if len(args) > 1 {
				for _, arg := range args {
					fmt.Printf("%s ", arg)
				}
			}
			fmt.Println()
		case "kill":
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
			}

			proc, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println(err)
			}

			err = proc.Kill()
			if err != nil {
				fmt.Println(err)
			}
		case "ps":
			pids, err := ioutil.ReadDir("/proc")
			if err != nil {
				fmt.Println(err)
			}

			for _, pid := range pids {
				if pidAsInt, err := strconv.Atoi(pid.Name()); err == nil {
					fmt.Println(pidAsInt)
				}
			}
		default:
			fmt.Println("unsupportable command")
		}
	}
}

func execBash(commandAndArgs string) {
	command := exec.Command("bash", "-c", commandAndArgs)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
