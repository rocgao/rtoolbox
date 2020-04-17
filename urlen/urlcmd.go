package urlen

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "url <command> <content>",
		Short: "URL编解码",
	}
	cmd.AddCommand(encoderCommand(), decoderCommand())
	return &cmd
}
