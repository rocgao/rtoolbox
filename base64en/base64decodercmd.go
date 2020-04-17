package base64en

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func decoderCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "decoder <content>",
		Short: "Base64解码",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要解码的值")
			}

			data, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				return errors.Wrap(err, "DecodeString()")
			}
			fmt.Println(string(data))
			return nil
		},
	}
	return &cmd
}
