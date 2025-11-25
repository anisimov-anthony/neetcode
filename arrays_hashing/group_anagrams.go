package arrayshashing

func groupAnagrams(strs []string) [][]string {
	table := make(map[[26]int][]string)

	for _, str := range strs {
		key := [26]int{}
		for i := range len(str) {
			key[str[i]-'a']++
		}
		table[key] = append(table[key], str)
	}

	result := make([][]string, 0, len(table))

	for _, vals := range table {
		result = append(result, vals)
	}

	return result
}
