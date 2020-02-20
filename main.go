package metatext

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	periodRegex  *regexp.Regexp
	specialRegex *regexp.Regexp
	err          error
)

func init() {
	periodRegex, err = regexp.Compile("[.!?]")
	if err != nil {
		fmt.Println(err)
	}

	specialRegex, err = regexp.Compile(`[^\w]`)
	if err != nil {
		fmt.Println(err)
	}
}

// CountWords returns a map of every unique word (case not included) to its count
func CountWords(text string) map[string]int {
	cleanedText := specialRegex.ReplaceAllString(strings.ToLower(text), " ")
	words := strings.Fields(cleanedText)
	wordToCount := make(map[string]int)
	for _, word := range words {
		wordToCount[word]++
	}

	return wordToCount
}

// GetSentencesOrdered returns a slice of sentences
// Sentence containing only special characters are removed
func GetSentencesOrdered(text string) []string {
	sentences := periodRegex.Split(text, -1)

	correctSentences := sentences[:0]
	for _, sentence := range sentences {
		cleanedSentence := specialRegex.ReplaceAllString(sentence, "")
		if cleanedSentence != "" {
			correctSentences = append(correctSentences, strings.TrimSpace(sentence))
		}
	}
	return correctSentences
}
