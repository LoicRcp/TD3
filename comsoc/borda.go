package comsoc

func BordaSWF(p Profile) (Count, error) {
	count := make(Count)
	numAlts := len(p[0]) // Supposant que tous les votants classent toutes les alternatives

	for _, prefs := range p {
		for rank, alt := range prefs {
			count[alt] += numAlts - rank - 1
		}
	}

	return count, nil
}

func BordaSCF(p Profile) ([]Alternative, error) {
	Count, _ := BordaSWF(p)
	return maxCount(Count), nil
}
