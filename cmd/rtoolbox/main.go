package main

import (
	"fmt"
	"os"

	"github.com/rocgao/rtoolbox/base64en"
	"github.com/rocgao/rtoolbox/logicnum"
	"github.com/rocgao/rtoolbox/timestamp"
	"github.com/rocgao/rtoolbox/urlen"
	"github.com/rocgao/rtoolbox/uuid"
	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
)

func main() {
	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func rootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use:     "rtoolbox",
		Short:   "Roc的工具箱CLI版本",
		Version: fmt.Sprintf("%s %s", Version, GitCommit),
	}
	rootCmd.AddCommand(uuid.Command(), timestamp.Command(), base64en.Command(), urlen.Command(), logicnum.Command())
	return &rootCmd
}
