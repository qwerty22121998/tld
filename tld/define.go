package tld

import (
	"net/url"
	"os"

	"github.com/qwerty22121998/go-trie/trie"
)

const tldFile = "https://publicsuffix.org/list/public_suffix_list.dat"
const dataFile = "./tld.txt"

var tld = trie.New()
var debugMode = false
var autoUpdate = true

type URL struct {
	Domain    string // domain
	Subdomain string // sub domain
	TLD       string // top level domain
	Port      string // port
	*url.URL         // net/url
}

func handle(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func isFileExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}
