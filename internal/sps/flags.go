package sps

// Global is a group of common flags for all subcommands.
type Global struct {
	Debug bool `long:"debug" description:"Enable debug logging"`
}

type Flags struct {
	Global *Global        `group:"Global Options"`
	Server *ServerCommand `command:"server" description:"Start server"`
}
