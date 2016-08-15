package vochdoni

type Status int

const (
	Created = iota
	ApprovalPending
	InProgress
	TallyPending
	Counting
	Done
	Archived
)

type Referendum struct {
	Title        string
	Description  string
	Tally        Tally
	Status       Status
	UID          string
	Organization Token
	// Implicit
	// Ballot      Ballot
	// Votes       []Vote
}

func (r *Referendum) Approve() {

}

func (r *Referendum) Freeze() {

}

func (r *Referendum) Open() {

}

func (r *Referendum) Close() {

}

func (r *Referendum) Archive() {

}
