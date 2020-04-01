package metatext

import "fmt"

func (metaData *MetaData) CalculateReadabilityIndex(index int) (float32, error) {
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
func (metaData *MetaData) CalculateLix() float32 {
	return metaData.WordCount/metaData.PeriodCount + (metaData.LongWordCount*100)/metaData.WordCount
}

// 100-0
func (metaData *MetaData) CalculateFleschReading() float32 {
	return 206.835 - 1.015*(metaData.WordCount/metaData.PeriodCount) - 84.6*(metaData.SyllableCount/metaData.WordCount)
}

// 0-18
func (metaData *MetaData) CalculateFleschGrade() float32 {
	return 0.39*(metaData.WordCount/metaData.PeriodCount) + 11.8*(metaData.SyllableCount/metaData.WordCount) - 15.59
}
