package arrayshashing

import "slices"

func twoSum(nums []int, target int) []int {
	history := make(map[int]int)

	for i, nm := range nums {
		history[nm] = i
	}

	result := make([]int, 0, 2)
	for i, nm := range nums {
		if ti, ok := history[target-nm]; ok {
			if i != ti {
				result = append(result, ti)
				result = append(result, i)
				break
			}
		}
	}

	slices.Sort(result)
	return result
}
