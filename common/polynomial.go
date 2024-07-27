package common

type Polynomial struct {
	degree int
}

func NewPolynomial(degree int) *Polynomial {
	return &Polynomial{degree: degree}
}
