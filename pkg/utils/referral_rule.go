package utils

func ReferralRule(count int) (level string, rate float64) {
	if count <= 0 {
		return "", 0
	}
	if count > 0 && count <= 30 {
		rate = 0.1
		level = "Standard"
	} else if count >= 31 && count <= 100 {
		rate = 0.2
		level = "Level I"
	} else if count >= 101 && count <= 1000 {
		rate = 0.3
		level = "Level II"
	} else if count >= 1001 && count <= 5000 {
		rate = 0.5
		level = "Level III"
	} else {
		rate = 0.7
		level = "Level IV"
	}
	return level, rate
}
