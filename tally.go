package vochdoni

type Result struct {
	Referendum Referendum
	Data       map[string]int64
}

func (r *Result) Publish() {

}

type Tally struct {
	Method string
}

func (t *Tally) Count(r *Referendum) (rs *Result) {
	return
}
