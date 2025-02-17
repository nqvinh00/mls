package utils

import (
	"fmt"
	"os"
	"sort"

	"github.com/nqvinh00/mls/models"
)

func SortFiles(files []models.File, sortType string) {
	switch sortType {
	case "s", "size":
		sort.Slice(files, func(i, j int) bool {
			return files[i].Size() < files[j].Size()
		})
	case "d", "date":
		sort.Slice(files, func(i, j int) bool {
			return files[i].FileInfo.ModTime().Before(files[j].FileInfo.ModTime())
		})
	case "x", "extension":
		sort.Slice(files, func(i, j int) bool {
			return files[i].Ext() < files[j].Ext()
		})
	case "t", "type":
		sort.Slice(files, func(i, j int) bool {
			return files[i].Type() < files[j].Type()
		})
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unknown sort type: %s\n", sortType)
		os.Exit(1)
	}
}
