CREATE TABLE IF NOT EXISTS tweets(
    tweet_id VARCHAR(20) PRIMARY KEY,
    tweet_text TEXT NOT NULL,
    tweet_timestamp TIMESTAMP WITH TIME ZONE,
    /*
     using `saved_at` to represent when we saved the tweet,
     since twitter already maintains a `created_at` field for when the tweet was created
     */
    saved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);