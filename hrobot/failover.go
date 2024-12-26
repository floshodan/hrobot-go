package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
)

type FailoverClient struct {
	client *Client
}

type Failover struct {
	IP             string
	Netmask        string
	ServerIP       string
	ServerNumber   int
	ActiveServerIP string
}

func (c *FailoverClient) List(ctx context.Context) ([]*Failover, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/failover/"), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.FailoverList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	failovertemplatelist := make([]*Failover, 0, len(body))

	for i := range body {
		failovertemplate := FailoverFromSchema(body[i])
		failovertemplatelist = append(failovertemplatelist, failovertemplate)
	}

	return failovertemplatelist, resp, nil
}

func (c *FailoverClient) GetFailoverIP(ctx context.Context, ip string) (*Failover, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/failover/%s", ip), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Failover
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FailoverFromSchema(body), resp, nil
}

// switches failover routing to given IP
func (c *FailoverClient) SwitchFailover(ctx context.Context, failover_ip, active_server_ip string) (*Failover, *Response, error) {
	requestBody := url.Values{}
	requestBody.Add("active_server_ip", active_server_ip)

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/failover/%s", failover_ip), requestBody)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Failover
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FailoverFromSchema(body), resp, nil
}

// deletes the routing of a given IP
func (c *FailoverClient) DeleteFailover(ctx context.Context, failover_ip string) (*Failover, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/failover/%s/", failover_ip), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.Failover
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FailoverFromSchema(body), resp, nil
}
