package schema

type ResetList []struct {
	Reset struct {
		ServerIP        string      `json:"server_ip"`
		ServerIpv6Net   string      `json:"server_ipv6_net"`
		ServerNumber    int         `json:"server_number"`
		Type            interface{} `json:"type"`
		OperatingStatus string      `json:"operating_status"`
	} `json:"reset"`
}

type Auto struct {
	Reset struct {
		ServerIP        string   `json:"server_ip"`
		ServerIpv6Net   string   `json:"server_ipv6_net"`
		ServerNumber    int      `json:"server_number"`
		Type            []string `json:"type"`
		OperatingStatus string   `json:"operating_status"`
	} `json:"reset"`
}

type Reset struct {
	Reset struct {
		ServerIP        string      `json:"server_ip"`
		ServerIpv6Net   string      `json:"server_ipv6_net"`
		ServerNumber    int         `json:"server_number"`
		Type            interface{} `json:"type"`
		OperatingStatus string      `json:"operating_status"`
	} `json:"reset"`
}
