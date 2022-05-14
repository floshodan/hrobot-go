package hrobot

import (
	"context"
	"fmt"

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
func (c *SubnetClient) UpdateTrafficBySubnet(ctx context.Context, netip string, opt *SubnetOps) (*Subnet, *Response, error) {
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

type SubnetOps struct {
	TrafficWarnings bool   `url:"traffic_warnings"` //enable/disable traffic warnings (true,false)
	TrafficHourly   string `url:"traffic_hourly"`   // hourly traffic limit in MB
	TrafficDaily    string `url:"traffic_daily"`    // daily traffic limit in MB
	TrafficMonthly  string `url:"traffic_monthly"`  // monthly traffic limit in GB
}
