package ip

import (
	"github.com/spf13/cobra"
)

// Command 获取IP相关命令
func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ip <command>",
		Short: "IP相关命令",
	}

	cmd.AddCommand(getIPLocatingCmd())

	return &cmd
}
