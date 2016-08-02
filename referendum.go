package vochdoni

type Referendum struct {
	Title       string
	Description string
	Method      Tally
	// Implicit
	// Ballot      Ballot
	// Votes       []Vote
}
