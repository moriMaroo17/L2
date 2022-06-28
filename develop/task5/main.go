package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

// init flags variables as global
var (
	count, ignore, invert, fix, numbers bool
	after, before, context              int
)

// cmdFlags contains command-line flags
type cmdFlags struct {
	after   int
	before  int
	context int
	count   bool
	ignore  bool
	invert  bool
	fix     bool
	numbers bool
}

// cmdFlagsConstructor is the constructor of cmdFlags (depends on package global variables)
func cmdFlagsConstructor() cmdFlags {
	return cmdFlags{after, before, context, count, ignore, invert, fix, numbers}
}

// cmdParams contains all command-line parameters for package
type cmdParams struct {
	fileName string
	pattern  string
	cmdFlags
}

// cmdParamsConstructor is the constructor of cmdParams
func cmdParamsConstructor(fileName, pattern string, cmdFlags cmdFlags) cmdParams {
	return cmdParams{fileName, pattern, cmdFlags}
}

// execute is the function for find intersection of pattern and rows in the given file
func (c *cmdParams) execute() ([]string, map[int]struct{}, error) {

	var patternAsRegExp *regexp.Regexp

	rows, err := readFile(c.fileName)
	if err != nil {
		return nil, nil, err
	}

	if c.ignore {
		patternAsRegExp, err = regexp.Compile(strings.ToLower(c.pattern))
	} else {
		patternAsRegExp, err = regexp.Compile(c.pattern)
	}
	if err != nil {
		return nil, nil, err
	}

	resultIdxs := make(map[int]struct{})

	for idx, row := range rows {
		if (c.fix && strings.Contains(row, c.pattern)) || patternAsRegExp.MatchString(row) {
			resultIdxs[idx] = struct{}{}
		}
	}
	return rows, resultIdxs, nil
}

// printOut is a helper function for printning out result depends on given flags
func (c *cmdParams) printOut(allRows []string, matchesIdxs map[int]struct{}) {
	upBorder := int(math.Max(float64(before), float64(context)))
	downBorder := int(math.Max(float64(after), float64(context)))

	printMethod := func(idx int, row string) {
		if c.numbers {
			fmt.Printf("%d: %s\n", idx, row)
		} else {
			fmt.Println(row)
		}
	}

	switch {
	case c.count:
		fmt.Println(len(matchesIdxs))
	case c.invert:
		for idx, row := range allRows {
			if _, ok := matchesIdxs[idx]; !ok {
				printMethod(idx, row)
			}
		}
	default:
		var previousIdx, postIdx int

		for idx, row := range allRows {
			if _, ok := matchesIdxs[idx]; !ok {
				continue
			}
			previousIdx = idx - upBorder
			if previousIdx < 0 {
				previousIdx = 0
			}
			postIdx = idx + downBorder
			if postIdx > len(allRows)-1 {
				postIdx = len(allRows) - 1
			}
			if previousIdx == postIdx {
				printMethod(idx, row)
				fmt.Println("")
				continue
			}
			for idxCtxRow, ctxRow := range allRows[previousIdx : postIdx+1] {
				printMethod(idxCtxRow, ctxRow)
			}
			fmt.Println("")
		}
	}
}

func main() {
	// init the command-line flags
	flag.IntVar(&after, "A", 0, "output n string after the coincidence")
	flag.IntVar(&before, "B", 0, "output n string before the coincidence")
	flag.IntVar(&context, "C", 0, "output n string around the coincidence")
	flag.BoolVar(&count, "c", false, "output number of strings")
	flag.BoolVar(&ignore, "i", false, "ignore case")
	flag.BoolVar(&invert, "v", false, "invert (instead of match, exclude)")
	flag.BoolVar(&fix, "F", false, "exact string match, not a pattern")
	flag.BoolVar(&numbers, "n", false, "output lines numbers")

	flag.Parse()

	if len(flag.Args()) < 2 {
		panic("too few arguments")
	}

	cmdParamsAsStruct := cmdParamsConstructor(flag.Args()[0], flag.Args()[1], cmdFlagsConstructor())

	allRows, matchesIdxs, err := cmdParamsAsStruct.execute()
	if err != nil {
		panic(err)
	}

	cmdParamsAsStruct.printOut(allRows, matchesIdxs)
}

// readFile reads a file and returns a slice of strings
func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rowsArr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowsArr = append(rowsArr, scanner.Text())
	}
	return rowsArr, nil
}
