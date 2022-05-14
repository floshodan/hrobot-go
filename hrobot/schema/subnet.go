package schema

type SubnetList []struct {
	Subnet struct {
		IP              string `json:"ip"`
		Mask            int    `json:"mask"`
		Gateway         string `json:"gateway"`
		ServerIP        string `json:"server_ip"`
		ServerNumber    int    `json:"server_number"`
		Failover        bool   `json:"failover"`
		Locked          bool   `json:"locked"`
		TrafficWarnings bool   `json:"traffic_warnings"`
		TrafficHourly   int    `json:"traffic_hourly"`
		TrafficDaily    int    `json:"traffic_daily"`
		TrafficMonthly  int    `json:"traffic_monthly"`
	} `json:"subnet"`
}

type Subnet struct {
	Subnet struct {
		IP              string `json:"ip"`
		Mask            int    `json:"mask"`
		Gateway         string `json:"gateway"`
		ServerIP        string `json:"server_ip"`
		ServerNumber    int    `json:"server_number"`
		Failover        bool   `json:"failover"`
		Locked          bool   `json:"locked"`
		TrafficWarnings bool   `json:"traffic_warnings"`
		TrafficHourly   int    `json:"traffic_hourly"`
		TrafficDaily    int    `json:"traffic_daily"`
		TrafficMonthly  int    `json:"traffic_monthly"`
	} `json:"subnet"`
}
