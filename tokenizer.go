package main

import (
	"regexp"
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

var stopwords = map[string]struct{}{
	"a": {}, "about": {}, "above": {}, "after": {}, "again": {}, "against": {}, "all": {},
	"am": {}, "an": {}, "and": {}, "any": {}, "are": {}, "aren't": {}, "as": {}, "at": {},
	"be": {}, "because": {}, "been": {}, "before": {}, "being": {}, "below": {}, "between": {},
	"both": {}, "but": {}, "by": {}, "can": {}, "can't": {}, "cannot": {}, "could": {}, "couldn't": {},
	"did": {}, "didn't": {}, "do": {}, "does": {}, "doesn't": {}, "doing": {}, "don't": {},
	"down": {}, "during": {}, "each": {}, "few": {}, "for": {}, "from": {}, "further": {},
	"had": {}, "hadn't": {}, "has": {}, "hasn't": {}, "have": {}, "haven't": {}, "having": {},
	"he": {}, "he'd": {}, "he'll": {}, "he's": {}, "her": {}, "here": {}, "here's": {},
	"hers": {}, "herself": {}, "him": {}, "himself": {}, "his": {}, "how": {}, "how's": {},
	"i": {}, "i'd": {}, "i'll": {}, "i'm": {}, "i've": {}, "if": {}, "in": {}, "into": {},
	"is": {}, "isn't": {}, "it": {}, "it's": {}, "its": {}, "itself": {}, "let's": {}, "me": {},
	"more": {}, "most": {}, "mustn't": {}, "my": {}, "myself": {}, "no": {}, "nor": {}, "not": {},
	"of": {}, "off": {}, "on": {}, "once": {}, "only": {}, "or": {}, "other": {}, "ought": {},
	"our": {}, "ours": {}, "ourselves": {}, "out": {}, "over": {}, "own": {}, "same": {},
	"shall": {}, "shan't": {}, "she": {}, "she'd": {}, "she'll": {}, "she's": {}, "should": {},
	"shouldn't": {}, "so": {}, "some": {}, "such": {}, "than": {}, "that": {}, "that's": {},
	"the": {}, "their": {}, "theirs": {}, "them": {}, "themselves": {}, "then": {}, "there": {},
	"there's": {}, "these": {}, "they": {}, "they'd": {}, "they'll": {}, "they're": {}, "they've": {},
	"this": {}, "those": {}, "through": {}, "to": {}, "too": {}, "under": {}, "until": {}, "up": {},
	"very": {}, "was": {}, "wasn't": {}, "we": {}, "we'd": {}, "we'll": {}, "we're": {}, "we've": {},
	"were": {}, "weren't": {}, "what": {}, "what's": {}, "when": {}, "when's": {}, "where": {},
	"where's": {}, "which": {}, "while": {}, "who": {}, "who's": {}, "whom": {}, "why": {},
	"why's": {}, "with": {}, "won't": {}, "would": {}, "wouldn't": {}, "you": {}, "you'd": {},
	"you'll": {}, "you're": {}, "you've": {}, "your": {}, "yours": {}, "yourself": {}, "yourselves": {},
}

func tokenize(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	// Find all matches in the text
	tokens := re.FindAllString(text, -1)
	tokens = lowerCase(tokens)
	tokens = stopWordsFilter(tokens)
	tokens = stemTokens(tokens)
	return tokens

}

func lowerCase(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}
	return tokens
}

func stopWordsFilter(tokens []string) []string {
	var result []string
	for _, token := range tokens {
		if _, exists := stopwords[token]; !exists {
			result = append(result, token)
		}
	}
	return result
}

func stemTokens(tokens []string) []string {
	var result []string
	for _, token := range tokens {
		stemmedToken := snowballeng.Stem(token, false)
		result = append(result, stemmedToken)
	}
	return result

}
