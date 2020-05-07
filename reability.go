package metatext

import (
	"fmt"
	"math"
	"strings"

	"github.com/teisnp/syllables"
)

func (text *Text) CalculateReadabilityIndex(index int) (float32, error) {
	switch index {
	case 0:
		return text.CalculateLix(), nil
	case 1:
		return text.CalculateFleschReading(), nil
	case 2:
		return text.CalculateFleschGrade(), nil
	case 3:
		return text.CalculateDaleChall(), nil
	case 4:
		index, err := text.CalculateGunningFog()
		return index, err
	case 5:
		x, y, err := text.CalculateFryGraph()
		return x + y, err
	case 6:
		index, err := text.CalculateSmog()
		return index, err
	case 7:
		index, err := text.CalculateForcast()
		return index, err
	default:
		return 0, fmt.Errorf("Readbility index calculation not supported for case: %d", index)
	}
}

// 0-55+
func (text *Text) CalculateLix() float32 {
	return text.WordCount/float32(len(text.Sentences)) + (text.LongWordCount*100)/text.WordCount
}

// 100-0
func (text *Text) CalculateFleschReading() float32 {
	return 206.835 - 1.015*(text.WordCount/float32(len(text.Sentences))) - 84.6*(text.SyllableCount/text.WordCount)
}

// 0-18
func (text *Text) CalculateFleschGrade() float32 {
	return 0.39*(text.WordCount/float32(len(text.Sentences))) + 11.8*(text.SyllableCount/text.WordCount) - 15.59
}

// -4.9-10+
func (text *Text) CalculateDaleChall() float32 {
	pdw := CountDifficultWordsDaleChall(text) / text.WordCount * 100
	rawScore := 0.1579*pdw + 0.0496*(text.WordCount/float32(len(text.Sentences)))
	if pdw > 0.05 {
		return 3.65365 + rawScore
	}
	return rawScore
}

//6-17
func (text *Text) CalculateGunningFog() (float32, error) {
	samples, err := SamplePassage(text, 100, 1)
	if err != nil {
		return 0.0, err
	}

	sampleText := samples[0]
	var ASL float32 = sampleText.WordCount / float32(len(sampleText.Sentences))
	var PHW float32 = (float32(CountWordsWithNSyllabes(text, 3)) / sampleText.WordCount)
	return 0.4 * (ASL + PHW), nil
}

// 108-174, 2-25+
func (text *Text) CalculateFryGraph() (float32, float32, error) {
	samples, err := SamplePassage(text, 100, 3)
	if err != nil {
		return 0.0, 0.0, err
	}

	var syllableCount float32
	var wordCount float32
	var sentenceCount float32
	for _, sample := range samples {
		for _, sentence := range sample.Sentences {
			syllableCount += sentence.SyllableCount
			wordCount += sentence.WordCount
			sentenceCount++
		}
	}

	return syllableCount / (wordCount / 100), sentenceCount / (wordCount / 100), nil
}

// 5- 18
func (text *Text) CalculateSmog() (float32, error) {
	sentences, err := sampleSenteces(text, 10, 3)
	if err != nil {
		return 0.0, err
	}
	var threeWordSyllableCount float32
	for _, sentenceData := range sentences {
		sentence := text.GetSentenceText(sentenceData)
		for _, word := range strings.Fields(sentence) {
			if syllables.In(word) > 2 {
				threeWordSyllableCount++
			}
		}
	}

	return float32(3 + math.Sqrt(float64(threeWordSyllableCount))), nil
}

// Recommended 9-10
func (text *Text) CalculateForcast() (float32, error) {
	samples, err := SamplePassage(text, 150, 1)
	if err != nil {
		return 0.0, err
	}

	var oneWordSyllableCount float32
	sampledText := (*samples[0])
	for _, sentence := range sampledText.Sentences {
		for _, word := range strings.Fields(sampledText.GetSentenceText(sentence)) {
			if syllables.In(word) > 2 {
				oneWordSyllableCount++
			}
		}
	}

	return 20.0 - (oneWordSyllableCount / 10), nil
}

// Syntactic Density Score
// 1-4
func (text *Text) CalculateSDS() float64 {
	return 0.0
}
