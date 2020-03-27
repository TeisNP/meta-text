package metatext

func CalculateLix(wordCount int, longWordCount int, periodCount int) int {
	return wordCount/periodCount + (longWordCount*100)/wordCount
}

func CalculateFleschReading(wordCount int, sentenceCount int, syllableCount int) int {
	return int(206.835 - 1.015*float64((wordCount/sentenceCount)) - 84.6*float64((syllableCount/wordCount)))
}

func CalculateFleschGrade(wordCount int, sentenceCount int, syllableCount int) int {
	return int(0.39*float64((wordCount/sentenceCount)) + 11.8*float64((syllableCount/wordCount)) - 15.59)
}
