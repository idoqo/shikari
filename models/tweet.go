package models

type Tweet struct {
	TweetId string `json:"id"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
	SavedAt string `json:"saved_at"`
}

type SearchHits struct {
	Data []Tweet
}
