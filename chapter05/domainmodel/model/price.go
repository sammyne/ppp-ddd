package model

type Price Money

//func NewPrice(amount Money) Price {
//	return Price(amount)
//}

func (price Price) Increment() Money {
	amount := Money(price)

	switch {
	case amount >= 0.01 && amount <= 0.99:
		return amount + 0.05
	case amount >= 1 && amount <= 4.99:
		return amount + 0.2
	case amount >= 5 && amount <= 14.99:
		return amount + 0.5
	}

	return amount + 1
}
