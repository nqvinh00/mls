package utils

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nqvinh00/mls/models"
)

func IsHidden(path string) bool {
	if path == "." {
		return false
	}

	parts := strings.Split(path, string(filepath.Separator))
	for i := 0; i < len(parts); i++ {
		f := filepath.Join(parts[i:]...)
		if strings.Contains(f, "..") {
			absolutePath, err := filepath.Abs(f)
			if err != nil {
				continue
			}
			f = absolutePath
		}

		file, err := models.NewFile(f)
		if err != nil {
			continue
		}

		if file.IsHidden() {
			return true
		}
	}

	return false
}

func ConvertFileSize(size int64) string {
	if size < 1<<10 {
		return strconv.FormatInt(size, 10) + models.SizeUnit[0]
	}

	fileSize, exp := float64(size), 0
	for i := fileSize / (1 << 10); i >= 1; i /= 1 << 10 {
		exp++
		fileSize /= 1 << 10
	}

	return fmt.Sprintf("%.2f%s", fileSize, models.SizeUnit[exp-1])
}
