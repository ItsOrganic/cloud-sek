package models

import (
	"sync"
	"time"
)

type Config struct {
	Db DBConfig `yaml:"db"`
}

type DBConfig struct {
	User     string `yaml:"user-name"`
	Pass     string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	SslMode  string `yaml:"ssl-mode"`
}

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Comments    []Comment `json:"comments"`
	Timestamp   time.Time `json:"timestamp"`
}

type PostCache struct {
	MuRW      sync.RWMutex
	PostCache map[string]*Post
}
