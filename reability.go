package metatext

func (metaData *MetaData) CalculateLix() int {
	return metaData.WordCount/metaData.PeriodCount + (metaData.LongWordCount*100)/metaData.WordCount
}

func (metaData *MetaData) CalculateFleschReading(wordCount int, sentenceCount int, syllableCount int) int {
	return int(206.835 - 1.015*float64((metaData.WordCount/metaData.PeriodCount)) - 84.6*float64((metaData.SyllableCount/metaData.WordCount)))
}

func (metaData *MetaData) CalculateFleschGrade() int {
	return int(0.39*float64((metaData.WordCount/metaData.PeriodCount)) + 11.8*float64((metaData.SyllableCount/metaData.WordCount)) - 15.59)
}
