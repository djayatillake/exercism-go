package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	unitmap := map[string]int{}
	unitmap["quarter_of_a_dozen"] = 3
	unitmap["half_of_a_dozen"] = 6
	unitmap["dozen"] = 12
	unitmap["small_gross"] = 120
	unitmap["gross"] = 144
	unitmap["great_gross"] = 1728
	return unitmap
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
	if ok {
		bill[item] += units[unit]
		return true } else {
		return false }
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, ok := bill[item]
	_, ok2 := units[unit]
	if ok && ok2 && bill[item] - units[unit] == 0 {
		delete(bill, item)
		return true
	} else if ok && ok2 && bill[item] - units[unit] > 0 {
		bill[item] = bill[item] - units[unit]
		return true
	} else {
		return false
	}
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	i, ok := bill[item]
	return i, ok
}
