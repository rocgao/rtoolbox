package command

import (
	"encoding/base64"
	"flag"
	"fmt"
	"strings"
)

type Base64CoderCommand struct{
	Meta

	help string
	fs *flag.FlagSet

	action string
	value string
}

func (c *Base64CoderCommand) Help() string {
	c.init()
	return c.help
}

func (c *Base64CoderCommand) Run(args []string) int {
	c.parseArgs(args)

	if c.action==""{
		fmt.Println("请指定-value参数")
		return 1
	}

	if c.action=="encode"{
		bytes:=[]byte(c.value)
		v:=base64.StdEncoding.EncodeToString(bytes)
		fmt.Println(v)
		return 0
	}
	if c.action=="decode"{
		v,err:=base64.StdEncoding.DecodeString(c.value)
		if err != nil {
			fmt.Printf("Error: %s",err.Error())
			return 1
		}
		fmt.Println(string(v))
		return 0
	}

	fmt.Printf("action:%s值无效\n",c.action)
	return 1
}

func (c* Base64CoderCommand) Synopsis() string {
	return "Base64编码工具"
}

func (c *Base64CoderCommand) init(){
	fs:= flag.NewFlagSet("base64coder", flag.ContinueOnError)
	fs.StringVar(&c.value,"value","","指定要转的值。")
	fs.StringVar(&c.action,"action","encode","操作类型。可能值：encode,decode")

	helpBuilder:=&strings.Builder{}
	fs.SetOutput(helpBuilder)
	fs.Usage()
	c.help =helpBuilder.String()

	c.fs=fs
}

func (c *Base64CoderCommand) parseArgs(args []string) error{
	c.init()
	return c.fs.Parse(args)
}