package term

import (
	"io"
	"os"
)

/*
	标准化工具类，用于屏蔽系统差异
*/

const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}
