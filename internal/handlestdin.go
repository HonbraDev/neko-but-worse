package internal

import (
	"bufio"
	"os"
)

func (i *Instance) HandleStdin() error {
	return i.HandleScanner(bufio.NewScanner(os.Stdin))
}
