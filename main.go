package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	translated, err := gtranslate.TranslateWithFromTo(
		text,
		gtranslate.FromTo{
			From: "en",
			To:   "zh-cn",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en   : %s\nzh-cn: %s \n", text, translated)
}
