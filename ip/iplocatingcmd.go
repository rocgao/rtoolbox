package ip

import (
	"fmt"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func getIPLocatingCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "locate <IP>",
		Short: "IP定位",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errors.New("请提供需要定位的IP")
			}

			region, err := ip2region.New("ip2region.db")
			defer region.Close()
			if err != nil {
				return errors.Wrap(err, "call ip2region.New()")
			}

			ipInfo, err := region.MemorySearch(args[0])
			if err != nil {
				return errors.Wrap(err, "call region.MemorySearch()")
			}
			fmt.Printf("%v\n", ipInfo)

			return nil
		},
	}

	return &cmd
}
