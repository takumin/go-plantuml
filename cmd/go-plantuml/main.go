package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"golang.org/x/crypto/ssh/terminal"

	puml "github.com/takumin/go-plantuml"
)

func main() {
	var (
		server string
		format string
		style  string
	)
	flag.StringVar(&server, "server", "http://plantuml.com/plantuml", "PlantUML Server Address")
	flag.StringVar(&format, "format", "png", "Output Format (Options: png, svg)")
	flag.StringVar(&style, "style", "link", "Output Style (Options: encode, link)")
	flag.Parse()

	if terminal.IsTerminal(0) {
		flag.PrintDefaults()
		return
	}

	if server == "" {
		fmt.Fprintf(os.Stderr, "require plantuml server address")
		return
	}

	if format != "png" && format != "svg" {
		fmt.Fprintf(os.Stderr, "require format option png or svg")
		return
	}

	if style != "encode" && style != "link" {
		fmt.Fprintf(os.Stderr, "require style option encode or link")
		return
	}

	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return
	}
	if !utf8.Valid(raw) {
		fmt.Fprintln(os.Stderr, "invalid utf8 strings:", raw)
		return
	}

	enc, err := puml.Encode(raw)
	if err != nil {
		fmt.Fprintln(os.Stderr, "plantuml encoding:", err)
		return
	}

	switch style {
	case "encode":
		fmt.Fprintf(os.Stdout, "%s\n", enc)
	case "link":
		fmt.Fprintf(os.Stdout, "<a href=\"%s/uml/%s\"><img src=\"%s/%s/%s\" alt=\"PlantUML PNG Image\" /></a>\n", server, enc, server, format, enc)
	}
}
