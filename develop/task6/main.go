package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cutter is a struct that holds a list of parameters
type cutter struct {
	delimeter string
	fields    map[int]struct{}
	separated bool
}

// newCutter is a constructor for cutter struct
func newCutter(delimeter string, fields string, separated bool) cutter {
	return cutter{delimeter: delimeter, fields: parseInterval(fields), separated: separated}
}

// execute is a function that takes a string and printing out the result depending on the parameters
func (c *cutter) execute(input string) {
	splittedString := strings.Split(input, c.delimeter)

	if c.separated && len(splittedString) < 2 {
		return
	}

	if len(c.fields) == 0 {
		for _, field := range splittedString {
			fmt.Println(field)
		}
		return
	}

	for idx, field := range splittedString {
		if _, ok := c.fields[idx]; ok {
			fmt.Println(field)
		}
	}
}

func main() {
	delimeter := flag.String("d", " ", "set delimeter")
	fields := flag.String("f", "", "set fields")
	separated := flag.Bool("s", false, "output only separated fields")

	flag.Parse()

	cutter := newCutter(*delimeter, *fields, *separated)

	fmt.Println(cutter)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cutter.execute(scanner.Text())
	}
}

// parseInterval is a function that parses an interval frm string and returns int map
func parseInterval(interval string) map[int]struct{} {
	result := make(map[int]struct{})

	if interval != "" {
		intervalParts := strings.Split(interval, ",")

		for _, part := range intervalParts {
			idxs := strings.Split(part, "-")
			for _, idx := range idxs {
				intIdx, err := strconv.Atoi(idx)
				if err != nil {
					fmt.Println(err)
				}
				result[intIdx-1] = struct{}{}
			}
		}
	}

	return result
}
