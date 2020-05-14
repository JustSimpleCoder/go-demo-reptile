package main

import (
	"fmt"
	"regexp"
	"strings"
)

const text = " hahah  Woiskobe@gmail.com       "

func main() {

	re := regexp.MustCompile(`(?i)(w.+)@(.+\..+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match[0][1])
	fmt.Printf(strings.TrimSpace(text))
}
