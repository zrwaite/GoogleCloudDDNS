package models

type Params struct {
	ClientID         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	AccessToken      string   `json:"access_token"`
	RefreshToken     string   `json:"refresh_token"`
	Domains          []string `json:"domains"`
	RefreshAttempted bool
}
