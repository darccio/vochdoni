package vochdoni

type Solution struct {
	Ballot      Ballot
	Title       string
	Description string
}

type Ballot struct {
	Referendum Referendum
	// Implicit
	// Solutions  []Solution
}
