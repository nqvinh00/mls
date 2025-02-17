//go:build !windows
package models

import (
	"fmt"
	"os/user"
	"syscall"
)

func (f File) Attrs() uint32 {
	return 0
}

func (f File) IsHidden() bool {
	return f.Name()[0] == '.'
}

func (f File) Stat_t() syscall.Stat_t {
	return *f.FileInfo.Sys().(*syscall.Stat_t)
}

func (f File) Group() string {
	group, err := user.LookupGroupId(fmt.Sprint(f.Stat_t().Gid))
	if err != nil {
		return "unknown"
	}

	return group.Name
}

func (f File) User() string {
	user, err := user.LookupId(fmt.Sprint(f.Stat_t().Uid))
	if err != nil {
		return "unknown"
	}

	return user.Username
}

func (f File) Nlink() uint32 {
	return uint32(f.Stat_t().Nlink)
}
