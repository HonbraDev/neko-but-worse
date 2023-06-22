package internal

type Instance struct {
	// runtime vars

	currentLine       int
	previousLineBlank bool

	// config vars

	NumberLines    bool
	NumberAllLines bool
	ShowEnds       bool
	SqueezeBlank   bool
}

func NewInstance() *Instance {
	return &Instance{
		currentLine: 1,
	}
}
