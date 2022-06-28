package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
)

type RDNSClient struct {
	client *Client
}

type RDNS struct {
	IP  string
	PTR string
}

func (c *RDNSClient) GetAll(ctx context.Context) ([]*RDNS, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/rdns"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RDNSList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	rdnslist := make([]*RDNS, 0, len(body))

	for i := range body {
		firewalltemplate := RDNSFromSchema(body[i])
		rdnslist = append(rdnslist, firewalltemplate)
	}

	return rdnslist, resp, nil

}

func (c *RDNSClient) GetByIP(ctx context.Context, ip string) (*RDNS, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/rdns/%s", ip), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RDNS
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RDNSFromSchema(body), resp, nil
}

func (c *RDNSClient) CreateRDNS(ctx context.Context, ip string, ptr string) (*RDNS, *Response, error) {
	params := url.Values{}
	params.Set("ptr", ptr)

	req, err := c.client.NewRequest(ctx, "PUT", fmt.Sprintf("/rdns/%s", ip), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RDNS
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RDNSFromSchema(body), resp, nil
}

func (c *RDNSClient) UpdateRDNS(ctx context.Context, ip string, ptr string) (*RDNS, *Response, error) {
	params := url.Values{}
	params.Set("ptr", ptr)

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/rdns/%s", ip), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RDNS
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RDNSFromSchema(body), resp, nil
}

func (c *RDNSClient) Delete(ctx context.Context, ip string) (*Response, error) {

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/rdns/%s", ip), nil)

	if err != nil {
		return nil, err
	}

	var body schema.RDNS
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
