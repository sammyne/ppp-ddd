package model

import (
	"errors"
	"math"
)

//type Money struct {
//	value float64 // float64 mirrors the decimal in C#
//}

type Money float64

const NilMoney Money = math.SmallestNonzeroFloat64

//func (money Money) GreaterThan(other Money) bool {
//	return money > other
//}

func NewMoney(value float64) (Money, error) {
	if err := ThrowErrorIfInvalid(value); nil != err {
		return NilMoney, err
	}

	return Money(value), nil
}

func ThrowErrorIfInvalid(value float64) error {
	if x := value * 100; int(math.Ceil(x)) != int(x) {
		return errors.New("there cannot be more than two decimal places")
	}

	if value < 0 {
		return errors.New("money cannot be a negative value")
	}

	return nil
}
