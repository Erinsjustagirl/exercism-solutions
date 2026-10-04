package cards

// OutOfBounds returns a boolean if the index given is out of the bounds of the slice.
func OutOfBounds(slice []int, index int) bool {
    if index >= len(slice) || index < 0 {
        return true;
    } else  {
        return false;
    }
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9};
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if OutOfBounds(slice, index) {
        return -1;
    } else {
    	return slice[index];
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    result := slice;
	if OutOfBounds(slice, index) {
	    result = append(slice, value);
	} else {
	    slice[index] = value;
		result = slice;
	}
    return result;
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...);
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    result := slice;
	if OutOfBounds(slice, index) {
	    return slice;
	} else {
	    result = append(slice[0:index], slice[index+1:]...);
	}
    return result;
}
