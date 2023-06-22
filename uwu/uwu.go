package uwu

import "strings"

func UwU(s string) string {
	// l => w
	s = w(s)

	// replace ! with random faces
	for strings.Contains(s, "!") {
		s = strings.Replace(s, "!", " "+randomFace(), 1)
	}

	// Na -> Nya (all vowels)
	for _, vowel := range vowels {
		s = strings.ReplaceAll(s, "N"+vowel, "Ny"+vowel)
		s = strings.ReplaceAll(s, "n"+vowel, "ny"+vowel)
	}

	return s
}
