package main

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/zrwaite/google-cloud-ddns/models"
	"github.com/zrwaite/google-cloud-ddns/requests"
)

func main() {
	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    paramsFilePath := path.Join(exPath, "params.json")
	file, err := os.Open(paramsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(file)
	var params models.Params
	json.Unmarshal(byteValue, &params)
	defer file.Close()

	initialRecord := requests.GetRecord(params.Domains[0], &params)
	currentIP := requests.GetIP()

	if initialRecord.Rrdatas[0] != currentIP {
		records := []models.DNSRecord{}
		for _, domain := range params.Domains {
			records = append(records, *requests.GetRecord(domain, &params))
		}
		fmt.Println("newip: <" + currentIP + ">")
		requests.PatchRecords(records, currentIP, &params)
	}
	if params.RefreshAttempted {
		fmt.Println("Updating access token")
		params.RefreshAttempted = false
		content, err := json.Marshal(params)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(paramsFilePath, content, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
