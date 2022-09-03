package requests

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetIP() string {
	url := "https://icanhazip.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(string(b))
}
