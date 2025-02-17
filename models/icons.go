package models

// Nerd Font (FiraCode patched)
const (
	ArrowIcon     = "\uea9c"
	DirectoryIcon = "\ueaf7"
	FileIcon      = "\uf4a5" // generic
	LinkFileIcon  = "\uf481"
	LinkDirIcon   = "\ueaed"

	ArchiveIcon = "\uf410"
	ImageIcon   = "\uf1c5"
	AudioIcon   = "\uf1c7"
	VideoIcon   = "\uf1c8"

	CIcon      = "\ue61e"
	CPPIcon    = "\ue61d"
	PythonIcon = "\ue73c"
	ShellIcon  = "\uea85"
	HtmlIcon   = "\uf13b"
	PDFIcon    = "\uf1c1"
	WordIcon   = "\uf1c2"
	ExcelIcon  = "\uf1c3"
	CSVIcon    = "\ueefc"
	JsonIcon   = "\ueb0f"
	LogIcon    = "\uf18d"
	JarIcon    = "\ue204"
	XMLIcon    = "\ue619"
	YamlIcon   = "\ue6a8"
	GoIcon     = "\ue627"
	SqlIcon    = "\ue706"
)

var extensionIconMap = map[string]string{
	".c":     CIcon,
	".h":     CIcon,
	".cpp":   CPPIcon,
	".cc":    CPPIcon,
	".cp":    CPPIcon,
	".cxx":   CPPIcon,
	".hpp":   CPPIcon,
	".py":    PythonIcon,
	".ipynb": PythonIcon,
	".pyc":   PythonIcon,
	".go":    GoIcon,

	".sh":    ShellIcon,
	".shell": ShellIcon,
	".bat":   ShellIcon,

	".7z":     ArchiveIcon,
	".a":      ArchiveIcon,
	".ar":     ArchiveIcon,
	".bz2":    ArchiveIcon,
	".cab":    ArchiveIcon,
	".cpio":   ArchiveIcon,
	".deb":    ArchiveIcon,
	".gz":     ArchiveIcon,
	".rar":    ArchiveIcon,
	".rpm":    ArchiveIcon,
	".tar":    ArchiveIcon,
	".tar.gz": ArchiveIcon,
	".tgz":    ArchiveIcon,
	".zip":    ArchiveIcon,

	".jpg":  ImageIcon,
	".jpeg": ImageIcon,
	".png":  ImageIcon,
	".gif":  ImageIcon,
	".svg":  ImageIcon,
	".webp": ImageIcon,
	".bmp":  ImageIcon,
	".ico":  ImageIcon,
	".tiff": ImageIcon,
	".tif":  ImageIcon,
	".heif": ImageIcon,
	".heic": ImageIcon,

	".mp3":  AudioIcon,
	".aac":  AudioIcon,
	".au":   AudioIcon,
	".flac": AudioIcon,
	".m4a":  AudioIcon,
	".wav":  AudioIcon,

	".avi":  VideoIcon,
	".mkv":  VideoIcon,
	".mp4":  VideoIcon,
	".mov":  VideoIcon,
	".webm": VideoIcon,
	".wmv":  VideoIcon,
	".m4v":  VideoIcon,
	".flv":  VideoIcon,
	".mpv":  VideoIcon,

	".jar":  JarIcon,
	".java": JarIcon,

	".pdf":  PDFIcon,
	".docx": WordIcon,
	".doc":  WordIcon,
	".xlsx": ExcelIcon,
	".xls":  ExcelIcon,
	".csv":  CSVIcon,

	".json": JsonIcon,
	".log":  LogIcon,

	".html": HtmlIcon,
	".yaml": YamlIcon,
	".xml":  XMLIcon,
	".yml":  YamlIcon,
}

func GetFileIcon(ext string) string {
	if icon, ok := extensionIconMap[ext]; ok {
		return icon
	}

	return FileIcon
}
