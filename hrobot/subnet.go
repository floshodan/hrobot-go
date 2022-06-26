package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type Subnet struct {
	IP              string
	Mask            int
	Gateway         string
	ServerIP        string
	ServerNumber    int
	Failover        bool
	Locked          bool
	TrafficWarnings bool
	TrafficHourly   int
	TrafficDaily    int
	TrafficMonthly  int
}

type SubnetClient struct {
	client *Client
}

type SubnetMac struct {
	IP          string
	Mask        int
	Mac         string
	PossibleMac interface{}
}

type SubnetCancellation struct {
	IP                       string
	Mask                     int
	ServerNumber             int
	EarliestCancellationDate string
	Cancelled                bool
	CancellationDate         interface{}
}

func (c *SubnetClient) List(ctx context.Context) ([]*Subnet, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/subnet/"), nil)
	//	fmt.Println(req)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetList
	//fmt.Println(body)
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(len(body))

	subnetlist := make([]*Subnet, 0, len(body))

	for i := range body {
		subnet := SubnetFromSchema(body[i])
		subnetlist = append(subnetlist, subnet)
	}
	return subnetlist, resp, nil
}

// requires the net ip of the subnet
func (c *SubnetClient) GetSubnetByIP(ctx context.Context, netip string) (*Subnet, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/subnet/%s", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Subnet
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetFromSchema(body), resp, nil
}

// requires the net ip of the subnet
func (c *SubnetClient) UpdateTraffic(ctx context.Context, netip string, opt *SubnetOps) (*Subnet, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/subnet/%s", netip), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Subnet
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetFromSchema(body), resp, nil
}

func (c *SubnetClient) GetMac(ctx context.Context, netip string) (*SubnetMac, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/subnet/%s/mac", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetMac
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetMacFromSchema(body), resp, nil
}

func (c *SubnetClient) GenerateMac(ctx context.Context, netip string) (*SubnetMac, *Response, error) {
	req, err := c.client.NewRequest(ctx, "PUT", fmt.Sprintf("/subnet/%s/mac", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetMac
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetMacFromSchema(body), resp, nil
}

func (c *SubnetClient) DeleteMAC(ctx context.Context, netip string) (*SubnetMac, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/subnet/%s/mac", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetMac
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetMacFromSchema(body), resp, nil
}

func (c *SubnetClient) GetCancellation(ctx context.Context, netip string) (*SubnetCancellation, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/subnet/%s/cancellation", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetCancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetCancellationFromSchema(body), resp, nil
}

func (c *SubnetClient) PostCancellation(ctx context.Context, netip string, cancellation_date string) (*SubnetCancellation, *Response, error) {
	data := url.Values{}
	data.Set("cancellation_date", cancellation_date)

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/subnet/%s/cancellation", netip), data)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetCancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetCancellationFromSchema(body), resp, nil
}

func (c *SubnetClient) DeleteCancellation(ctx context.Context, netip string) (*SubnetCancellation, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/subnet/%s/cancellation", netip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SubnetCancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SubnetCancellationFromSchema(body), resp, nil
}

type SubnetOps struct {
	TrafficWarnings bool   `url:"traffic_warnings"` //enable/disable traffic warnings (true,false)
	TrafficHourly   string `url:"traffic_hourly"`   // hourly traffic limit in MB
	TrafficDaily    string `url:"traffic_daily"`    // daily traffic limit in MB
	TrafficMonthly  string `url:"traffic_monthly"`  // monthly traffic limit in GB
}
