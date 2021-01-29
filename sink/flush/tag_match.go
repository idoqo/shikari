package flush

import (
	"gitlab.com/idoko/shikari/models"
	"strings"
)

var (
	// checkers is a map where the key is a stack e.g php, backend, mysql, etc.
	// the value is a function that takes a tweet string, and returns if the key (i.e the stack) was found in the tweet.
	// the second return value is a string slice of 'synonyms' or similar stack. The idea is that if we run a check for javascript,
	// it should tell the caller to use it's result for say, 'js'.
	checkers = map[string]func(string)(bool, []string) {
		"php": func(tweet string) (bool, []string) {
			var alsoMatches []string

			found := strings.Contains(tweet, "php") || strings.Contains(tweet, "#php")
			return found, alsoMatches
		},

		"javascript": func(tweet string) (bool, []string) {
			alsoMatches := []string{"js"}
			found := strings.Contains(tweet, "js") ||
				strings.Contains(tweet, "#js") ||
				strings.Contains(tweet, "javascript") ||
				strings.Contains(tweet, "#javascript")
			return found, alsoMatches
		},
	}

	stacks = []models.Tag {
		{1, "javascript"},
		{2, "cpp"},
		{3, "go"},
		{4, "php"},
		{5, "docker"},
	}
)

func extractStacks(tweet string) []models.Tag {
	var matchedStacks []models.Tag

	for _, stack := range stacks {
		if cond, ok := checkers[stack.Tag]; ok {
			if found, _ := cond(tweet); found {
				matchedStacks = append(matchedStacks, stack)
			}
		} else {
			// we don't have a way to check for this stack yet, so do a straight up string match
			if strings.Contains(tweet, stack.Tag) || strings.Contains(tweet, "#"+stack.Tag) {
				matchedStacks = append(matchedStacks, stack)
			}
		}
	}

	return matchedStacks
}