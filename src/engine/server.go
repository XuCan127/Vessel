package engine

import (
	"Vessel/src/common/term"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func newDaemonCommand() *cobra.Command {
	opts := newDaemonOptions(config.New())

	cmd := &cobra.Command{
		Use:           "dockerd [OPTIONS]",
		Short:         "A self-sufficient runtime for containers.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.flags = cmd.Flags()
			daemonCli := NewDaemonCli()
			return daemonCli.start(opts)
		},
		DisableFlagsInUseLine: true,
		Version:               fmt.Sprintf("%s, build %s", dockerversion.Version, dockerversion.GitCommit),
	}
	cli.SetupRootCommand(cmd)

	flags := cmd.Flags()
	flags.BoolP("version", "v", false, "Print version information and quit")
	flags.StringVar(&opts.configFile, "config-file", defaultDaemonConfigFile, "Daemon configuration file")
	opts.InstallFlags(flags)
	installConfigFlags(opts.daemonConfig, flags)
	installServiceFlags(flags)

	return cmd
}

func main() {
	// 第一步开启日志记录程序
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: term.RFC3339NanoFixed,
		FullTimestamp:   true,
	})
	// 设置
	_, _, stderr := term.StdStreams()
	logrus.SetOutput(stderr)

	cmd := newDaemonCommand()
	cmd.SetOut(stdout)
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
}
