package schema

type RDNSList []struct {
	Rdns struct {
		IP  string `json:"ip"`
		Ptr string `json:"ptr"`
	} `json:"rdns"`
}

type RDNS struct {
	Rdns struct {
		IP  string `json:"ip"`
		Ptr string `json:"ptr"`
	} `json:"rdns"`
}
