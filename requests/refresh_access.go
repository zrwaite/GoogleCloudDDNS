package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/zrwaite/google-cloud-ddns/models"
)

func RefreshAccess(params *models.Params) {
	if params.RefreshAttempted {
		log.Fatal("Error: Refresh failed multiple times")
	}
	params.RefreshAttempted = true

	refresh := models.Refresh{
		RefreshToken: params.RefreshToken,
		ClientID:     params.ClientID,
		ClientSecret: params.ClientSecret,
		GrantType:    "refresh_token",
	}

	json_data, err := json.Marshal(refresh)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://oauth2.googleapis.com/token"
	body := bytes.NewBuffer(json_data)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error refreshing token: " + resp.Status)
	}
	var access models.Access
	err = json.NewDecoder(resp.Body).Decode(&access)
	if err != nil {
		log.Fatalln(err)
	}
	params.AccessToken = access.AccessToken
}
