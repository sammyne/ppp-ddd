package domain

type Commander interface {
	Execute() error
}
