package cmd

import (
	"flag"

	"github.com/honbradev/neko/internal"
)

var (
	NumberAllLines = flag.Bool("b", false, "number nonempty output lines, overrides -n")
	NumberLines    = flag.Bool("n", false, "number all output lines")
	ShowEnds       = flag.Bool("e", false, "display $ at end of each line")
	SqueezeBlank   = flag.Bool("s", false, "squeeze repeated empty output lines to a single empty line")
)

func Run() error {
	flag.Parse()

	instance := internal.NewInstance()

	instance.NumberAllLines = *NumberAllLines
	instance.NumberLines = *NumberLines
	instance.ShowEnds = *ShowEnds
	instance.SqueezeBlank = *SqueezeBlank

	err := instance.HandleStdin()
	if err != nil {
		return err
	}

	return nil
}
