package parser

import (
	"go-demo-reptile/bussiness/engine"
	"go-demo-reptile/model"
	"log"
	"regexp"
	"strings"
)

var regTitle = regexp.MustCompile(`<title>(.+)</title>`)
var regKeywords = regexp.MustCompile(`<meta name="keywords" content="(.+)"[^>]*>`)
var regDes = regexp.MustCompile(`<meta name="description" content="(.+)"`)

func ParseInfo(contents []byte, url string) engine.ParseResult {

	page := model.PageInfo{}
	page.Url = url

	page.Title = regPaperInfo(contents, regTitle)
	keywords := regPaperInfo(contents, regKeywords)
	if len(keywords) > 0 {
		page.Keywords = strings.Split(keywords, ",")
	} else {
		page.Keywords = []string{}
	}
	page.Description = regPaperInfo(contents, regDes)

	result := engine.ParseResult{
		Items: []interface{}{page},
	}

	nextRequest := ParseBaseUrl(contents)
	result.Requests = nextRequest.Requests

	log.Printf("page: %v", page)

	return result

}

func regPaperInfo(contents []byte, compile *regexp.Regexp) string {

	submatch := compile.FindSubmatch(contents)

	if len(submatch) >= 2 {
		return string(submatch[1])
	}
	return ""

}
