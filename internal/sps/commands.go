package sps

import (
	"context"
)

// Command is a common part of all subcommands.
type Command struct {
	Ctx         context.Context
	CriticalCtx context.Context
}

type ServerCommand struct {
	Command
}

func (cmd *ServerCommand) Execute([]string) error {
	server := Server{
		Addr: ":8080",
		Api:  &Api{},
	}
	return server.Start()
}
