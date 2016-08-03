package vochdoni

import (
	"time"
)

type Option struct {
	Vote     Vote
	Value    string
	Modifier uint16
}

type Vote struct {
	Referendum  Referendum
	Voter       Token
	CastTime    time.Time
	Tracker     string
	Fingerprint string
	// Implicit
	// Options    []Option
}
