package main

import (
	"Vessel/src/common/term"
	"github.com/sirupsen/logrus"
)

/*
	Client入口文件
	用户通过命令行与Client交互，Client通过restfulAPI与Daemon(Server)交互
*/

func main() {
	// 用户终端的标准流
	stdin, stdout, stderr := term.StdStreams()
	logrus.SetOutput(stderr)

}
