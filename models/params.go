package models

type Params struct {
	AccessToken      string   `json:"access_token"`
	RefreshToken     string   `json:"refresh_token"`
	Domains          []string `json:"domains"`
	RefreshAttempted bool
}
