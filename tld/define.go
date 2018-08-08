package tld

import (
	"os"
)

const tldFile = "https://publicsuffix.org/list/public_suffix_list.dat"
const dataFile = "./tld.txt"

var tld = make([]string, 0)
var debugMode = false

func handle(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func isFileExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// Update : Update list of tld to newest
func Update() bool {
	return update()
}

// SetDebugMode : Set pkg debug mode
func SetDebugMode(mode bool) {
	debugMode = mode
}
