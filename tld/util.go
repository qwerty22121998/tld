package tld

import "fmt"

func debug(s ...string) {
	if debugMode {
		fmt.Println(s)
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func split(url string) (string, string) {
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == ':' {
			return url[:i], url[i+1:]
		}
	}
	return url, ""
}
