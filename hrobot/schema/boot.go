package schema

type BootLinux struct {
	Linux struct {
		ServerIP      string        `json:"server_ip"`
		ServerIpv6Net string        `json:"server_ipv6_net"`
		ServerNumber  int           `json:"server_number"`
		Dist          string        `json:"dist"`
		Arch          int           `json:"arch"`
		Lang          string        `json:"lang"`
		Active        bool          `json:"active"`
		Password      string        `json:"password"`
		AuthorizedKey []interface{} `json:"authorized_key"`
		HostKey       []interface{} `json:"host_key"`
	} `json:"linux"`
}

type BootRescue struct {
	Rescue struct {
		ServerIP      string        `json:"server_ip"`
		ServerIpv6Net string        `json:"server_ipv6_net"`
		ServerNumber  int           `json:"server_number"`
		Os            string        `json:"os"`
		Arch          int           `json:"arch"`
		Active        bool          `json:"active"`
		Password      string        `json:"password"`
		AuthorizedKey []interface{} `json:"authorized_key"`
		HostKey       []interface{} `json:"host_key"`
	} `json:"rescue"`
}
