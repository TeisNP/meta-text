package metatext

import "fmt"

func (metaData *MetaData) CalculateReadabilityIndex(index int) (uint, error) {
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

// 0-55+
func (metaData *MetaData) CalculateLix() uint {
	return metaData.WordCount/metaData.PeriodCount + (metaData.LongWordCount*100)/metaData.WordCount
}

// 100-0
func (metaData *MetaData) CalculateFleschReading() uint {
	return uint(206.835 - 1.015*float64((metaData.WordCount/metaData.PeriodCount)) - 84.6*float64((metaData.SyllableCount/metaData.WordCount)))
}

// 0-18
func (metaData *MetaData) CalculateFleschGrade() uint {
	return uint(0.39*float64((metaData.WordCount/metaData.PeriodCount)) + 11.8*float64((metaData.SyllableCount/metaData.WordCount)) - 15.59)
}
