package factory

import "math"

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

func NewCalculator(kind int, principal float64, time int) InterestCalculator {
	if kind == 1 {
		return &simpleInterest{
			principal:      principal,
			rateOfInterest: 10.00,
			time:           12,
		}
	}
	return &compoundInterest{
		principal:      principal,
		rateOfInterest: 10.00,
		time:           12,
	}
}

type InterestCalculator interface {
	Calculate() float64
}

func (si *simpleInterest) Calculate() float64 {
	return si.time * (1 + (si.rateOfInterest * si.time))
}

func (ci *compoundInterest) Calculate() float64 {
	return math.Pow(ci.principal*(1+(ci.rateOfInterest/12)), ci.time)
}
