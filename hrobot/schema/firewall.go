package schema

type Firewall struct {
	Firewall struct {
		ServerIP     string `json:"server_ip"`
		ServerNumber int    `json:"server_number"`
		Status       string `json:"status"`
		Filter_IPv6  bool   `json:"filter_ipv6"`
		WhitelistHos bool   `json:"whitelist_hos"`
		Port         string `json:"port"`
		Rules        struct {
			Input []struct {
				IPVersion string      `json:"ip_version"`
				Name      string      `json:"name"`
				DstIP     interface{} `json:"dst_ip"`
				SrcIP     string      `json:"src_ip"`
				DstPort   string      `json:"dst_port"`
				SrcPort   interface{} `json:"src_port"`
				Protocol  interface{} `json:"protocol"`
				TCPFlags  interface{} `json:"tcp_flags"`
				Action    string      `json:"action"`
			} `json:"input"`
			Output []struct {
				IPVersion string      `json:"ip_version"`
				Name      string      `json:"name"`
				DstIP     interface{} `json:"dst_ip"`
				SrcIP     string      `json:"src_ip"`
				DstPort   string      `json:"dst_port"`
				SrcPort   interface{} `json:"src_port"`
				Protocol  interface{} `json:"protocol"`
				TCPFlags  interface{} `json:"tcp_flags"`
				Action    string      `json:"action"`
			} `json:"output"`
		} `json:"rules"`
	} `json:"firewall"`
}

type FirewallTemplateList []struct {
	FirewallTemplate struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Filter_IPv6  bool   `json:"filter_ipv6"`
		WhitelistHos bool   `json:"whitelist_hos"`
		IsDefault    bool   `json:"is_default"`
	} `json:"firewall_template"`
}

type FirewallTemplate struct {
	FirewallTemplate struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Filter_IPv6  bool   `json:"filter_ipv6"`
		WhitelistHos bool   `json:"whitelist_hos"`
		IsDefault    bool   `json:"is_default"`
	} `json:"firewall_template"`
}

type FirewallTemplateWithRules struct {
	FirewallTemplate struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Filter_IPv6  bool   `json:"filter_ipv6"`
		WhitelistHos bool   `json:"whitelist_hos"`
		IsDefault    bool   `json:"is_default"`
		Rules        struct {
			Input []struct {
				IPVersion    string      `json:"ip_version"`
				Name         string      `json:"name"`
				DstIP        interface{} `json:"dst_ip"`
				SrcIP        string      `json:"src_ip"`
				DstPort      string      `json:"dst_port"`
				SrcPort      interface{} `json:"src_port"`
				Protocol     interface{} `json:"protocol"`
				TCPFlags     interface{} `json:"tcp_flags"`
				PacketLength interface{} `json:"packet_length"`
				Action       string      `json:"action"`
			} `json:"input"`
			Output []struct {
				IPVersion    string      `json:"ip_version"`
				Name         string      `json:"name"`
				DstIP        interface{} `json:"dst_ip"`
				SrcIP        string      `json:"src_ip"`
				DstPort      string      `json:"dst_port"`
				SrcPort      interface{} `json:"src_port"`
				Protocol     interface{} `json:"protocol"`
				TCPFlags     interface{} `json:"tcp_flags"`
				PacketLength interface{} `json:"packet_length"`
				Action       string      `json:"action"`
			} `json:"output"`
		} `json:"rules"`
	} `json:"firewall_template"`
}
