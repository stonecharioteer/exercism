package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = map[string]int {
        "dozen": 12,
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill = map[string]int {}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if value, ok:= units[unit]; ok {
		bill[item] += value
    	return true
    } else {
        return false
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok:=units[unit]; !ok {
        return false
    }
    if _, ok:=bill[item]; !ok {
        return false
    }
    if bill[item] < units[unit] {
        return false
    }
    bill[item] -= units[unit]
    if bill[item] == 0 {
        delete(bill, item)
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, ok:=bill[item]
    return v, ok
}
