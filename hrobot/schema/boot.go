package schema

type BootList struct {
	Boot struct {
		Rescue  `json:"rescue"`
		Linux   Linux       `json:"linux"`
		Vnc     Vnc         `json:"vnc"`
		Windows interface{} `json:"windows"`
		Plesk   interface{} `json:"plesk"`
		Cpanel  interface{} `json:"cpanel"`
	} `json:"boot"`
}

type RescueList struct {
	Rescue `json:"rescue"`
}

type Rescue struct {
	ServerIP      string        `json:"server_ip"`
	ServerIpv6Net string        `json:"server_ipv6_net"`
	ServerNumber  int           `json:"server_number"`
	Os            interface{}   `json:"os"`
	Arch          interface{}   `json:"arch"`
	Active        bool          `json:"active"`
	Password      interface{}   `json:"password"`
	AuthorizedKey []interface{} `json:"authorized_key"`
	HostKey       []interface{} `json:"host_key"`
	BootTime      interface{}   `json:"boot_time"`
}

type Linux struct {
	ServerIP      string        `json:"server_ip"`
	ServerIpv6Net string        `json:"server_ipv6_net"`
	ServerNumber  int           `json:"server_number"`
	Dist          []string      `json:"dist"`
	Arch          []int         `json:"arch"`
	Lang          []string      `json:"lang"`
	Active        bool          `json:"active"`
	Password      interface{}   `json:"password"`
	AuthorizedKey []interface{} `json:"authorized_key"`
	HostKey       []interface{} `json:"host_key"`
}

type Vnc struct {
	ServerIP      string      `json:"server_ip"`
	ServerIpv6Net string      `json:"server_ipv6_net"`
	ServerNumber  int         `json:"server_number"`
	Dist          []string    `json:"dist"`
	Arch          []int       `json:"arch"`
	Lang          []string    `json:"lang"`
	Active        bool        `json:"active"`
	Password      interface{} `json:"password"`
}
