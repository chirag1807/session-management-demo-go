package response

type Topic struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TopicResponse struct {
	Topics  []Topic `json:"topics"`
}
