package metatext

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

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
	Sentences       map[int]*Sentence
}

type Sentence struct {
	SyllableCount   float32
	WordCount       float32
	TotalWordLength float32
	LongWordCount   float32
	Words           []*Word
}

type Word struct {
	Word          string
	SyllableCount float32
}

func AnalyseText(text string) (*Text, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	sentences := doc.Sentences()
	textData := Text{
		Sentences: make(map[int]*Sentence),
	}

	for _, sentence := range sentences {
		cleanedSentence := CleanText(sentence.Text)
		if cleanedSentence != "" {
			textData.AppendSentence(analyseSentence(cleanedSentence))
		}
	}

	return &textData, nil
}

func (text *Text) AppendSentence(sentence *Sentence) {
	text.SyllableCount += sentence.SyllableCount
	text.WordCount += sentence.WordCount
	text.TotalWordLength += sentence.TotalWordLength
	text.LongWordCount += sentence.LongWordCount
	text.Sentences[int(text.WordCount)] = sentence

}

func analyseSentence(sentence string) *Sentence {
	words := strings.Fields(sentence)
	sentenceData := Sentence{
		WordCount: float32(len(words)),
		Words:     make([]*Word, len(words)),
	}

	for i, word := range words {
		wordData := analyseWord(word)
		sentenceData.SyllableCount += wordData.SyllableCount
		sentenceData.TotalWordLength += float32(len(word))
		if len(word) > 6 {
			sentenceData.LongWordCount++
		}
		sentenceData.Words[i] = wordData
	}

	return &sentenceData
}

func analyseWord(word string) *Word {
	wordData := Word{
		Word:          word,
		SyllableCount: float32(syllables.In(word)),
	}

	return &wordData
}

func CountWordsWithNSyllabes(sentences map[int]*Sentence, n float32) float32 {
	var total float32
	for _, sentence := range sentences {
		for _, data := range sentence.Words {
			if data.SyllableCount >= n {
				total += data.SyllableCount
			}
		}

	}

	return total
}

func SamplePassage(text *Text, length int, number int) ([]*Text, error) {
	textSplitLength := int(text.WordCount) / number

	if textSplitLength < length {
		return nil, errors.New("The text is too short for the split")
	}

	//rand.Seed(time.Now().UnixNano())
	passages := make([]*Text, number)
	currentIndex := 0
	for i := 0; i < number; i++ {
		passage := Text{
			Sentences: make(map[int]*Sentence),
		}

		start := rand.Intn(textSplitLength-length) + currentIndex
		end := start + length
		for index, sentence := range text.Sentences {
			if start >= index {
				continue
			}
			if start < index {
				passage.AppendSentence(sentence)
				continue
			}
			if end <= index {
				passage.AppendSentence(sentence)
				break
			}
		}

		passages[i] = &passage
		currentIndex += textSplitLength
	}

	return passages, nil
}
