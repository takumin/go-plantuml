package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var (
		server string
		format string
		style  string
	)
	flag.StringVar(&server, "server", "http://plantuml.com/plantuml", "PlantUML Server Address")
	flag.StringVar(&format, "format", "png", "Output Format (Options: png, svg, ascii)")
	flag.StringVar(&style, "style", "png", "Output Style (Options: enc, link, img)")
	flag.Parse()

	if terminal.IsTerminal(0) {
		flag.PrintDefaults()
		return
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return
	}
	if !utf8.Valid(buf) {
		fmt.Fprintln(os.Stderr, "invalid utf8 strings")
		return
	}
}
