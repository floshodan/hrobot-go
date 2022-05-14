package schema

type IPList []struct {
	IP struct {
		IP              string      `json:"ip"`
		ServerIP        string      `json:"server_ip"`
		ServerNumber    int         `json:"server_number"`
		Locked          bool        `json:"locked"`
		SeparateMac     interface{} `json:"separate_mac"`
		TrafficWarnings bool        `json:"traffic_warnings"`
		TrafficHourly   int         `json:"traffic_hourly"`
		TrafficDaily    int         `json:"traffic_daily"`
		TrafficMonthly  int         `json:"traffic_monthly"`
	} `json:"ip"`
}

type IP struct {
	IP struct {
		IP              string      `json:"ip"`
		ServerIP        string      `json:"server_ip"`
		ServerNumber    int         `json:"server_number"`
		Locked          bool        `json:"locked"`
		SeparateMac     interface{} `json:"separate_mac"`
		TrafficWarnings bool        `json:"traffic_warnings"`
		TrafficHourly   int         `json:"traffic_hourly"`
		TrafficDaily    int         `json:"traffic_daily"`
		TrafficMonthly  int         `json:"traffic_monthly"`
	} `json:"ip"`
}

type MAC struct {
	Mac struct {
		IP  string `json:"ip"`
		Mac string `json:"mac"`
	} `json:"mac"`
}
