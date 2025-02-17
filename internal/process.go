package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nqvinh00/mls/models"
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

func Tree(writer io.Writer, paths []string, maxDepth int, showAll, noColor, noIcon bool) {
	wd, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, path := range paths {
		if filepath.IsAbs(path) {
			err = os.Chdir(path)
		} else {
			err = os.Chdir(filepath.Join(wd, path))
		}

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}

		files, err := getFiles(".", showAll)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}

		cleanPath := filepath.Clean(path)
		if !noColor {
			cleanPath = fmt.Sprintf("%s%s%s", models.GetColor(models.DirType), cleanPath, models.Reset)
		}
		_, _ = fmt.Fprintln(writer, cleanPath)

		printAsTree(writer, files, "", maxDepth, 0, showAll, noColor, noIcon)
	}
}
