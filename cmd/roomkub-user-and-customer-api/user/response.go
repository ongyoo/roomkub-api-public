package user

import "time"

type CreateUserResponse struct {
	Id string `json:"id" example:"004"`
}

type LoginUserResponse struct {
	Token *string `json:"token"`
}

type GetUserItemList struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	NickName     string    `json:"nick_name"`
	ThumbnailURl string    `json:"thumbnail_url"`
	IsActive     bool      `json:"is_active"`
	IsBaned      bool      `json:"is_baned"`
	IsVerify     bool      `json:"is_verify"`
	Slug         string    `json:"slug"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}
