package main

import (
	"Vessel/src/common/term"
	"Vessel/src/server/daemon"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = "Vessel is a tiny container."

func initConfig() {
	//term.CheckFile(constant.ConfigFile,"")
}

// 创建Cli前，初始化Logrus
func before(*cli.Context) error {
	_, _, stderr := term.StdStreams()
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: term.RFC3339NanoFixed,
		FullTimestamp:   true,
	})
	logrus.SetOutput(stderr)
	return nil
}

// 创建Cli
func main() {
	after(nil)
	//_, _, stderr := term.StdStreams()
	//app := cli.NewApp()
	//app.Name = "Vessel-Server"
	//app.Usage = usage
	//app.Before = before
	//app.After = after
	//if err := app.Run(os.Args); err != nil {
	//	logrus.Fatal(err)
	//	_, _ = fmt.Fprintf(stderr, term.ErrorFmt, "App", err)
	//	os.Exit(1)
	//}
}

// Cli创建后，启动TCP监听
func after(*cli.Context) error {
	wildcard := term.Wildcard(886)
	d := daemon.Daemon{TcpConfig: wildcard}
	return d.ActivateListeners()
}
