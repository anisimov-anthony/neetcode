package arrayshashing

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	maxLen := 0

	for num := range set {
		if !set[num-1] {
			currentNum := num
			currentLen := 1

			for set[currentNum+1] {
				currentNum++
				currentLen++
			}

			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	return maxLen
}
