package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/zrwaite/google-cloud-ddns/models"
	"github.com/zrwaite/google-cloud-ddns/requests"
)

func main() {
	file, err := os.Open("params.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(file)
	var params models.Params
	json.Unmarshal(byteValue, &params)
	defer file.Close()

	initialRecord := requests.GetRecord(params.Domains[0], &params)
	currentIP := requests.GetIP()
	records := []models.DNSRecord{}
	for _, domain := range params.Domains {
		records = append(records, *requests.GetRecord(domain, &params))
	}

	if initialRecord.Rrdatas[0] != currentIP {
		fmt.Println("newip: <" + currentIP + ">")
		requests.PatchRecords(records, currentIP, &params)
	} else {
		fmt.Println("IPs match, no update needed")
	}
}
