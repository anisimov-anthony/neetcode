package arrayshashing

func isAnagram(s string, t string) bool {
	sfreqs := make(map[byte]int)
	tfreqs := make(map[byte]int)

	for i := range len(s) {
		sfreqs[s[i]]++
	}

	for i := range len(t) {
		tfreqs[t[i]]++
	}

	for sb, sfr := range sfreqs {
		if tfreqs[sb] != sfr {
			return false
		}
	}

	for tb, tfr := range tfreqs {
		if sfreqs[tb] != tfr {
			return false
		}
	}

	return true
}
