package database

import (
	"cloud-sek/globals"
	"cloud-sek/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db = globals.DbConn

type PostRepository struct{}

func InitDbConn() {
	loadDbConn()
}

func loadDbConn() {
	config := globals.Config.Db
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Pass, config.Database, config.SslMode)
	db, err := sql.Open(config.Driver, connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	Db = db
}

func (p *PostRepository) InsertPost(query string, post models.Post) error {
	_, err := Db.Exec(query, post.Title, post.Description)
	if err != nil {
		fmt.Print("Error inserting post: ", err)
		return err
	}
	return nil
}

func (p *PostRepository) GetPostById(query, id string) (models.Post, error) {
	row := Db.QueryRow(query, id)
	var post models.Post
	err := row.Scan(&post.ID, &post.Title, &post.Description)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p *PostRepository) GetCommentsByPostID(query, postID string) ([]models.Comment, error) {
	comments := []models.Comment{}
	rows, err := Db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		comment.PostID = postID
		if err := rows.Scan(&comment.ID, &comment.Author, &comment.Message); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (p *PostRepository) InsertComment(query string, comment models.Comment) (string, error) {
	var commentId string
	err := Db.QueryRow(query, comment.PostID, comment.Author, comment.Message).Scan(&commentId)
	if err != nil {
		fmt.Print("Error inserting comment: ", err)
		return "", err
	}
	return commentId, nil
}
