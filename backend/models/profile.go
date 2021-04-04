package models

import "time"

// profile private
type ProfileResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	Profile Profile `json:"Content"`
}

type Profile struct {
	Username       string    `json:"username"`
	DateRegistered time.Time `json:"date_registered"`
	RecoverBy      string    `json:"recover_by"`
	Contact        string    `json:"contact"`
	DateConnected  time.Time `json:"date_connected"`
	AvatarName     string    `json:"avatar_name"`
	PlayingTitle   string    `json:"playing_title"`
	PlayingArtist  string    `json:"playing_artist"`
	PlayingAlbum   string    `json:"playing_album"`
	PlayingURL     string    `json:"playing_url"`
}

// profile public
type ProfilePublicResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	ProfilePublic ProfilePublic  `json:"Content"`
}

type ProfilePublic struct {
	Username       string    `json:"username"`
	DateRegistered time.Time `json:"date_registered"`
	DateConnected  time.Time `json:"date_connected"`
	AvatarName     string    `json:"avatar_name"`
	PlayingTitle   string    `json:"playing_title"`
	PlayingArtist  string    `json:"playing_artist"`
	PlayingAlbum   string    `json:"playing_album"`
	PlayingURL     string    `json:"playing_url"`
}


// signin
type SignInResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	JWT string `json:"Content"`
}

type SignInCheckResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	Username string `json:"Content"`
}

// avatars
type AvatarListResponse struct {
	Status  bool     `json:"Status"`
	Message string   `json:"Message"`
	AvatarList []string `json:"Content"`
}

// check username
type UsernameCheckResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	IsUsed bool   `json:"Content"`
}

// sign up
type SignUpResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
}
