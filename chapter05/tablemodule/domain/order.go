package domain

type Orders struct{}

func NewOrders(dataSet map[int64]interface{}) *Orders {
	return new(Orders)
}
