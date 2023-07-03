package term

import (
	"io"
	"os"
)

/*
	标准化工具类，用于屏蔽系统差异
*/

func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}
