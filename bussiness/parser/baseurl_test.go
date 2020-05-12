package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseBaseUrl(t *testing.T) {

	fetch, e := ioutil.ReadFile("testdata/sina.com.cn.txt")
	if e != nil {
		panic(e)
	}

	const hopeCount = 590
	result := ParseBaseUrl(fetch)
	if len(result.Requests) != hopeCount {
		t.Errorf(" hope request %d , but %d", hopeCount, len(result.Requests))
	}

	hopeList := []string{
		"军事", "国内", "国际",
	}

	for i, n := range hopeList {
		if result.Items[i].(string) != n {
			t.Errorf(" hope item is %s , now:%s", n, result.Items[i].(string))
		}
	}
}
