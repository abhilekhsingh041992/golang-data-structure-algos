
func binarySearch(array *int, n int, target int) int {
	var low int = 0
	var high int = n - 1
	for low <= high {
		var middle int = (low + high) / 2
		if array[middle] == target {
			return middle
		} else {
			if array[middle] < target {
				low = middle+1
			} else {
				high = middle-1
			}
		}
	}
	return -1
}
func binarySearchFirst(array *int, n int, target int) int {
	var low int = 0
	var high int = n - 1
	var first int = -1
	for low <= high {
		var middle int = (low + high) / 2
		if array[middle] == target {
			first = middle
			high = middle-1
		} else {
			if array[middle] < target {
				low = middle+1
			} else {
				high = middle-1
			}
		}
	}
	return first
}
func binarySearchLast(array *int, n int, target int) int {
	var low int = 0
	var high int = n - 1
	var last int = -1
	for low <= high {
		var middle int = (low + high) / 2
		if array[middle] == target {
			last = middle
			low = middle+1
		} else {
			if array[middle] < target {
				low = middle+1
			} else {
				high = middle-1
			}
		}
	}
	return last
}

