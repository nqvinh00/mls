package main

import (
	"bufio"
	"log"
	"os"
	"runtime"

	"github.com/nqvinh00/mls/internal"
	"github.com/nqvinh00/mls/internal/utils"

	"github.com/spf13/pflag"
)

var (
	paths      []string
	showAll    bool
	showList   bool
	showAsTree bool
	noColor    bool
	noIcon     bool
	noLink     bool
	depth      int
	sort       string
)

func parseArgs() {
	pflag.CommandLine.SortFlags = false
	pflag.BoolVarP(&showAll, "all", "a", false, "Show all files including hidden ones")
	pflag.BoolVarP(&showList, "list", "l", false, "Show list of files")
	pflag.BoolVarP(&showAsTree, "tree", "t", false, "Show files as a tree")
	pflag.BoolVarP(&noColor, "no-color", "C", false, "Disable color output")
	pflag.BoolVarP(&noIcon, "no-icon", "I", false, "Disable icon output")
	pflag.BoolVarP(&noLink, "no-link", "L", false, "Disable link output")
	pflag.StringVarP(&sort, "sort", "s", "s", "Sort files by extenstion (x), type (t), size (s), or date (d)")
	pflag.IntVarP(&depth, "depth", "d", -1, "Maximum depth of the tree")

	help := pflag.BoolP("help", "h", false, "Show help")

	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	paths = pflag.Args()
}

func main() {
	parseArgs()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	if runtime.GOOS == "windows" && !noColor {
		if err := utils.EnableColor(); err != nil {
			log.Fatal(err)
		}
	}

	if showAsTree {
		internal.Tree(writer, paths, depth, showAll, noColor, noIcon)
	} else {
		internal.Ls(writer, paths, sort, showList, showAll, noColor, noLink, noIcon)
	}
}
