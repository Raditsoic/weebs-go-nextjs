// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Anime struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Genre       *Genre `json:"genre"`
}

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Manga struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Genre       *Genre `json:"genre"`
}

type Mutation struct {
}

type NewAnime struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	GenreID     string `json:"genreId"`
}

type NewManga struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	GenreID     string `json:"genreId"`
}

type NewReview struct {
	ContentID   *string `json:"contentID,omitempty"`
	ContentType string  `json:"contentType"`
	Comment     string  `json:"comment"`
	Rating      float64 `json:"rating"`
}

type Query struct {
}

type RegisterUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Review struct {
	ID          string  `json:"id"`
	UserID      string  `json:"userID"`
	ContentID   string  `json:"contentID"`
	ContentType string  `json:"contentType"`
	Rating      float64 `json:"rating"`
	Comment     string  `json:"comment"`
	CreatedAt   string  `json:"createdAt"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`
	Hash     string `json:"hash"`
	Role     string `json:"role"`
}
