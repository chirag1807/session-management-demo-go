package response

import "time"

type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       *string   `json:"image"`
	Topic       int64     `json:"topic"`
	Author      int64     `json:"author"`
	Likes       int       `json:"likes"`
	Views       int       `json:"views"`
	PublishedAt time.Time `json:"publishedat"`
}

type ArticleResponse struct {
	Article []Article `json:"article"`
}
