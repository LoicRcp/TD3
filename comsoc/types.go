package comsoc

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

func rank(alt Alternative, prefs []Alternative) int {
	for rank, alternative := range prefs {
		if alternative == alt {
			return rank
		}
	}
	return -1
}

func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}
