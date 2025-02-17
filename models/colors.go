package models

import (
	"fmt"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31"
	Green   = "\033[32"
	Yellow  = "\033[33"
	Blue    = "\033[34"
	Magenta = "\033[35"
	Cyan    = "\033[36"
	Gray    = "\033[37"
	White   = "\033[97"

	NonFormat = "m"
	Bold      = ";1m"
	Italic    = ";3m"
	Underline = ";4m"
)

var typeColorMap = map[int]string{
	FileType:       fmt.Sprintf("%s%s", Green, NonFormat),
	DirType:        fmt.Sprintf("%s%s", Blue, Bold),
	LinkType:       fmt.Sprintf("%s%s", Cyan, NonFormat),
	BrokenType:     fmt.Sprintf("%s%s", Red, Italic),
	ArchiveType:    fmt.Sprintf("%s%s", Red, Underline),
	ExecutableType: fmt.Sprintf("%s%s", Green, Bold),
	CodeType:       fmt.Sprintf("%s%s", Magenta, NonFormat),
	ImageType:      fmt.Sprintf("%s%s", Yellow, NonFormat),
	AudioType:      fmt.Sprintf("%s%s", Green, Bold),
	VideoType:      fmt.Sprintf("%s%s", Yellow, NonFormat),
}

func GetColor(t int) string {
	if color, ok := typeColorMap[t]; ok {
		return color
	}
	return typeColorMap[FileType]
}
