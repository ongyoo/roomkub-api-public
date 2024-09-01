package businessChannel

import "time"

type GetBusinessChannelItem struct {
	ID               string    `json:"id"`
	ChannelID        string    `json:"channel_id"`
	Name             string    `json:"name"`
	LogoThumbnailURl string    `json:"logo_thumbnail_url"`
	IsActive         bool      `json:"is_active"`
	IsBaned          bool      `json:"is_ban"`
	IsSeparate       bool      `json:"is_separate"`
	CreateAt         time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}

type GetRootBusinessChannelItem struct {
	ID               string    `json:"id"`
	ChannelID        string    `json:"channel_id"`
	MyPackageID      string    `json:"my_package_id"`
	Name             string    `json:"name"`
	LogoThumbnailURl string    `json:"logo_thumbnail_url"`
	IsActive         bool      `json:"is_active"`
	IsBaned          bool      `json:"is_ban"`
	CreateAt         time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}

type GetBusinessChannelIteById struct {
	ID               string    `json:"id"`
	ChannelID        string    `json:"channel_id"`
	Name             string    `json:"name"`
	LogoThumbnailURl string    `json:"logo_thumbnail_url"`
	IsActive         bool      `json:"is_active"`
	IsBaned          bool      `json:"is_ban"`
	IsSeparate       bool      `json:"is_separate"`
	CreateAt         time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}
