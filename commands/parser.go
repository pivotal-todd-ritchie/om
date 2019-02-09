package commands

import (
	"github.com/jessevdk/go-flags"
)

func ParseArgs(data interface{}, args []string) ([]string, error) {
	return flags.NewParser(data, flags.HelpFlag|flags.PassDoubleDash).ParseArgs(args)
}
