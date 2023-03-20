package hrobot

import (
	"context"
	"fmt"
	"time"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type ServerMarketProduct struct {
	ID             int
	Name           string
	Description    []string
	Traffic        string
	Dist           []string
	Arch           []int
	Lang           []string
	CPU            string
	CPUBenchmark   int
	MemorySize     int
	HddSize        int
	HddText        string
	HddCount       int
	Datacenter     string
	NetworkSpeed   string
	Price          string
	PriceSetup     string
	PriceVat       string
	PriceSetupVat  string
	FixedPrice     bool
	NextReduce     int
	NextReduceDate string
}

type ServerProduct struct {
	ID          string
	Name        string
	Description []string
	Traffic     string
	Dist        []string
	Arch        []int
	Lang        []string
	Location    []string
	Prices      []struct {
		Location string
		Price    struct {
			Net   string
			Gross string
		}
		PriceSetup struct {
			Net   string
			Gross string
		}
	}
	OrderableAddons []struct {
		ID     string
		Name   string
		Min    int
		Max    int
		Prices interface{}
	}
}

type ServerOrderTransaction struct {
	ID            string
	Date          time.Time
	Status        string
	ServerNumber  interface{}
	ServerIP      interface{}
	AuthorizedKey []struct {
		Key struct {
			Name        string
			Fingerprint string
			Type        string
			Size        int
		}
	}
	HostKey []interface{}
	Comment interface{}
	Product struct {
		ID          string
		Name        string
		Description []string
		Traffic     string
		Dist        string
		Arch        int
		Lang        string
		Location    string
	}
	Addons []interface{}
}

type ServerMarketTransaction struct {
	ID            string
	Date          time.Time
	Status        string
	ServerNumber  interface{}
	ServerIP      interface{}
	AuthorizedKey []struct {
		Key struct {
			Name        string
			Fingerprint string
			Type        string
			Size        int
		}
	}
	HostKey []interface{}
	Comment interface{}
	Product struct {
		ID           int
		Name         string
		Description  []string
		Traffic      string
		Dist         string
		Arch         string
		Lang         string
		CPU          string
		CPUBenchmark int
		MemorySize   int
		HddSize      int
		HddText      string
		HddCount     int
		Datacenter   string
		NetworkSpeed string
	}
}

type OrderClient struct {
	client *Client
}

func (c *OrderClient) ProductList(ctx context.Context, opt *OrderServerListOpts) ([]*ServerProduct, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server/product?%s", params.Encode()), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerProductList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	server_product := make([]*ServerProduct, 0, len(body))

	for i := range body {
		server := ServerProductFromSchema(body[i])
		server_product = append(server_product, server)
	}

	return server_product, resp, nil

}

func (c *OrderClient) GetProductById(ctx context.Context, id string) (*ServerProduct, *Response, error) {

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server/product/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerProduct
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerProductFromSchema(body), resp, nil
}

// overview of the last orders within 30 days
func (c *OrderClient) GetServerTransactionList(ctx context.Context) ([]*ServerOrderTransaction, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", "/order/server/transaction/", nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerOrderTransactionList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	servertransactionlist := make([]*ServerOrderTransaction, 0, len(body))

	for i := range body {
		serverorder := ServerOrderTransactionFromSchema(body[i])
		servertransactionlist = append(servertransactionlist, serverorder)
	}

	return servertransactionlist, resp, nil

}

// orders a server return 201 if succesful
func (c *OrderClient) OrderServer(ctx context.Context, opt *OrderServerOpts) (*ServerOrderTransaction, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", "/order/server/transaction", params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerOrderTransaction
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerOrderTransactionFromSchema(body), resp, nil
}

func (c *OrderClient) GetServerTransactionById(ctx context.Context, id string) (*ServerOrderTransaction, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server/transaction/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerOrderTransaction
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerOrderTransactionFromSchema(body), resp, nil

}

func (c *OrderClient) ListMarket(ctx context.Context, opt *OrderMarketListOpts) ([]*ServerMarketProduct, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server_market/product?%s", params.Encode()), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Server_market_list
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	server_market := make([]*ServerMarketProduct, 0, len(body))

	for i := range body {
		server := ServerMarketOrderFromSchema(body[i])
		server_market = append(server_market, server)
	}

	return server_market, resp, nil

}

func (c *OrderClient) GetMarketServerById(ctx context.Context, transaction int) (*ServerMarketProduct, *Response, error) {

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server_market/product/%s", fmt.Sprint(transaction)), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Server_market_product
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerMarketOrderFromSchema(body), resp, nil
}

func (c *OrderClient) GetServerMarketTransactionList(ctx context.Context) ([]*ServerMarketTransaction, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", "/order/server_market/transaction/", nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerMarketTransactionList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	servertransactionlist := make([]*ServerMarketTransaction, 0, len(body))

	for i := range body {
		serverorder := ServerMarketTransactionFromSchema(body[i])
		servertransactionlist = append(servertransactionlist, serverorder)
	}

	return servertransactionlist, resp, nil

}

func (c *OrderClient) OrderServerMarket(ctx context.Context, opt *OrderMarketOpts) (*ServerMarketTransaction, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", "/order/server_market/transaction", params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerMarketTransaction
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerMarketTransactionFromSchema(body), resp, nil

}

func (c *OrderClient) GetMarketTransactionById(ctx context.Context, id string) (*ServerMarketTransaction, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/order/server_market/transaction/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerMarketTransaction
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ServerMarketTransactionFromSchema(body), resp, nil

}

type OrderMarketListOpts struct {
	CPU               string `url:"cpu"`
	Min_CPU_Benchmark string `url:"min_cpu_benchmark"`
	Max_CPU_Benchmark string `url:"max_cpu_benchmark"`
	Min_Memory_Size   string `url:"min_memory_size"`
	Max_Memory_Size   string `url:"max_memory_size"`
	Min_HDD_Size      string `url:"min_hdd_size"`
	Max_HDD_Size      string `url:"max_hdd_size"`
	Search            string `url:"search"`
	Min_Price         string `url:"min_price"`
	Max_Price         string `url:"max_price"`
}

type OrderServerListOpts struct {
	Min_Price       string `url:"min_price"`
	Max_Price       string `url:"max_price"`
	Min_Price_Setup string `url:"min_price_setup"`
	Max_Price_Setup string `url:"max_price_setup"`
	Location        string `url:"location"`
}

type OrderMarketOpts struct {
	Product_ID      string `url:"product_id"`       //required
	Authorized_Keys string `url:"authorized_key[]"` //required
	//Password        string `url:"password"`
	Dist    string `url:"dist"`
	Arch    string `url:"arch"`
	Lang    string `url:"lang"`
	Comment string `url:"comment"`
	Test    bool   `url:"test"`
}

type OrderServerOpts struct {
	Product_ID     string `url:"product_id"`       //required
	Authorized_Key string `url:"authorized_key[]"` //required
	//Password string `url:"password"`
	Location string `url:"location"` //required
	Dist     string `url:"dist"`
	Arch     string `url:"arch"`
	Lang     string `url:"lang"`
	Comment  string `url:"comment"`
	Addons   string `url:"addons[]"`
	Test     bool   `url:"test"`
}
