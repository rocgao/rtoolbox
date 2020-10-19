package timestamp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	ms  = 1000000
	sec = ms * 1000
)

var (
	unit = "sec"
)

// Command get the Timestamp command
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

			// 根据单位转换成纳秒 nsec
			switch unit {
			case "sec":
				value = value * sec
			case "ms":
				value = value * ms
			case "tick":
				value = (value - 621355968000000000) * 100 // 621355968000000000指1970/1/1T0:0:0到0001/1/1T0:0:0的tick数，1 tick = 100 nsec
			}

			t := time.Unix(0, value)
			fmt.Printf("%s -> %s\n", color.RedString(args[0]),color.GreenString(t.Format("2006-01-02 15:04:05")))
			return nil
		},
	}
	cmd.Flags().StringVarP(&unit, "unit", "u", "sec", "单位（sec: 秒 ms:毫秒 tick: DateTimeOffset）")

	return &cmd
}
