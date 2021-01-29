package flush

import "testing"

func TestTweetStacks(t *testing.T) {
	tweet := "✨ Wunder Mobility is now looking for a Senior Backend PHP Engineer - Infrastructure (m/f/d) #rest #php #javascript #api #amazonwebservices #remotejobs #remotework #JobSearch #RemoteWork #workfromhome  👉 https://t.co/gDAamik7Ly"
	var hasPHP, hasJS bool
	for _, tag := range extractStacks(tweet) {
		if tag.Tag == "php" {
			hasPHP = true
		}
		if tag.Tag == "javascript" {
			hasJS = true
		}
	}
	if !hasPHP {
		t.Errorf("Expected to find PHP in stacks for %s but didn't\n", tweet)
	}
	if !hasJS {
		t.Errorf("expected to find javascript in stacks for %s but didn't\n", tweet)
	}
}