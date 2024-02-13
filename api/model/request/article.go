package request

type Article struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Image       string `json:"image,omitempty"`
	Topic       int64  `json:"topic,omitempty"`
	Author      int64  `json:"author,omitempty"`
	PublishedAt int64  `json:"publishedat,omitempty"`
}
