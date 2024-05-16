package services

import (
	"database/sql"
	"errors"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

type ArticleService struct {
	DB *sql.DB
}

func (s *ArticleService) GetAllArticles() ([]Article, error) {
	rows, err := s.DB.Query("select ID, TITLE, CONTENT from ARTICLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (s *ArticleService) GetArticleById(id int) (*Article, error) {
	var article Article

	err := s.DB.QueryRow("select ID, TITLE, CONTENT from ARTICLES where ID = ?", id).Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	return &article, nil
}
