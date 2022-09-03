package requests

import (
	"encoding/json"
	"log"

	"github.com/zrwaite/google-cloud-ddns/models"
)

func GetRecord(domain string, params *models.Params) *models.DNSRecord {
	url := "https://dns.googleapis.com/dns/v1beta2/projects/insomnizac/managedZones/insomnizac/rrsets/" + domain + "./A"
	resp, err := AuthorizedGetRequest(url, "Bearer "+params.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	var record models.DNSRecord

	if resp.StatusCode == 401 {
		RefreshAccess(params)
		return GetRecord(domain, params)
	} else if resp.StatusCode != 200 {
		log.Fatal("Error getting record " + domain + ": " + resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		log.Fatalln(err)
	}
	return &record
}
