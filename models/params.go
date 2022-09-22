package models

type Params struct {
	AccessToken      string   `json:"access_token"`
	ContactEmail     string   `json:"contact_email"`
	FromEmail        string   `json:"from_email"`
	SendgridAPIKey   string   `json:"sendgrid_api_key"`
	Domains          []string `json:"domains"`
	ISS              string   `json:"iss"`
	JWT_KEY          string   `json:"jwt_key"`
	RefreshAttempted bool
}
