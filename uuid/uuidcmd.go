package uuid

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/spf13/cobra"
	"strings"
)

var (
	upper  bool
	noDash bool
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "uuid",
		Short: "生成GUID",
		Long:  "生成GUID",
		Run: func(cmd *cobra.Command, args []string) {
			u := guid.New()

			var v string
			if upper {
				v = u.StringUpper()
			} else {
				v = u.String()
			}
			if noDash {
				v = strings.ReplaceAll(v, "-", "")
			}
			fmt.Println(v)
		},
	}

	cmd.Flags().BoolVarP(&upper, "upper", "u", false, "大写形式")
	cmd.Flags().BoolVarP(&noDash, "no-dash", "n", false, "不包含-")
	return &cmd
}
