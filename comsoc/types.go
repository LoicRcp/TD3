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
