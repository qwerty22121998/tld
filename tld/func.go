package tld

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"sort"
)

type URL struct {
	Domain, Subdomain, TLD string
	*url.URL
}

const DOMAIN_REGEX = `((http|https)\:/\/)?\w+(\.\w+)+(\:\d+)?`

func fetch() {
	file, err := os.Open(dataFile)
	handle(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tld = append(tld, scanner.Text())
	}
}

func update() bool {

	// connect to host
	if debugMode {
		println("Connecting to", tldFile, "...")
	}
	resp, err := http.Get(tldFile)
	if err != nil {
		println("Have problem with internet connection")
		println("Use offline data")
		return false
	}
	if debugMode {
		println("Connected!")
	}

	// read the data file
	if debugMode {
		println("Reading data...")
	}
	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	if debugMode {
		println("Created!")
	}

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || text[0] == '/' {
			continue
		}
		if text[0] == '!' || text[0] == '*' {
			text = text[2:]
		}
		tld = append(tld, text)

	}

	sort.Strings(tld)

	// check if file exist
	if isFileExist(dataFile) {
		if debugMode {
			println("Data file found! Removing...")
		}
		err := os.RemoveAll(dataFile)
		handle(err)
		if debugMode {
			println("File removed!")
		}
	}
	// create data file
	if debugMode {
		println("Creating data file...")
	}
	file, err := os.Create(dataFile)
	handle(err)
	if debugMode {
		println("Created!")
	}

	if debugMode {
		println("Export data to", dataFile)
	}
	for _, v := range tld {
		file.WriteString(v + "\n")
	}

	if debugMode {
		println("Exported!")
	}

	file.Close()
	return true

}

func newParser() (func(domain string), error) {
	if !update() {
		fetch()
	}

	return nil, nil
}
