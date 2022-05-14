package schema

type VSwitchList []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Vlan      int    `json:"vlan"`
	Cancelled bool   `json:"cancelled"`
}

type VSwitch struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Vlan      int    `json:"vlan"`
	Cancelled bool   `json:"cancelled"`
}

type VSwitchSingle struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Vlan         int           `json:"vlan"`
	Cancelled    bool          `json:"cancelled"`
	Server       []interface{} `json:"server"`
	Subnet       []interface{} `json:"subnet"`
	CloudNetwork []interface{} `json:"cloud_network"`
}
