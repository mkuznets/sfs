package sps

import (
	"context"
	"fmt"
)

// Command is a common part of all subcommands.
type Command struct {
	Ctx         context.Context
	CriticalCtx context.Context
}

type Server struct {
	Command
}

func (cmd *Server) Execute([]string) error {
	fmt.Println("hello world!")
	return nil
}
