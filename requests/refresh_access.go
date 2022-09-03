package requests

import (
	"log"

	"github.com/zrwaite/google-cloud-ddns/models"
)

func RefreshAccess(params *models.Params) {
	if params.RefreshAttempted {
		log.Fatal("Error: Refresh failed multiple times")
	}
	params.RefreshAttempted = true
	//refresh
}
