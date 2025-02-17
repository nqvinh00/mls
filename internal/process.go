package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Ls(writer io.Writer, paths []string, sortType string, showList, showAll, noColor, noLink, noIcon bool) {
	for _, path := range paths {
		if strings.ContainsRune(path, '*') {
			glob(writer, path, sortType, showList, showAll, noColor, noLink, noIcon)
		} else {
			files, err := getFiles(path, showAll)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}

			if len(paths) > 1 {
				_, _ = fmt.Fprintln(writer, filepath.Clean(path)+":")
			}
			printFiles(writer, files, sortType, showList, noColor, noLink, noIcon)
		}
	}
}

func Tree() {

}
