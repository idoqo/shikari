package flush

import "testing"

func TestTweetStacks(t *testing.T) {
	tweet := "✨ Wunder Mobility is now looking for a Senior Backend PHP Engineer - Infrastructure (m/f/d) #rest #php #javascript #api #amazonwebservices #remotejobs #remotework #JobSearch #RemoteWork #workfromhome  👉 https://t.co/gDAamik7Ly"
	tags := extractStacks(tweet)
	if _, ok := tags["php"]; !ok {
		t.Errorf("Expected to find PHP in stacks for %s but didn't\n", tweet)
	}
	if _, ok := tags["js"]; !ok {
		t.Errorf("expected to find javascript in stacks for %s but didn't\n", tweet)
	}
}