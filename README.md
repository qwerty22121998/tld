# tld
Use [Go](https://golang.org/) and [Public Suffix List](https://publicsuffix.org/list/) stored in a [Trie](https://en.wikipedia.org/wiki/Trie) to extract top level domain 

## Installation
Run the following code
>go get github.com/qwerty22121998/tld
## Usage
Struct of an URL
```go
type URL struct {
	Domain    string // domain
	Subdomain string // sub domain
	TLD       string // top level domain
	Port      string // port
	*url.URL         // net/url
}
```

To show the progress
>tld.SetDebugMode(true/false) //Default False

To auto update before parse
>tld.SetAutoUpdate(true/false) //Default True

To force update the lastest tld list

>tld.Update() 

To create new Parser
>parse := tld.NewParser()

To parse an url 

> url, err := parse.Parse("http://github.com")