package schema

import "time"

type Server_market_list []struct {
	Product struct {
		ID             int      `json:"id"`
		Name           string   `json:"name"`
		Description    []string `json:"description"`
		Traffic        string   `json:"traffic"`
		Dist           []string `json:"dist"`
		Arch           []int    `json:"arch"`
		Lang           []string `json:"lang"`
		CPU            string   `json:"cpu"`
		CPUBenchmark   int      `json:"cpu_benchmark"`
		MemorySize     int      `json:"memory_size"`
		HddSize        int      `json:"hdd_size"`
		HddText        string   `json:"hdd_text"`
		HddCount       int      `json:"hdd_count"`
		Datacenter     string   `json:"datacenter"`
		NetworkSpeed   string   `json:"network_speed"`
		Price          string   `json:"price"`
		PriceSetup     string   `json:"price_setup"`
		PriceVat       string   `json:"price_vat"`
		PriceSetupVat  string   `json:"price_setup_vat"`
		FixedPrice     bool     `json:"fixed_price"`
		NextReduce     int      `json:"next_reduce"`
		NextReduceDate string   `json:"next_reduce_date"`
	} `json:"product"`
}

type Server_market_product struct {
	Product struct {
		ID             int      `json:"id"`
		Name           string   `json:"name"`
		Description    []string `json:"description"`
		Traffic        string   `json:"traffic"`
		Dist           []string `json:"dist"`
		Arch           []int    `json:"arch"`
		Lang           []string `json:"lang"`
		CPU            string   `json:"cpu"`
		CPUBenchmark   int      `json:"cpu_benchmark"`
		MemorySize     int      `json:"memory_size"`
		HddSize        int      `json:"hdd_size"`
		HddText        string   `json:"hdd_text"`
		HddCount       int      `json:"hdd_count"`
		Datacenter     string   `json:"datacenter"`
		NetworkSpeed   string   `json:"network_speed"`
		Price          string   `json:"price"`
		PriceSetup     string   `json:"price_setup"`
		PriceVat       string   `json:"price_vat"`
		PriceSetupVat  string   `json:"price_setup_vat"`
		FixedPrice     bool     `json:"fixed_price"`
		NextReduce     int      `json:"next_reduce"`
		NextReduceDate string   `json:"next_reduce_date"`
	} `json:"product"`
}

type ServerProductList []struct {
	Product struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Description    []string `json:"description"`
		Traffic        string   `json:"traffic"`
		Dist           []string `json:"dist"`
		DeprecatedArch []int    `json:"@deprecated arch"`
		Lang           []string `json:"lang"`
		Location       []string `json:"location"`
		Prices         []struct {
			Location string `json:"location"`
			Price    struct {
				Net   string `json:"net"`
				Gross string `json:"gross"`
			} `json:"price"`
			PriceSetup struct {
				Net   string `json:"net"`
				Gross string `json:"gross"`
			} `json:"price_setup"`
		} `json:"prices"`
		OrderableAddons []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Min    int         `json:"min"`
			Max    int         `json:"max"`
			Prices interface{} `json:"prices"`
		} `json:"orderable_addons"`
	} `json:"product"`
}

type ServerProduct struct {
	Product struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Description    []string `json:"description"`
		Traffic        string   `json:"traffic"`
		Dist           []string `json:"dist"`
		DeprecatedArch []int    `json:"@deprecated arch"`
		Lang           []string `json:"lang"`
		Location       []string `json:"location"`
		Prices         []struct {
			Location string `json:"location"`
			Price    struct {
				Net   string `json:"net"`
				Gross string `json:"gross"`
			} `json:"price"`
			PriceSetup struct {
				Net   string `json:"net"`
				Gross string `json:"gross"`
			} `json:"price_setup"`
		} `json:"prices"`
		OrderableAddons []struct {
			ID     string      `json:"id"`
			Name   string      `json:"name"`
			Min    int         `json:"min"`
			Max    int         `json:"max"`
			Prices interface{} `json:"prices"`
		} `json:"orderable_addons"`
	} `json:"product"`
}

