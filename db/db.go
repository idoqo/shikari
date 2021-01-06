package db

import (
	"database/sql"
	"fmt"
	"gitlab.com/idoko/shikari/models"
	"log"

	_ "github.com/lib/pq"
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
	var id int
	var savedAt string

	query := "INSERT INTO tweets (tweet_id, tweet_text, tweet_timestamp) VALUES ($1, $2, $3) RETURNING id, saved_at"
	err := db.Conn.QueryRow(query, tweet.TweetId, tweet.Text, tweet.CreatedAt).Scan(&id, &savedAt)
	if err != nil {
		return err
	}
	tweet.ID = id
	tweet.SavedAt = savedAt
	return nil
}