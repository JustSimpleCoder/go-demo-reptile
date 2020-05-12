package main

import (
	"fmt"
	"regexp"
	"strings"
)

const text = " hahah  woiskobe@gmail.com       "

func main() {

	re := regexp.MustCompile(`(.+)@(.+\..+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	fmt.Printf(strings.TrimSpace(text))
}
