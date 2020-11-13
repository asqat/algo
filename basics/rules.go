package basics

// FindLargest returns the largest value from arr
func FindLargest(arr []int) int {
	largest := arr[0]
	for i := range arr {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}

// ContainDuplicates checks an existent of duplicates from arr
func ContainDuplicates(arr []int) bool {
	for i := range arr {
		for j := range arr {
			if i != j {
				if arr[i] == arr[j] {
					return true
				}
			}
		}
	}
	return false
}
