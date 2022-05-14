package schema

type ServerList []struct {
	Server struct {
		ServerIP      string   `json:"server_ip"`
		ServerIpv6Net string   `json:"server_ipv6_net"`
		ServerNumber  int      `json:"server_number"`
		ServerName    string   `json:"server_name"`
		Product       string   `json:"product"`
		Dc            string   `json:"dc"`
		Traffic       string   `json:"traffic"`
		Status        string   `json:"status"`
		Cancelled     bool     `json:"cancelled"`
		PaidUntil     string   `json:"paid_until"`
		IP            []string `json:"ip"`
		Subnet        []struct {
			IP   string `json:"ip"`
			Mask string `json:"mask"`
		} `json:"subnet"`
		LinkedStoragebox interface{} `json:"linked_storagebox"`
	} `json:"server"`
}

type Server struct {
	Server struct {
		ServerIP      string   `json:"server_ip"`
		ServerIpv6Net string   `json:"server_ipv6_net"`
		ServerNumber  int      `json:"server_number"`
		ServerName    string   `json:"server_name"`
		Product       string   `json:"product"`
		Dc            string   `json:"dc"`
		Traffic       string   `json:"traffic"`
		Status        string   `json:"status"`
		Cancelled     bool     `json:"cancelled"`
		PaidUntil     string   `json:"paid_until"`
		IP            []string `json:"ip"`
		Subnet        []struct {
			IP   string `json:"ip"`
			Mask string `json:"mask"`
		} `json:"subnet"`
		LinkedStoragebox interface{} `json:"linked_storagebox"`
	} `json:"server"`
}

type SingleServer struct {
	Server struct {
		ServerIP      string   `json:"server_ip"`
		ServerIpv6Net string   `json:"server_ipv6_net"`
		ServerNumber  int      `json:"server_number"`
		ServerName    string   `json:"server_name"`
		Product       string   `json:"product"`
		Dc            string   `json:"dc"`
		Traffic       string   `json:"traffic"`
		Status        string   `json:"status"`
		Cancelled     bool     `json:"cancelled"`
		PaidUntil     string   `json:"paid_until"`
		IP            []string `json:"ip"`
		Subnet        []struct {
			IP   string `json:"ip"`
			Mask string `json:"mask"`
		} `json:"subnet"`
		Reset            bool        `json:"reset"`
		Rescue           bool        `json:"rescue"`
		Vnc              bool        `json:"vnc"`
		Windows          bool        `json:"windows"`
		Plesk            bool        `json:"plesk"`
		Cpanel           bool        `json:"cpanel"`
		Wol              bool        `json:"wol"`
		HotSwap          bool        `json:"hot_swap"`
		LinkedStoragebox interface{} `json:"linked_storagebox"`
	} `json:"server"`
}

type Cancellation struct {
	Cancellation struct {
		ServerIP                 string      `json:"server_ip"`
		ServerIpv6Net            string      `json:"server_ipv6_net"`
		ServerNumber             int         `json:"server_number"`
		ServerName               string      `json:"server_name"`
		EarliestCancellationDate string      `json:"earliest_cancellation_date"`
		Cancelled                bool        `json:"cancelled"`
		ReservationPossible      bool        `json:"reservation_possible"`
		Reserved                 bool        `json:"reserved"`
		CancellationDate         interface{} `json:"cancellation_date"`
		CancellationReason       []string    `json:"cancellation_reason"`
	} `json:"cancellation"`
}
