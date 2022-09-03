package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/zrwaite/google-cloud-ddns/models"
)

func PatchRecords(records []models.DNSRecord, updatedIP string, params *models.Params) {
	patch := new(models.DNSPatch)
	for _, record := range records {
		patch.Deletions = append(patch.Deletions, record)
		patch.Additions = append(patch.Additions, models.DNSRecord{
			Name:    record.Name,
			Type:    record.Type,
			Ttl:     record.Ttl,
			Rrdatas: []string{updatedIP},
		})
	}
	json_data, err := json.Marshal(patch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_data))

	url := "https://dns.googleapis.com/dns/v1beta2/projects/insomnizac/managedZones/insomnizac/changes"
	body := bytes.NewBuffer(json_data)
	fmt.Println(body)
	resp, err := AuthorizedBodyRequest(url, "POST", "Bearer "+params.AccessToken, body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error: " + resp.Status)
	}
}
