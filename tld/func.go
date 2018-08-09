package tld

import (
	"bufio"
	"errors"
	"net/http"
	"net/url"
	"os"
)

func fetch() {
	file, err := os.Open(dataFile)
	handle(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tld.Add(scanner.Text())
	}
}

func update() bool {

	debug("Connecting to", tldFile, "...")

	resp, err := http.Get(tldFile)
	if err != nil {
		debug("Have problem with internet connection")
		debug("Use offline data")
		return false
	}
	debug("Connected!")

	debug("Reading data...")

	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	debug("Created!")

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || text[0] == '/' {
			continue
		}
		if text[0] == '!' || text[0] == '*' {
			text = text[2:]
		}
		tld.Add(reverse(text))

	}

	if isFileExist(dataFile) {
		debug("Data file found! Removing...")

		err := os.RemoveAll(dataFile)
		handle(err)
		debug("File removed!")

	}

	debug("Creating data file...")

	file, err := os.Create(dataFile)
	handle(err)
	debug("Created!")

	debug("Export data to", dataFile)

	for _, v := range tld.Prefix("") {
		file.WriteString(reverse(v) + "\n")
	}

	debug("Exported!")

	file.Close()
	return true

}

// Update : Update list of tld to newest
func Update() bool {
	debug("Force Update")
	return update()
}

// SetDebugMode : Set pkg debug mode
func SetDebugMode(mode bool) {
	debugMode = mode
}

// SetAutoUpdate : set autoupdate mode
func SetAutoUpdate(update bool) {
	autoUpdate = update
}

// NewParser : return a parser
func NewParser() func(string) (*URL, error) {
	if !autoUpdate || !update() {
		fetch()
	}

	debug("Parse Initialized!")

	return func(s string) (*URL, error) {
		debug("Parsing", s)
		url, err := url.Parse(s)

		if err != nil {
			return nil, err
		}

		if len(url.Host) == 0 {
			debug("Parsed")
			return &URL{URL: url}, nil

		}
		domain, port := split(url.Host)

		tldNow := reverse(tld.CommonWord(reverse(domain)))

		if len(tldNow) == 0 {
			debug("TLD not found!")
			return nil, errors.New("tld not found!")
		}
		domain = domain[:len(domain)-len(tldNow)]

		subdomain := ""
		for i := len(domain) - 1; i >= 0; i-- {
			if domain[i] == '.' {
				subdomain = domain[:i]
				domain = domain[i+1:]
				break
			}
		}
		debug("Parsed")
		return &URL{
			Subdomain: subdomain,
			Domain:    domain,
			TLD:       tldNow,
			Port:      port,
			URL:       url,
		}, nil
	}
}
