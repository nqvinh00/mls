package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/nqvinh00/mls/internal/utils"
	"github.com/nqvinh00/mls/models"
	"golang.org/x/term"
)

func getGlob(path string) []string {
	files, err := filepath.Glob(path)
	if err != nil {
		return []string{}
	}

	return files
}

func glob(writer io.Writer, path string, sortType string, showList, showAll, noColor, noLink, noIcon bool) {
	files := getGlob(path)
	parents := make(map[string][]string)
	parentsDir := make([]string, 0)
	for _, file := range files {
		dir := filepath.Dir(file)
		parents[dir] = append(parents[dir], file)
		parentsDir = append(parentsDir, dir)
	}

	sort.Slice(parentsDir, func(i, j int) bool {
		return parentsDir[i] < parentsDir[j]
	})

	for _, dir := range parentsDir {
		if !showAll && utils.IsHidden(path) {
			continue
		}

		children := getParentFiles(parents[dir], showAll)
		if len(children) == 0 {
			continue
		}

		_, _ = fmt.Fprintf(writer, "%s:\n", dir)
		printFiles(writer, children, sortType, showList, noColor, noLink, noIcon)
	}
}

func getParentFiles(files []string, showAll bool) []models.File {
	var filesToReturn []models.File
	for _, f := range files {
		file, err := models.NewFile(f)
		if err != nil {
			continue
		}

		if showAll || !file.IsHidden() {
			filesToReturn = append(filesToReturn, file)
		}
	}

	return filesToReturn
}

func getFiles(path string, showAll bool) ([]models.File, error) {
	var filesToReturn []models.File
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, i := range fileInfos {
		info, err := i.Info()
		if err != nil {
			return nil, err
		}
		file := models.File{FileInfo: info, Path: filepath.Join(path, i.Name())}
		if showAll || !file.IsHidden() {
			filesToReturn = append(filesToReturn, file)
		}
	}

	return filesToReturn, nil
}

func printFiles(writer io.Writer, files []models.File, sortType string, showList, noColor, noLink, noIcon bool) {
	utils.SortFiles(files, sortType)

	if showList {
		printAsList(writer, files, showList, noColor, noLink, noIcon)
	} else {
		printFilesInColumns(writer, files, noColor, true, noIcon)
	}
}

func printAsList(writer io.Writer, files []models.File, showList, noColor, noLink, noIcon bool) {
	var (
		sizes        []string
		totalSize    int64
		sizeCharLen  int
		modeCharLen  int
		nlinkCharLen int
		userCharLen  int
		groupCharLen int
	)

	for _, f := range files {
		totalSize += f.Size()
		sizeF := utils.ConvertFileSize(f.Size())
		sizes = append(sizes, sizeF)

		if len(sizeF) > sizeCharLen {
			sizeCharLen = len(sizeF)
		}

		if showList {
			if len(f.FileMode()) > modeCharLen {
				modeCharLen = len(f.FileMode())
			}

			if len(fmt.Sprintf("%d", f.Nlink())) > nlinkCharLen {
				nlinkCharLen = len(fmt.Sprintf("%d", f.Nlink()))
			}

			if runtime.GOOS != "windows" {
				if len(f.User()) > userCharLen {
					userCharLen = len(f.User())
				}

				if len(f.Group()) > groupCharLen {
					groupCharLen = len(f.Group())
				}
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "Total %s\n", utils.ConvertFileSize(totalSize))
	for i, f := range files {
		var line string
		if showList {
			line += fmt.Sprintf("%-*s ", modeCharLen, f.FileMode())
			line += fmt.Sprintf("%*d ", nlinkCharLen, f.Nlink())

			if runtime.GOOS != "windows" {
				user, group := f.User(), f.Group()
				if user == "" {
					user = group
				}

				line += fmt.Sprintf("%*s ", userCharLen-1, user)
				line += fmt.Sprintf("%*s ", groupCharLen-1, group)
			}

			line += fmt.Sprintf("%*s ", sizeCharLen, sizes[i])
			line += f.ModTime().Format(models.ISO8601) + " "
		}

		if !noColor {
			line += f.Colorize(noColor, noLink, noIcon)
		} else {
			line += f.Name()
		}

		_, _ = fmt.Fprintln(writer, line)
	}
}

func printFilesInColumns(writer io.Writer, files []models.File, noColor, noLink, noIcon bool) {
	maxNameLen := 0
	for _, file := range files {
		nameLen := utf8.RuneCountInString(file.Colorize(noColor, noLink, noIcon))
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
	}

	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		terminalWidth = 80
	}

	columnWidth := maxNameLen + 2
	columns := terminalWidth / columnWidth
	if columns < 1 {
		columns = 1
	}
	rows := (len(files) + columns - 1) / columns

	// Print each row
	var sb strings.Builder
	sb.Grow(len(files) * (columnWidth + 1))

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fileIndex := i + j*rows
			if fileIndex < len(files) {
				name := files[fileIndex].Colorize(noColor, noLink, noIcon)
				_, _ = sb.WriteString(name)
				_, _ = sb.WriteString(strings.Repeat(" ", columnWidth-utf8.RuneCountInString(name)))
			}
		}
		_ = sb.WriteByte('\n')
	}

	_, _ = fmt.Fprint(writer, sb.String())
}

func printAsTree(writer io.Writer, files []models.File, indent string, maxDepth int, depth int, showAll, noColor, noIcon bool) {
	if len(files) == 0 {
		return
	}

	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	for i, f := range files {
		connector := "├──"
		if i == len(files)-1 {
			connector = "└──"
		}
		_, _ = fmt.Fprintf(writer, "%s%s %s\n", indent, connector, f.Colorize(noColor, false, noIcon))
		if f.IsDir() {
			childIndent := indent + "    "
			if i != len(files)-1 {
				childIndent = indent + "│   "
			}
			childrenFiles, err := getFiles(f.Path, showAll)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}
			printAsTree(writer, childrenFiles, childIndent, maxDepth, depth+1, showAll, noColor, noIcon)
		}
	}
}
