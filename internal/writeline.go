package internal

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func (i *Instance) WriteLine(s string) error {
	var err error

	isBlank := len(s) == 0

	if i.SqueezeBlank {
		if isBlank {
			if i.previousLineBlank {
				i.currentLine++
				return nil
			}
			i.previousLineBlank = true
		} else {
			i.previousLineBlank = false
		}

	}

	if (i.NumberLines && !isBlank) || i.NumberAllLines {
		number := fmt.Sprint(i.currentLine)
		prefixLength := 6 - len(number)
		if prefixLength < 0 {
			prefixLength = 0
		}
		prefix := strings.Repeat(" ", prefixLength)
		s = prefix + number + "  " + s
		i.currentLine++
	}

	// handle ShowEnds
	if i.ShowEnds {
		s = s + "$"
	}

	_, err = io.WriteString(os.Stdout, s+"\n")

	return err
}
