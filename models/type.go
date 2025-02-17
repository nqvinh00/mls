package models

const (
	DirType = iota
	FileType
	LinkType
	BrokenType
	ArchiveType
	ExecutableType
	CodeType
	ImageType
	AudioType
	VideoType
)

func GetFileType(ext string) int {
	if fType, ok := extensionTypeMap[ext]; ok {
		return fType
	}
	return FileType
}

var extensionTypeMap = map[string]int{
	".c":     CodeType,
	".h":     CodeType,
	".cpp":   CodeType,
	".cc":    CodeType,
	".cp":    CodeType,
	".cxx":   CodeType,
	".hpp":   CodeType,
	".py":    CodeType,
	".ipynb": CodeType,
	".pyc":   CodeType,
	".go":    CodeType,
	".jar":   CodeType,
	".java":  CodeType,

	".sh":    ExecutableType,
	".shell": ExecutableType,
	".bat":   ExecutableType,

	".7z":     ArchiveType,
	".a":      ArchiveType,
	".ar":     ArchiveType,
	".bz2":    ArchiveType,
	".cab":    ArchiveType,
	".cpio":   ArchiveType,
	".deb":    ArchiveType,
	".gz":     ArchiveType,
	".rar":    ArchiveType,
	".rpm":    ArchiveType,
	".tar":    ArchiveType,
	".tar.gz": ArchiveType,
	".tgz":    ArchiveType,
	".zip":    ArchiveType,

	".png":  ImageType,
	".jpg":  ImageType,
	".jpeg": ImageType,
	".gif":  ImageType,
	".bmp":  ImageType,
	".svg":  ImageType,
	".ico":  ImageType,
	".tiff": ImageType,
	".tif":  ImageType,
	".heif": ImageType,
	".heic": ImageType,
	".webp": ImageType,

	".mp3":  AudioType,
	".wav":  AudioType,
	".aac":  AudioType,
	".au":   AudioType,
	".flac": AudioType,
	".m4a":  AudioType,

	".avi":  VideoType,
	".mkv":  VideoType,
	".mp4":  VideoType,
	".mov":  VideoType,
	".webm": VideoType,
	".wmv":  VideoType,
	".m4v":  VideoType,
	".flv":  VideoType,
	".mpv":  VideoType,
}
