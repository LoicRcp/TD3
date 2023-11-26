package comsoc

func MajoritySWF(p Profile) (Count, error) {
	count := make(Count)

	for _, prefs := range p {
		if len(prefs) > 0 {
			firstChoice := prefs[0]
			count[firstChoice]++
		}
	}

	return count, nil
}

func MajoritySCF(p Profile) ([]Alternative, error) {
	Count, _ := MajoritySWF(p)
	return maxCount(Count), nil
}
