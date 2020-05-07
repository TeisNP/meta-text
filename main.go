package metatext

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/teisnp/syllables"
	"gopkg.in/jdkato/prose.v2"
)

var (
	letterRegex *regexp.Regexp
	spaceRegex  *regexp.Regexp

	err error
)

func init() {
	letterRegex, err = regexp.Compile(`[^A-ZÆØÅa-zæøå ][s]?|[^\w][^AIOØÅaioØÅ]{1}[^\w]{1}`)
	if err != nil {
		fmt.Println(err)
	}

	spaceRegex, err = regexp.Compile(`[ ]{2,}`)
	if err != nil {
		fmt.Println(err)
	}
}

func CleanText(text string) string {
	cleanedText := spaceRegex.ReplaceAllString(letterRegex.ReplaceAllString(text, ""), " ")
	return strings.ToLower(cleanedText)
}

type Text struct {
	SyllableCount   float32
	WordCount       float32
	TotalWordLength float32
	LongWordCount   float32
	Sentences       []*Sentence
	Text            string
}

type Sentence struct {
	SyllableCount   float32
	WordCount       float32
	TotalWordLength float32
	LongWordCount   float32
	Start           int
	End             int
}

func (text *Text) GetSentenceText(sentence *Sentence) string {
	//fmt.Printf("Start: %d, End: %d", sentence.Start, sentence.End)
	return text.Text[sentence.Start:sentence.End]
}

func AnalyseText(text string) (*Text, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	textData := Text{}
	var cleanedText bytes.Buffer

	sentences := doc.Sentences()
	for _, sentence := range sentences {
		cleanedSentence := CleanText(sentence.Text)
		if cleanedSentence != "" {
			sentenceData := analyseSentence(cleanedSentence, cleanedText.Len())

			textData.SyllableCount += sentenceData.SyllableCount
			textData.WordCount += sentenceData.WordCount
			textData.TotalWordLength += sentenceData.TotalWordLength
			textData.LongWordCount += sentenceData.LongWordCount
			textData.Sentences = append(textData.Sentences, sentenceData)
			cleanedText.WriteString(cleanedSentence + " ")
		}
	}
	textData.Text = cleanedText.String()

	return &textData, nil
}

func analyseSentence(sentence string, start int) *Sentence {
	words := strings.Fields(sentence)
	sentenceData := Sentence{
		Start:         start,
		End:           start + len(sentence) + 1,
		WordCount:     float32(len(words)),
		SyllableCount: float32(syllables.In(sentence)),
	}

	for _, word := range words {
		sentenceData.TotalWordLength += float32(len(word))
		if len(word) > 6 {
			sentenceData.LongWordCount++
		}
	}

	return &sentenceData
}

func CountWordsWithNSyllabes(text *Text, n int) int {
	var total int
	for _, sentenceData := range text.Sentences {
		sentence := text.GetSentenceText(sentenceData)
		for _, word := range strings.Fields(sentence) {
			syllables := syllables.In(word)
			if syllables >= n {
				total += syllables
			}
		}

	}

	return total
}

func sampleSenteces(text *Text, length int, number int) ([]*Sentence, error) {
	sentenceSplit := len(text.Sentences) / number
	if sentenceSplit <= length {
		return nil, errors.New("The text is too short for the split")
	}

	sentences := make([]*Sentence, number*length)
	currentIndex := 0
	for i := 0; i < number; i++ {
		start := rand.Intn(sentenceSplit-length) + i*sentenceSplit
		for j := start; j < start+length; j++ {
			sentences[currentIndex] = text.Sentences[j]
			currentIndex++
		}
	}

	return sentences, nil
}

func SamplePassage(text *Text, length int, number int) ([]*Text, error) {
	textSplitLength := int(text.WordCount) / number

	if textSplitLength <= length {
		return nil, errors.New("The text is too short for the split")
	}

	rand.Seed(time.Now().UnixNano())
	passages := make([]*Text, number)
	for i := 0; i < number; i++ {
		startWordindex := rand.Intn(textSplitLength-length) + i*textSplitLength

		start, err := text.wordIndexToStringIndex(startWordindex)
		end, err := text.wordIndexToStringIndex(startWordindex + length)

		passageText, err := AnalyseText(text.Text[start:end])
		if err != nil {
			return nil, err
		}
		passages[i] = passageText
	}

	return passages, nil
}

func (text *Text) wordIndexToStringIndex(index int) (int, error) {
	currentWordIndex := 0
	for _, sentence := range text.Sentences {
		currentWordIndex += int(sentence.WordCount)

		if index <= currentWordIndex {
			return text.stringOndexOfWord(sentence, currentWordIndex-index), nil
		}
	}
	return 0, errors.New("Word index not found")
}

func (text *Text) stringOndexOfWord(sentence *Sentence, index int) int {
	sentenceString := text.GetSentenceText(sentence)
	senteceWords := strings.Fields(sentenceString)
	stringIndex := 0
	for i := 0; i < index; i++ {
		stringIndex += len(senteceWords[i]) + 1
	}

	return stringIndex + sentence.Start
}
