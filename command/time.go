package command

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type TimeCommand struct {
	Meta

	help string
	fs *flag.FlagSet

	value int64
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

	v:=c.value
	switch c.unit {
	case "sec":
		v=v*sec
	case "ms":
		v=v*ms
	}

	t:=time.Unix(0,v)
	fmt.Printf("Result: %d -> %s",c.value,t.Format(c.format))

	return 0
}

func (c *TimeCommand) Synopsis() string {
	return "将UNIX时间戳显示为人类可读时间。执行rtoolbox time_unix --help获取该命令的详细用法"
}

func (c *TimeCommand) Help() string {
	c.init()
	return c.help
}

func (c *TimeCommand) init(){
	fs:= flag.NewFlagSet("time_unix", flag.ContinueOnError)
	fs.Int64Var(&c.value,"value",0,"指定要转的时间值。")
	fs.StringVar(&c.unit,"unit","unix_sec","时间单位。可能值：sec,ms,ns")
	fs.StringVar(&c.format,"format",time.RFC3339,"时间格式。")

	helpBuilder:=&strings.Builder{}
	fs.SetOutput(helpBuilder)
	fs.Usage()
	c.help =helpBuilder.String()

	c.fs=fs
}


func (c *TimeCommand) parseArgs(args []string) error{
	c.init()
	return c.fs.Parse(args)
}
