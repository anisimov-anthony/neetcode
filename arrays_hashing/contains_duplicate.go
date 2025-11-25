package arrayshashing

func hasDuplicate(nums []int) bool {
	history := make(map[int]bool)
	for _, nm := range nums {
		_, exists := history[nm]
		if !exists {
			history[nm] = true
		} else {
			return true
		}
	}
	return false
}
