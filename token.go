package vochdoni

type Authorization struct {
	Referendum Referendum
	Token      Token
}

type Token struct {
	Value string
}
