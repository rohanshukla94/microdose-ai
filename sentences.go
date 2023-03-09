package main

import "strings"

type Rule interface {
	IsWordSeparator(rune rune) bool
	IsSentenceSeparator(rune rune) bool
}

type RuleDefault struct {
	wordSeparators     [21]string
	sentenceSeparators [3]string
}

// NewRule constructor retrieves a RuleDefault pointer.
func NewRule() *RuleDefault {
	return &RuleDefault{
		[21]string{" ", ",", "'", "’", "\"", ")", "(", "[", "]", "{", "}", "\"", ";", "\n", ">", "<", "%", "@", "&", "=", "#"},
		[3]string{"!", ".", "?"},
	}
}

func (r *RuleDefault) IsWordSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.wordSeparators {
		if chr == val {
			return true
		}
	}

	return r.IsSentenceSeparator(rune)
}

func (r *RuleDefault) IsSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}

func TokenizeText(rawText string, rule Rule) Text {
	return findSentences(rawText, rule)
}

func findSentences(rawText string, rule Rule) Text {
	text := Text{}

	var sentence string

	for _, chr := range rawText {
		if !rule.IsSentenceSeparator(chr) {
			sentence = sentence + string(chr)
		} else if len(sentence) > 0 {
			sentence = sentence + string(chr)

			text.Append(sentence, findWords(sentence, rule))

			sentence = ""
		}
	}

	if len(sentence) > 0 {
		text.Append(sentence, findWords(sentence, rule))
	}

	return text
}

func findWords(rawSentence string, rule Rule) (words []string) {
	words = []string{}

	var word string

	for _, chr := range rawSentence {
		if !rule.IsWordSeparator(chr) {
			word = word + string(chr)
		} else if len(word) > 0 {
			words = append(words, strings.ToLower(word))
			word = ""
		}
	}

	if len(word) > 0 {
		words = append(words, strings.ToLower(word))
	}

	return
}
