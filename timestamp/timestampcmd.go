package timestamp

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

const (
	ms  = 1000000
	sec = ms * 1000
)

var (
	unit = "sec"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "timestamp <timestamp> --unit sec",
		Short: "将Unix timestamp时间转换可读日期格式",
		Long:  "将Unix timestamp时间转换可读日期格式",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要转换的timestamp")
			}

			value, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return errors.Wrap(err, "解析timestamp出错")
			}

			switch unit {
			case "sec":
				value = value * sec
			case "ms":
				value = value * ms
			}

			t := time.Unix(0, value)
			fmt.Printf("%s -> %s", args[0], t.Format(time.RFC3339))
			return nil
		},
	}
	cmd.Flags().StringVarP(&unit, "unit", "u", "sec", "单位（sec: 秒 ms:毫秒）")

	return &cmd
}
