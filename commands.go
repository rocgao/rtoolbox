package main

import (
	"github.com/RocGao/rtoolbox/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"time_unix": func() (cli.Command, error) {
			return &command.TimeCommand{
				Meta: *meta,
			}, nil
		},
		"url_coder":func() (cli.Command,error){
			return &command.UrlCoderCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
