package uwu

import "strings"

func w(s string) string {
	s = strings.ReplaceAll(s, "r", "w")
	s = strings.ReplaceAll(s, "l", "w")
	s = strings.ReplaceAll(s, "R", "W")
	s = strings.ReplaceAll(s, "Th", "D")
	s = strings.ReplaceAll(s, "th", "d")
	s = strings.ReplaceAll(s, "v", "ff")
	s = strings.ReplaceAll(s, "V", "Ff")
	s = strings.ReplaceAll(s, "Ove", "Uv")
	s = strings.ReplaceAll(s, "ove", "uv")

	return s
}
