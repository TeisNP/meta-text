package metatext

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
func (text *Text) CalculateGunningFog() float32 {
	samples, err := SamplePassage(text, 100, 1)
	if err != nil {
		return -1
	}

	sampleText := samples[0]
	var ASL float32 = sampleText.WordCount / float32(len(sampleText.Sentences))
	var PHW float32 = (float32(CountWordsWithNSyllabes(sampleText.Sentences, 3)) / sampleText.WordCount)
	return 0.4 * (ASL + PHW)
}

// 108-174, 2-25+
func (text *Text) CalculateFryGraph() (float32, float32) {
	return text.SyllableCount / (text.WordCount / 100), float32(len(text.Sentences)) / (text.WordCount / 100)
}
