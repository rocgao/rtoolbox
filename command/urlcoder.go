package command

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

type UrlCoderCommand struct{
	Meta

	help string
	fs *flag.FlagSet

	action string
	value string
}

func (c *UrlCoderCommand) Help() string {
	c.init()
	return c.help
}

func (c *UrlCoderCommand) Run(args []string) int {
	c.parseArgs(args)

	if c.action==""{
		fmt.Println("请指定-value参数")
		return 1
	}

	if c.action=="encode"{
		fmt.Println(url.PathEscape(c.value))
		return 0
	}
	if c.action=="decode"{
		v,err:=url.PathUnescape(c.value)
		if err != nil {
			fmt.Printf("Error: %s",err.Error())
			return 1
		}
		fmt.Println(v)
		return 0
	}

	fmt.Printf("action:%s值无效\n",c.action)
	return 1
}

func (c* UrlCoderCommand) Synopsis() string {
	return "URL转码工具"
}

func (c *UrlCoderCommand) init(){
	fs:= flag.NewFlagSet("urlcoder", flag.ContinueOnError)
	fs.StringVar(&c.value,"value","","指定要转的值。")
	fs.StringVar(&c.action,"action","encode","操作类型。可能值：encode,decode")

	helpBuilder:=&strings.Builder{}
	fs.SetOutput(helpBuilder)
	fs.Usage()
	c.help =helpBuilder.String()

	c.fs=fs
}

func (c *UrlCoderCommand) parseArgs(args []string) error{
	c.init()
	return c.fs.Parse(args)
}

