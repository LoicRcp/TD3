package comsoc

import (
	"errors"
)

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

func maxCount(count Count) (bestAlts []Alternative) {
	bestScore := 0
	for alternative, score := range count {
		if score > bestScore {
			bestScore = score
			bestAlts = []Alternative{alternative}
		} else if score == bestScore {
			bestAlts = append(bestAlts, alternative)
		}
	}
	return bestAlts
}

func checkProfile(prefs []Alternative, alts []Alternative) error {
	altsCounter := make(map[Alternative]int, len(alts))
	for _, alternative := range prefs {
		altsCounter[alternative] += 1
		if altsCounter[alternative] > 1 {
			return errors.New("L'unicité n'est pas respecté dans ce profil.")
		}
	}
	if len(prefs) != len(alts) {
		return errors.New("L'ensemble des alternatives n'est pas représenté dans ce profil.")
	}
	return nil
}
