package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	os.FileInfo
	Path string
}

func (f File) Name() string {
	return f.FileInfo.Name()
}

func (f File) Size() int64 {
	return f.FileInfo.Size()
}

func (f File) Ext() string {
	return filepath.Ext(f.Name())
}

func (f File) Mode() string {
	return f.FileInfo.Mode().String()
}

func (f File) IsLink() bool {
	return f.FileInfo.Mode()&os.ModeSymlink != 0
}

func (f File) IsSymLinkBroken() bool {
	target, err := filepath.EvalSymlinks(f.Path)
	if err != nil {
		return true
	}

	_, err = os.Stat(target)
	return err != nil
}

func (f File) Link() string {
	target, err := os.Readlink(f.Path)
	if err != nil {
		return ""
	}

	absPath, err := filepath.Abs(target)
	if err != nil {
		return ""
	}

	if absPath != "" && strings.HasPrefix(absPath, "..") {
		return absPath
	}

	return target
}

func (f File) PrettyPrint(noLink, noIcon bool) string {
	name := f.Name()
	if !noLink && f.IsLink() {
		arrow := "->"
		if !noIcon {
			arrow = ArrowIcon
		}
		name = name + " " + arrow + " " + f.Link()
	}

	if !noIcon {
		name = f.Icon() + " " + name
	}

	return name
}

func (f File) Icon() string {
	if f.IsLink() {
		if f.IsDir() {
			return LinkDirIcon
		}
		return LinkFileIcon
	}

	if f.IsDir() {
		return DirectoryIcon
	}

	return GetFileIcon(f.Ext())
}

func (f File) Type() int {
	if f.IsLink() {
		if f.IsSymLinkBroken() {
			return BrokenType
		}
		return LinkType
	}

	if f.IsDir() {
		return DirType
	}

	return GetFileType(f.Ext())
}

func (f File) FileMode() string {
	return f.FileInfo.Mode().String()
}

func (f File) Colorize(noColor, noLink, noIcon bool) string {
	pretty := f.PrettyPrint(noLink, noIcon)
	if noColor {
		return pretty
	}

	return fmt.Sprintf("%s%s%s", GetColor(f.Type()), pretty, Reset)
}

func NewFile(path string) (File, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return File{}, err
	}

	return File{
		FileInfo: fileInfo,
		Path:     path,
	}, nil
}
