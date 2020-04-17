package urlen

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/url"
)

func decoderCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "decoder <content>",
		Short: "URL解码",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要编码的值")
			}

			value, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "PathUnescape()")
			}
			fmt.Println(value)

			return nil
		},
	}
	return &cmd
}
