
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	var b strings.Builder
	b.WriteString("# Report\n\n| File | Size |\n|---|---|\n")
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		b.WriteString("| " + p + " | " + strconv.FormatInt(info.Size(), 10) + " |\n")
		return nil
	})
	fmt.Print(b.String())
}
