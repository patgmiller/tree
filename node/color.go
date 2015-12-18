package node

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const Escape = "\x1b"
const (
	Reset int = 0
	Bold  int = 1
	Black int = iota + 28
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// ANSIColor
func ANSIColor(node *Node, s string) string {
	var color int
	switch ext := filepath.Ext(node.Name()); strings.ToLower(ext) {
	case ".bat", ".btm", ".cmd", ".com", ".dll", ".exe":
		color = Green
	case ".arj", ".bz2", ".deb", ".gz", ".lzh", ".rpm", ".tar", ".taz", ".tb2", ".tbz2",
		".tbz", ".tgz", ".tz", ".tz2", ".z", ".zip", ".zoo":
		color = Red
	case ".asf", ".avi", ".bmp", ".flac", ".gif", ".jpg", "jpeg", ".m2a", ".m2v", ".mov",
		".mp3", ".mpeg", ".mpg", ".ogg", ".ppm", ".rm", ".tga", ".tif", ".wav", ".wmv",
		".xbm", ".xpm":
		color = Magenta
	default:
		// IsDir
		if node.IsDir() {
			color = Blue
		}
		// IsSymlink
		if node.Mode()&os.ModeSymlink == os.ModeSymlink {
			// IsOrphan
			if _, err := filepath.EvalSymlinks(node.path); err != nil {
				// Error link color
				return fmt.Sprintf("%s[40;%d;%dm%s%s[%dm", Escape, Bold, Red, s, Escape, Reset)
			} else {
				color = Cyan
			}
		}
		// IsSocket
		if node.Mode()&os.ModeSocket == os.ModeSocket {
			return fmt.Sprintf("%s[40;%d;%dm%s%s[%dm", Escape, Bold, Magenta, s, Escape, Reset)
		}
		// IsFifo
		if node.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
			return fmt.Sprintf("%s[40;%dm%s%s[%dm", Escape, Yellow, s, Escape, Reset)
		}
		// IsBlk - a block special file (a device like a disk)
		if node.Mode()&os.ModeDevice == os.ModeDevice {
			return fmt.Sprintf("%s[40;%d;01m%s%s[%dm", Escape, Yellow, s, Escape, Reset)
		}
		// IsChr
		if node.Mode()&os.ModeCharDevice == os.ModeCharDevice {
			return fmt.Sprintf("%s[40;%d;01m%s%s[%dm", Escape, Yellow, s, Escape, Reset)
		}
		// IsExecutable
		// https://groups.google.com/forum/#!msg/golang-nuts/_6YqjJdfYyA/mOnV5zSpP8oJ
	}
	return fmt.Sprintf("%s[%d;%dm%s%s[%dm", Escape, Bold, color, s, Escape, Reset)
}

// TODO: HTMLColor
