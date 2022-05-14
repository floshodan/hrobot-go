package schema

type WOL struct {
	Wol struct {
		ServerIP      string `json:"server_ip"`
		ServerIpv6Net string `json:"server_ipv6_net"`
		ServerNumber  int    `json:"server_number"`
	} `json:"wol"`
}
