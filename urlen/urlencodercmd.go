package urlen

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/url"
)

func encoderCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "encoder <content>",
		Short: "URL编码",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要编码的值")
			}

			value := url.PathEscape(args[0])
			fmt.Println(value)

			return nil
		},
	}
	return &cmd
}
