package models

type Params struct {
	ClientID         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	AccessToken      string   `json:"access_token"`
	RefreshToken     string   `json:"refresh_token"`
	ContactEmail     string   `json:"contact_email"`
	FromEmail        string   `json:"from_email"`
	SendgridAPIKey   string   `json:"sendgrid_api_key"`
	Domains          []string `json:"domains"`
	RefreshAttempted bool
}
