package hrobot

import (
	"context"
	"fmt"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type IPClient struct {
	client *Client
}

type IP struct {
	IP              string
	ServerIP        string
	ServerNumber    int
	Locked          bool
	SeparateMac     interface{}
	TrafficWarnings bool
	TrafficHourly   int
	TrafficDaily    int
	TrafficMonthly  int
}

type MAC struct {
	IP  string
	MAC string
}

func (c *IPClient) List(ctx context.Context) ([]*IP, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/ip"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.IPList
	//fmt.Println(body)
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(len(body))

	iplist := make([]*IP, 0, len(body))

	for i := range body {
		ip := IPFromSchema(body[i])
		iplist = append(iplist, ip)
	}

	return iplist, resp, nil

}

func (c *IPClient) GetIPByIP(ctx context.Context, ip string) (*IP, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/ip/%s", ip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.IP
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return IPFromSchema(body), resp, nil
}

func (c *IPClient) UpdateTrafficByIP(ctx context.Context, ip string, opt *IPOps) (*IP, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/ip/%s", ip), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.IP
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return IPFromSchema(body), resp, nil
}

func (c *IPClient) GetMACByIP(ctx context.Context, ip string) (*MAC, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/ip/%s/mac", ip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.MAC
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return MACFromSchema(body), resp, nil
}

func (c *IPClient) GenerateMAC(ctx context.Context, ip string) (*MAC, *Response, error) {
	req, err := c.client.NewRequest(ctx, "PUT", fmt.Sprintf("/ip/%s/mac", ip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.MAC
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return MACFromSchema(body), resp, nil

}

func (c *IPClient) DeleteMAC(ctx context.Context, ip string) (*MAC, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/ip/%s/mac", ip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.MAC
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return MACFromSchema(body), resp, nil

}

type IPOps struct {
	TrafficWarnings bool   `url:"traffic_warnings"` //enable/disable traffic warnings (true,false)
	TrafficHourly   string `url:"traffic_hourly"`   // hourly traffic limit in MB
	TrafficDaily    string `url:"traffic_daily"`    // daily traffic limit in MB
	TrafficMonthly  string `url:"traffic_monthly"`  // monthly traffic limit in GB
}
