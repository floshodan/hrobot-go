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

type SubnetMac struct {
	Mac struct {
		IP          string `json:"ip"`
		Mask        int    `json:"mask"`
		Mac         string `json:"mac"`
		PossibleMac interface {
		} `json:"possible_mac"`
	} `json:"mac"`
}

type SubnetCancellation struct {
	Cancellation struct {
		IP                       string      `json:"ip"`
		Mask                     int         `json:"mask"`
		ServerNumber             int         `json:"server_number"`
		EarliestCancellationDate string      `json:"earliest_cancellation_date"`
		Cancelled                bool        `json:"cancelled"`
		CancellationDate         interface{} `json:"cancellation-date"`
	} `json:"cancellation"`
}
