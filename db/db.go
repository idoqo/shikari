package db

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/idoko/shikari/models"
	"log"

	_ "github.com/lib/pq"
)

var (
	ErrNoRecord = errors.New("no matching row was found")
)
type Database struct {
	Conn *sql.DB
	Error error
}

func Connect(host, username, password, database, sslmode string, port int) Database {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port = %d user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		db.Error = err
	}
	db.Conn = conn
	log.Println("Database connection established")
	return db
}

func (db Database) SaveTweet(tweet *models.Tweet) error {
	var savedAt string

	query := "INSERT INTO tweets (tweet_id, tweet_text, tweet_timestamp) VALUES ($1, $2, $3) RETURNING saved_at"
	err := db.Conn.QueryRow(query, tweet.TweetId, tweet.Text, tweet.CreatedAt).Scan(&savedAt)
	if err != nil {
		return err
	}
	tweet.SavedAt = savedAt
	return nil
}

func (db Database) SaveTweetTag(tag *models.Tag, tweetId string) error {
	query := "INSERT INTO tweet_tags(twitter_tweet_id, tag_id) VALUES($1, $2)"
	_, err := db.Conn.Exec(query, tweetId, tag.ID)
	return err
}