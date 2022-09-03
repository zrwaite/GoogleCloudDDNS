package models

type DNSRecord struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Ttl     int      `json:"ttl"`
	Rrdatas []string `json:"rrdatas"`
}

type DNSPatch struct {
	Additions []DNSRecord `json:"additions"`
	Deletions []DNSRecord `json:"deletions"`
}

type Refresh struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type Access struct {
	AccessToken string `json:"access_token"`
}