type ServerMarketTransactionList []struct {
	Transaction struct {
		ID            string      `json:"id"`
		Date          time.Time   `json:"date"`
		Status        string      `json:"status"`
		ServerNumber  interface{} `json:"server_number"`
		ServerIP      interface{} `json:"server_ip"`
		AuthorizedKey []struct {
			Key struct {
				Name        string `json:"name"`
				Fingerprint string `json:"fingerprint"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
			} `json:"key"`
		} `json:"authorized_key"`
		HostKey []interface{} `json:"host_key"`
		Comment interface{}   `json:"comment"`
		Product struct {
			ID           int      `json:"id"`
			Name         string   `json:"name"`
			Description  []string `json:"description"`
			Traffic      string   `json:"traffic"`
			Dist         string   `json:"dist"`
			Arch         string   `json:"arch"`
			Lang         string   `json:"lang"`
			CPU          string   `json:"cpu"`
			CPUBenchmark int      `json:"cpu_benchmark"`
			MemorySize   int      `json:"memory_size"`
			HddSize      int      `json:"hdd_size"`
			HddText      string   `json:"hdd_text"`
			HddCount     int      `json:"hdd_count"`
			Datacenter   string   `json:"datacenter"`
			NetworkSpeed string   `json:"network_speed"`
		} `json:"product"`
	} `json:"transaction"`
}

type ServerMarketTransaction struct {
	Transaction struct {
		ID            string      `json:"id"`
		Date          time.Time   `json:"date"`
		Status        string      `json:"status"`
		ServerNumber  interface{} `json:"server_number"`
		ServerIP      interface{} `json:"server_ip"`
		AuthorizedKey []struct {
			Key struct {
				Name        string `json:"name"`
				Fingerprint string `json:"fingerprint"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
			} `json:"key"`
		} `json:"authorized_key"`
		HostKey []interface{} `json:"host_key"`
		Comment interface{}   `json:"comment"`
		Product struct {
			ID           int      `json:"id"`
			Name         string   `json:"name"`
			Description  []string `json:"description"`
			Traffic      string   `json:"traffic"`
			Dist         string   `json:"dist"`
			Arch         string   `json:"arch"`
			Lang         string   `json:"lang"`
			CPU          string   `json:"cpu"`
			CPUBenchmark int      `json:"cpu_benchmark"`
			MemorySize   int      `json:"memory_size"`
			HddSize      int      `json:"hdd_size"`
			HddText      string   `json:"hdd_text"`
			HddCount     int      `json:"hdd_count"`
			Datacenter   string   `json:"datacenter"`
			NetworkSpeed string   `json:"network_speed"`
		} `json:"product"`
	} `json:"transaction"`
}

type ServerOrderTransaction struct {
	Transaction struct {
		ID            string      `json:"id"`
		Date          time.Time   `json:"date"`
		Status        string      `json:"status"`
		ServerNumber  interface{} `json:"server_number"`
		ServerIP      interface{} `json:"server_ip"`
		AuthorizedKey []struct {
			Key struct {
				Name        string `json:"name"`
				Fingerprint string `json:"fingerprint"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
			} `json:"key"`
		} `json:"authorized_key"`
		HostKey []interface{} `json:"host_key"`
		Comment interface{}   `json:"comment"`
		Product struct {
			ID          string   `json:"id"`
			Name        string   `json:"name"`
			Description []string `json:"description"`
			Traffic     string   `json:"traffic"`
			Dist        string   `json:"dist"`
			Arch        int      `json:"arch"`
			Lang        string   `json:"lang"`
			Location    string   `json:"location"`
		} `json:"product"`
	} `json:"transaction"`
}

type ServerOrderTransactionList []struct {
	Transaction struct {
		ID            string      `json:"id"`
		Date          time.Time   `json:"date"`
		Status        string      `json:"status"`
		ServerNumber  interface{} `json:"server_number"`
		ServerIP      interface{} `json:"server_ip"`
		AuthorizedKey []struct {
			Key struct {
				Name        string `json:"name"`
				Fingerprint string `json:"fingerprint"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
			} `json:"key"`
		} `json:"authorized_key"`
		HostKey []interface{} `json:"host_key"`
		Comment interface{}   `json:"comment"`
		Product struct {
			ID          string   `json:"id"`
			Name        string   `json:"name"`
			Description []string `json:"description"`
			Traffic     string   `json:"traffic"`
			Dist        string   `json:"dist"`
			Arch        int      `json:"arch"`
			Lang        string   `json:"lang"`
			Location    string   `json:"location"`
		} `json:"product"`
	} `json:"transaction"`
}
