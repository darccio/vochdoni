package vochdoni

type Option struct {
	Vote     Vote
	Value    string
	Modifier uint16
}

type Vote struct {
	Referendum Referendum
	// Implicit
	// Options    []Option
}
