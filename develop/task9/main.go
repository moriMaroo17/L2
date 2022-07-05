package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/opesun/goquery"
)

type flags struct {
	input    string
	output   string
	download bool
}

var flagsStore *flags

func (f *flags) readFile() ([]string, error) {
	file, err := os.Open(f.input)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	result := make([]string, 0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func main() {

	flagsStore = new(flags)

	flag.StringVar(&flagsStore.input, "i", "", "input file")
	flag.StringVar(&flagsStore.output, "O", "", "output file")
	flag.BoolVar(&flagsStore.download, "d", false, "download all data from site")

	flag.Parse()

	var domains []string

	switch true {
	case flagsStore.input != "":
		var err error
		domains, err = flagsStore.readFile()
		if err != nil {
			log.Fatal("cannot read file: ", err)
		}
	default:
		fmt.Println(flag.Args())
		domains = flag.CommandLine.Args()
	}

	if len(domains) == 0 {
		log.Fatal("no domains specified")
	}

	for _, domain := range domains {
		if err := getPage(domain); err != nil {
			fmt.Println(err)
		}
	}

}

// getPage is a function that get a page and work with it depending on flags
func getPage(domain string) error {

	switch true {
	case flagsStore.download:
		err := getPageResourses(domain)
		if err != nil {
			fmt.Println(err)
		}
	case flagsStore.output != "":
		if err := downloadPageResource(domain, flagsStore.output); err != nil {
			return err
		}
	default:
		response, err := http.Get(domain)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		scanner := bufio.NewScanner(response.Body)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	return nil
}

// getPageResourses is a function that gets a list of all page resources and downloads it
func getPageResourses(domain string) error {
	parsedURL, err := goquery.ParseUrl(domain)
	if err != nil {
		return err
	}

	for _, url := range parsedURL.Find("").Attrs("href") {
		if strings.Contains(url, ".css") ||
			strings.Contains(url, ".js") ||
			strings.Contains(url, ".jpeg") ||
			strings.Contains(url, ".png") ||
			strings.Contains(url, ".gif") {
			fmt.Printf("trying download: %s\n", url)
			splittedURL := strings.Split(url, "/")
			if splittedURL[0] == "." {
				url = strings.Replace(url, "./", domain+"/", 1)
			}
			filename := splittedURL[len(splittedURL)-1]
			if err := downloadPageResource(url, filename); err != nil {
				fmt.Println("error downloading page resource ", filename, " Error: ", err)
			}
		}

	}
	return nil
}

// DownloadPageResource is a helper function to download single page resource
func downloadPageResource(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, response.Body); err != nil {
		return err
	}
	return nil
}
