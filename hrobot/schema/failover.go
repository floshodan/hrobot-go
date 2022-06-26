package schema

type FailoverList []struct {
	Failover struct {
		IP             string `json:"ip"`
		Netmask        string `json:"netmask"`
		ServerIP       string `json:"server_ip"`
		ServerNumber   int    `json:"server_number"`
		ActiveServerIP string `json:"active_server_ip"`
	} `json:"failover"`
}

type Failover struct {
	Failover struct {
		IP             string `json:"ip"`
		Netmask        string `json:"netmask"`
		ServerIP       string `json:"server_ip"`
		ServerNumber   int    `json:"server_number"`
		ActiveServerIP string `json:"active_server_ip"`
	} `json:"failover"`
}
