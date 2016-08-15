package vochdoni

import (
	"testing"
)

func TestCreateReferendum(t *testing.T) {
	settings := &Settings{}
	r := &Referendum{}
	setUp(r)
	for i := 0; i < settings.ApprovalThreshold; i++ {
		r.Approve()
	}
	r.Freeze()
	r.Open()
	simulateVoting()
	r.Close()
	tl := r.Tally
	result := tl.Count(r)
	result.Publish()
	r.Archive()
}

func setUp(r *Referendum) {

}

func simulateVoting() {

}
