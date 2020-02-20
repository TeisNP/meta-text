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

// WordsToCount returns a map of every unique word (case not included) to its count
func WordsToCount(text string) map[string]int {
	cleanedText := specialRegex.ReplaceAllString(strings.ToLower(text), " ")
	words := strings.Fields(cleanedText)
	wordToCount := make(map[string]int)
	for _, word := range words {
		wordToCount[word]++
	}

	return wordToCount
}

// CountWords returns the number of words of at least length n
func CountWords(text string, n int) int {
	var total int
	wordCount := WordsToCount(text)
	for word, count := range wordCount {
		if len(word) >= n {
			total += count
		}
	}

	return total
}

// AverageWordLength returns the average length of all words in the text
func AverageWordLength(text string) int {
	var totalWordLen int
	var totalWords int
	wordCount := WordsToCount(text)
	for word, count := range wordCount {
		totalWordLen += len(word) * count
		totalWords += count
	}

	return totalWordLen / totalWords
}

// CountSentences returns number of sentences in a text, split by [.!?] \n
// Sentence containing only special characters are removed
func CountSentences(text string, n int) int {
	sentences := periodRegex.Split(text, -1)

	var total int
	for _, sentence := range sentences {
		cleanedSentence := specialRegex.ReplaceAllString(sentence, "")
		if cleanedSentence != "" {
			total++
		}
	}
	return total
}

// GetSentencesOrdered returns a slice of sentences \n
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
