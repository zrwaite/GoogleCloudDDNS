package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zrwaite/google-cloud-ddns/auth"
	"github.com/zrwaite/google-cloud-ddns/mail"
	"github.com/zrwaite/google-cloud-ddns/models"
)

func RefreshAccess(params *models.Params) {
	if params.RefreshAttempted {
		mail.ErrorMessage("Error: Refresh failed multiple times", params)
		log.Fatal("Error: Refresh failed multiple times")
	}
	params.RefreshAttempted = true

	token, tokenSuccess := auth.EncodeToken(params)
	if !tokenSuccess {
		mail.ErrorMessage("Error: Refresh failed", params)
		log.Fatal("Error: Refresh failed")
	}
	refresh := models.Refresh{
		Assertion: token,
		GrantType: "urn:ietf:params:oauth:grant-type:jwt-bearer",
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
		fmt.Println("Error refreshing token: " + resp.Status)
		log.Fatal(resp)
	}
	var access models.Access
	err = json.NewDecoder(resp.Body).Decode(&access)
	if err != nil {
		log.Fatalln(err)
	}
	params.AccessToken = access.AccessToken
}
