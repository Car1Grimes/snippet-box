package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (model *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (model *SnippetModel) Recent() ([]Snippet, error) {
	return nil, nil
}
