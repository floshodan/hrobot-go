package schema

type StorageBoxList []struct {
	StorageBox struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		Name         string `json:"name"`
		Product      string `json:"product"`
		Cancelled    bool   `json:"cancelled"`
		Locked       bool   `json:"locked"`
		Location     string `json:"location"`
		LinkedServer int    `json:"linked_server"`
		PaidUntil    string `json:"paid_until"`
	} `json:"storagebox"`
}

type StorageBox struct {
	StorageBox struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		Name         string `json:"name"`
		Product      string `json:"product"`
		Cancelled    bool   `json:"cancelled"`
		Locked       bool   `json:"locked"`
		Location     string `json:"location"`
		LinkedServer int    `json:"linked_server"`
		PaidUntil    string `json:"paid_until"`
	} `json:"storagebox"`
}

type StorageBoxSingle struct {
	StorageBox struct {
		ID                   int    `json:"id"`
		Login                string `json:"login"`
		Name                 string `json:"name"`
		Product              string `json:"product"`
		Cancelled            bool   `json:"cancelled"`
		Locked               bool   `json:"locked"`
		Location             string `json:"location"`
		LinkedServer         int    `json:"linked_server"`
		PaidUntil            string `json:"paid_until"`
		DiskQuota            int    `json:"disk_quota"`
		DiskUsage            int    `json:"disk_usage"`
		DiskUsageData        int    `json:"disk_usage_data"`
		DiskUsageSnapshots   int    `json:"disk_usage_snapshots"`
		WebDAV               bool   `json:"webdav"`
		Samba                bool   `json:"samba"`
		SSH                  bool   `json:"ssh"`
		ExternalReachability bool   `json:"external_reachability"`
		ZFS                  bool   `json:"zfs"`
		Server               string `json:"server"`
		HostSystem           string `json:"host_system"`
	} `json:"storagebox"`
}

type Password struct {
	Password string `json:"password"`
}
