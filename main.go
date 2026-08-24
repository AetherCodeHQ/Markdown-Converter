package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: markdown-converter <file.md>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	inList := false
	for _, line := range lines {
		t := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(t, "# "):
			closeList(&inList)
			fmt.Printf("<h1>%s</h1>\n", strings.TrimPrefix(t, "# "))
		case strings.HasPrefix(t, "## "):
			closeList(&inList)
			fmt.Printf("<h2>%s</h2>\n", strings.TrimPrefix(t, "## "))
		case strings.HasPrefix(t, "### "):
			closeList(&inList)
			fmt.Printf("<h3>%s</h3>\n", strings.TrimPrefix(t, "### "))
		case strings.HasPrefix(t, "- "), strings.HasPrefix(t, "* "):
			if !inList {
				fmt.Println("<ul>")
				inList = true
			}
			fmt.Printf("  <li>%s</li>\n", t[2:])
		default:
			closeList(&inList)
			if t != "" {
				fmt.Printf("<p>%s</p>\n", t)
			}
		}
	}
	closeList(&inList)
}

func closeList(inList *bool) {
	if *inList {
		fmt.Println("</ul>")
		*inList = false
	}
}
