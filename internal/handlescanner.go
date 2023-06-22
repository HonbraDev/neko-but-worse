package internal

import (
	"bufio"

	"github.com/honbradev/neko/uwu"
)

func (i *Instance) HandleScanner(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Text()
		if err := i.WriteLine(uwu.UwU(line)); err != nil {
			return err
		}
	}

	return nil
}
