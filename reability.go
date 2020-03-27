package metatext

import "fmt"

func (metaData *MetaData) CalculateReadabilityIndex(index int) (int, error) {
	switch index {
	case 0:
		return metaData.CalculateLix(), nil
	case 1:
		return metaData.CalculateFleschReading(), nil
	case 2:
		return metaData.CalculateFleschGrade(), nil
	default:
		return 0, fmt.Errorf("Readbility index calculation not supported for case: %d", index)
	}
}

func (metaData *MetaData) CalculateLix() int {
	return metaData.WordCount/metaData.PeriodCount + (metaData.LongWordCount*100)/metaData.WordCount
}

func (metaData *MetaData) CalculateFleschReading() int {
	return int(206.835 - 1.015*float64((metaData.WordCount/metaData.PeriodCount)) - 84.6*float64((metaData.SyllableCount/metaData.WordCount)))
}

func (metaData *MetaData) CalculateFleschGrade() int {
	return int(0.39*float64((metaData.WordCount/metaData.PeriodCount)) + 11.8*float64((metaData.SyllableCount/metaData.WordCount)) - 15.59)
}
