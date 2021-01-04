package flush

import "strings"

var (
	// checkers is a map where the key is a stack e.g php, backend, mysql, etc.
	// the value is a function that takes a tweet string, and returns if the key (i.e the stack) was found in the tweet.
	// the second return value is a string slice of 'synonyms' or similar stack. The idea is that if we run a check for javascript,
	// it should tell the caller to use it's result for say, 'js'.
	checkers = make(map[string]func(string)(bool, []string))

	stacks = []string {"php", "javascript"} // load this from db?
)

func extractStacks(tweet string) map[string]bool {
	initCheckers()
	matchedStacks := map[string]bool{}
	for _, stack := range stacks {
		if cond, ok := checkers[stack]; ok {
			if found, synonyms := cond(tweet); found {
				matchedStacks[stack] = true
				for _, alt := range synonyms {
					matchedStacks[alt] = true
				}
			}
		} else {
			// we don't have a way to check for this stack yet, so do a straight up string match
			if strings.Contains(tweet, stack) || strings.Contains(tweet, "#"+stack) {
				matchedStacks[stack] = true
			}
		}
	}

	return matchedStacks
}

func initCheckers() {
	checkers["php"] = func(tweet string) (bool, []string) {
		var alsoMatches []string

		found := strings.Contains(tweet, "php") || strings.Contains(tweet, "#php")
		return found, alsoMatches
	}

	checkers["javascript"] = func(tweet string) (bool, []string) {
		alsoMatches := []string{"js"}
		found := strings.Contains(tweet, "js") ||
			strings.Contains(tweet, "#js") ||
			strings.Contains(tweet, "javascript") ||
			strings.Contains(tweet, "#javascript")
		return found, alsoMatches
	}
}