package vochdoni

type Referendum struct {
	Title       string
	Description string
	Method      Tally
	UID         string
	// Implicit
	// Ballot      Ballot
	// Votes       []Vote
}
