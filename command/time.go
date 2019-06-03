package command

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TimeCommand struct {
	Meta
	help string

	value string
	unit string
	format string
}

const(
	ms=1000000
	sec=ms * 1000
)

func (c *TimeCommand) Run(args []string) int {

	err:=c.parseArgs(args)
	if err!=nil{
		return 1
	}

	v,err:=strconv.ParseInt(c.value,10,64)
	if err!=nil{
		fmt.Printf("参数value(%s)格式无效！",c.value)
		return 1
	}

	switch c.unit {
	case "sec":
		v=v*sec
	case "ms":
		v=v*ms
	}

	t:=time.Unix(0,v)
	fmt.Printf("Result: %s -> %s",c.value,t.Format(c.format))

	return 0
}

func (c *TimeCommand) Synopsis() string {
	return "将UNIX时间戳显示为人类可读时间"
}

func (c *TimeCommand) Help() string {
	helpText := "将UNIX时间戳显示为人类可读时间"
	return strings.TrimSpace(helpText)
}

func (c *TimeCommand) parseArgs(args []string) error{
	fs:= flag.NewFlagSet("time_unix", flag.ContinueOnError)
	fs.StringVar(&c.value,"value","","指定要转的时间值。")
	fs.StringVar(&c.unit,"unit","unix_sec","时间单位。可能值：sec,ms,ns")
	fs.StringVar(&c.format,"format",time.RFC3339,"时间格式。")
	return fs.Parse(args)

}
