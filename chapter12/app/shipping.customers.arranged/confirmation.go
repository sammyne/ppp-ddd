package arranged

type ShippingStatus int

const (
	Success ShippingStatus = iota
	Failure
)

type Confirmation struct {
	Status ShippingStatus
}
