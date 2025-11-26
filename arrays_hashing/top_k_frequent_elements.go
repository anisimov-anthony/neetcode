package arrayshashing

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	stats := make([][]int, len(nums)+1)
	result := make([]int, 0, k)

	for _, nm := range nums {
		freqs[nm]++
	}

	for nm, fr := range freqs {
		stats[fr] = append(stats[fr], nm)
	}

	for i := len(stats) - 1; i >= 0 && len(result) <= k; i-- {
		if len(stats[i]) > 0 {
			result = append(result, stats[i]...)
		}
	}

	return result[:k]
}
