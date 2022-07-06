package factory

import "math"

// A Factory Pattern says that just define an interface creating an implementation blueprint but let the subclasses
// decide which class to instantiate.

type simpleInterest struct {
	principal      float64
	rateOfInterest float64
	time           float64
}

type compoundInterest struct {
	principal      float64
	rateOfInterest float64
	time           float64
}

func NewCalculator(kind int, principal float64, time float64) InterestCalculator {
	if kind == 1 {
		return &simpleInterest{
			principal:      principal,
			rateOfInterest: 10.00,
			time:           time,
		}
	}
	return &compoundInterest{
		principal:      principal,
		rateOfInterest: 2.00,
		time:           time,
	}
}

type InterestCalculator interface {
	Calculate() float64
}

// Calculate users simpleInterest to instantiate
func (si *simpleInterest) Calculate() float64 {
	return si.time * (1 + (si.rateOfInterest * si.time))
}

// Calculate users compound to instantiate
func (ci *compoundInterest) Calculate() float64 {
	return math.Pow(ci.principal*(1+(ci.rateOfInterest/12)), ci.time)
}
