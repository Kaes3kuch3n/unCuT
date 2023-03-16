package bin

func GetLightest[T Bin](bins []T) (bin T) {
	if len(bins) == 0 {
		panic("Empty bin slice not allowed in GetLightest")
	}
	lightestIndex := 0
	lightestWeight := bins[0].GetWeight()
	for i, b := range bins {
		if i == 0 {
			continue
		}
		if b.GetWeight() < lightestWeight {
			lightestIndex = i
		}
	}
	return bins[lightestIndex]
}
