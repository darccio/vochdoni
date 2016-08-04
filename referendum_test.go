package vochdoni

import (
	"testing"
)

func TestCreateReferendum(t *testing.T) {
	r := &Referendum{}
	r.SetUp()
	r.Approve()
	r.Open()
	r.Close()
	tl := r.Tally
	result := tl.Count(r)
	result.Publish()
	r.Archive()
}
