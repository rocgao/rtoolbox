package base64en

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func encoderCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "encoder <content>",
		Short: "生成base64字符串",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要编码的值")
			}

			bytes := []byte(args[0])
			v := base64.StdEncoding.EncodeToString(bytes)
			fmt.Println(v)
			return nil
		},
	}

	return &cmd
}
