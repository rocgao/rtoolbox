package logicnum

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"math"
	"strconv"
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "logicnum <number>",
		Short: "显示逻辑和",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供<number>参数")
			}

			logicNum, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			var result []int64
			for logicNum > 0 {
				r := logicNum % 2
				result = append(result, r)
				logicNum = logicNum / 2
			}
			fmt.Printf("%v\n", result)

			l := len(result)
			for index, item := range result {
				if item == 0 {
					continue
				}

				v := math.Pow(2, float64(l-index-1))
				fmt.Print(v)
				fmt.Print(",")
			}

			if l > 0 {
				fmt.Println()
			}

			return nil
		},
	}
	return &cmd
}
