//go:build windows
package models

import (
	"syscall"
)

func (f File) Attrs() uint32 {
	return f.FileInfo.Sys().(*syscall.Win32FileAttributeData).FileAttributes
}

func (f File) IsDir() bool {
	if f.IsLink() {
		return f.Attrs()&syscall.FILE_ATTRIBUTE_DIRECTORY != 0
	}

	return f.FileInfo.IsDir()
}

func (f File) IsHidden() bool {
	if f.Name()[0] == '.' {
		return true
	}

	return f.Attrs()&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}

func (f File) Nlink() uint32 {
	encode, err := syscall.UTF16PtrFromString(f.Path)
	if err != nil {
		return 0
	}

	handle, err := syscall.CreateFile(encode, syscall.GENERIC_READ, syscall.FILE_SHARE_READ, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return 0
	}

	fileInfo := syscall.ByHandleFileInformation{}
	if err = syscall.GetFileInformationByHandle(handle, &fileInfo); err != nil {
		return 0
	}

	return fileInfo.NumberOfLinks
}
