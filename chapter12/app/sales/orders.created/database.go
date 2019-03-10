package created

import "strconv"

var nOrder int

// This can be any database technology; it can differ between
// business components
func saveOrder(productIDs []string, userID, shippingTypeID string) string {

	id := strconv.Itoa(nOrder)
	nOrder++

	return id
}
