package base64en

import "github.com/spf13/cobra"

// Command return the command instance of base64en
func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "base64 <command> <content>",
		Short: "Base64编解码",
	}
	cmd.AddCommand(encoderCommand(), decoderCommand())
	return &cmd
}
