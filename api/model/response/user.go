package response

type User struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Bio      string     `json:"bio"`
	Email    string     `json:"email"`
	Password string     `json:"password,omitempty"`
	Image    *string 	`json:"image,omitempty"`
	IsAdmin  bool       `json:"isadmin,omitempty"`
}