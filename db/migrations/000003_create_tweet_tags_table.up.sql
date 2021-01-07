CREATE TABLE IF NOT EXISTS tweet_tags(
    id SERIAL PRIMARY KEY,
    twitter_tweet_id VARCHAR(20) NOT NULL REFERENCES tweets(tweet_id),
    tag_id INTEGER NOT NULL REFERENCES tags(id)
);

CREATE INDEX idx_tweet_tags_twitter_tweet_id ON tweet_tags(twitter_tweet_id);