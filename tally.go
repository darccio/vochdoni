package vochdoni

type Result map[string]int64

type Tally interface {
	Count(Referendum) Result
}
