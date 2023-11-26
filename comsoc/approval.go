package comsoc

func ApprovalSWF(p Profile, thresholds []int) (Count, error) {
	count := make(Count)

	for i, prefs := range p {
		threshold := thresholds[i]
		for _, alt := range prefs[:threshold] {
			count[alt]++
		}
	}

	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) ([]Alternative, error) {
	Count, _ := ApprovalSWF(p, thresholds)
	return maxCount(Count), nil
}
