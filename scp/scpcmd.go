package scp

import (
	"os"

	goscp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// Command 获取scp命令实例
func Command() *cobra.Command {
	cmd := cobra.Command{
		Use: "scp <file> <remote-host>",
		Short: "使用SCP协议传输文件",
		RunE: func(c *cobra.Command, args []string) error {
			// clientConfig, err := auth.PasswordKey("root", "1qaz2wsx!QAZ@WSX[].", nil)
			clientConfig, err := auth.PasswordKey("root", "FJOiQBjHm0jBXYVxG1lr", ssh.InsecureIgnoreHostKey())
			client := goscp.NewClient("106.14.81.166:22", &clientConfig)

			err = client.Connect()
			if err != nil {
				return errors.Wrap(err, "call client.Connect()")
			}
			defer client.Close()

			f, err := os.Open("/Users/rocgao/Downloads/ILSpy_binaries_4.0.1.4530.zip")
			if err != nil {
				return errors.Wrap(err, "call os.Open()")
			}
			defer f.Close()

			err = client.CopyFile(f, "~/", "0655")
			if err != nil {
				return errors.Wrap(err, "call client.CopyFile()")
			}

			return nil
		},
	}
	return &cmd
}

func parseArgs(args []string) (remoteHost, file string) {
	panic("")
}
