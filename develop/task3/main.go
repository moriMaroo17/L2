package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	n, r, u, m, b, c, h bool
	k                   int
)

func main() {
	flag.BoolVar(&n, "n", false, "sort by numerical value")
	flag.BoolVar(&h, "h", false, "sort by numerical value considering suffixes")
	flag.BoolVar(&m, "M", false, "sort by month")
	flag.BoolVar(&c, "c", false, "check if data is sorted")

	flag.BoolVar(&r, "r", false, "reverse sort order")
	flag.BoolVar(&b, "b", false, "ignore bask whitespaces")
	flag.BoolVar(&u, "u", false, "don't output duplicate lines")
	flag.IntVar(&k, "k", 0, "column for sort")

	flag.Parse()

	inputFile := flag.Args()[0]
	outputFile := flag.Args()[1]

	if inputFile == "" || outputFile == "" {
		log.Fatal("U mast specify input and output file as two first args")
	}

	if k > 0 {
		k--
	} else {
		k = 0
	}

	fileAsMatrix := readFileToMatrix(inputFile)

	var sortFunction func(firstElemIdx, secondElemIdx int) bool

	switch true {
	case n:
		sortFunction = func(firstElemIdx, secondElemIdx int) bool {
			firstElemAsFloat, _ := strconv.ParseFloat(getMatrixElement(fileAsMatrix, firstElemIdx, k), 64)
			secondElemAsFloat, _ := strconv.ParseFloat(getMatrixElement(fileAsMatrix, secondElemIdx, k), 64)
			if r {
				return firstElemAsFloat > secondElemAsFloat
			}
			return firstElemAsFloat < secondElemAsFloat
		}
	case m:
		sortFunction = func(firstElemIdx, secondElemIdx int) bool {
			firstElemAsMonth := converStringToTime(getMatrixElement(fileAsMatrix, firstElemIdx, k))
			secondElemAsMonth := converStringToTime(getMatrixElement(fileAsMatrix, secondElemIdx, k))
			if r {
				return firstElemAsMonth.After(secondElemAsMonth)
			}
			return firstElemAsMonth.Before(secondElemAsMonth)
		}
	case h:
		sortFunction = func(firstElemIdx, secondElemIdx int) bool {
			firstElemAsFloat := convertSuffixes(getMatrixElement(fileAsMatrix, firstElemIdx, k))
			secondElemAsFloat := convertSuffixes(getMatrixElement(fileAsMatrix, secondElemIdx, k))
			if r {
				return firstElemAsFloat > secondElemAsFloat
			}
			return firstElemAsFloat < secondElemAsFloat
		}
	default:
		sortFunction = func(firstElemIdx, secondElemIdx int) bool {
			firstElem := getMatrixElement(fileAsMatrix, firstElemIdx, k)
			secondElem := getMatrixElement(fileAsMatrix, secondElemIdx, k)
			if r {
				return firstElem > secondElem
			}
			return firstElem < secondElem
		}
	}

	if c {
		isSorted := sort.SliceIsSorted(fileAsMatrix, sortFunction)
		fmt.Printf("File sorted: %v\n", isSorted)
		return
	}

	sort.Slice(fileAsMatrix, sortFunction)

	if r {

	}

	writeResultIntoFile(fileAsMatrix, outputFile)
}

func readFileToMatrix(inputFile string) [][]string {

	var result [][]string

	openedFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer openedFile.Close()

	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		if b {
			result = append(result, strings.Split(strings.TrimRight(scanner.Text(), "\t \n"), " "))
		}
		result = append(result, strings.Split(scanner.Text(), " "))
	}
	return result
}

func getMatrixElement(matrix [][]string, stringIdx, columnIdx int) string {
	if columnIdx < len(matrix[stringIdx]) {
		return matrix[stringIdx][columnIdx]
	}
	return ""
}

func converStringToTime(month string) time.Time {
	if tm, err := time.Parse("Jan", month); err == nil {
		return tm
	}
	if tm, err := time.Parse("January", month); err == nil {
		return tm
	}
	if tm, err := time.Parse("01", month); err == nil {
		return tm
	}
	if tm, err := time.Parse("1", month); err == nil {
		return tm
	}
	return time.Time{}
}

func convertSuffixes(s string) float64 {
	sAsRunes := []rune(s)
	suffix := sAsRunes[len(sAsRunes)-1]
	number := string(sAsRunes[:len(sAsRunes)-1])
	result, _ := strconv.ParseFloat(number, 64)
	switch suffix {
	case 'd':
		return result * 10
	case 'h':
		return result * 100
	case 'k':
		return result * 1000
	case 'M':
		return result * 1000000
	case 'G':
		return result * 1000000000
	default:
		return result
	}
}

func writeResultIntoFile(matrix [][]string, output string) {
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linesToWrite := make([]string, len(matrix))
	for idx, strToWrite := range matrix {
		formatted := strings.Join(strToWrite, " ")
		linesToWrite[idx] = formatted
	}

	if u {
		linesToWrite = unique(linesToWrite)
	}

	_, err = file.WriteString(strings.Join(linesToWrite, "\n"))
	if err != nil {
		panic(err)
	}
}

func unique(arr []string) []string {
	occurred := map[string]struct{}{}
	result := []string{}
	for e := range arr {
		if _, ok := occurred[arr[e]]; !ok {
			occurred[arr[e]] = struct{}{}
			result = append(result, arr[e])
		}
	}
	return result
}
