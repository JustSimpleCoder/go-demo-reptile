package parser

import (
	"go-demo-reptile/bussiness/engine"
	"regexp"
	"strings"
)

const urlReg = `<a href="([^"]+)"[^>]*>([^<]+)</a>`

func ParseBaseUrl(contents []byte) engine.ParseResult {

	compile := regexp.MustCompile(urlReg)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		url := strings.TrimSpace(string(m[1]))
		result.Items = append(result.Items, strings.TrimSpace(string(m[2])))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseInfo(bytes, url)
			},
		})

	}

	return result
}
