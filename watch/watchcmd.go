package watch

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var (
	interval int64
	output   string
)

// Command get watch command
func Command() *cobra.Command {
	cmd := cobra.Command{
		Use:   "watch -i <sec> -o <output> <cmd>",
		Short: "周期性地执行某个命令",
		RunE: func(c *cobra.Command, args []string) error {
			l := len(args)
			if l == 0 {
				return errors.New("指定的cmd无效")
			}

			// 定义command
			cmd := newExecutable(args)

			// 定义timer
			d := time.Duration(interval) * time.Second
			timer := time.NewTimer(d)
			defer timer.Stop()

			fmt.Printf("命令已开始执行。间隔：%ds\n", interval)
			for {
				t := <-timer.C
				cmd.Exec(t)
				timer.Reset(d)
			}
		},
	}

	cmd.Flags().Int64VarP(&interval, "interval", "i", 5, "指定间隔时间，默认值：5s")
	cmd.Flags().StringVarP(&output, "output", "o", "", "命令输出的文件")
	return &cmd
}

type executable struct {
	out *bytes.Buffer
	cmd *exec.Cmd
}

func (e *executable) Exec(t time.Time) {
	fmt.Printf("开执行 %s\n", t.Format("2006-01-02T03:04:05"))
	if err := e.cmd.Run(); err != nil {
		fmt.Println(e.out.String())
		return
	}
	fmt.Println(e.out.String())
}

func newExecutable(args []string) *executable {
	c := exec.Command(args[0], args[1:]...)
	var o bytes.Buffer
	c.Stdout = &o
	c.Stderr = &o
	return &executable{
		cmd: c,
		out: &o,
	}
}
