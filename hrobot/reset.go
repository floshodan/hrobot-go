package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
)

type ResetClient struct {
	client *Client
}

type Reset struct {
	ServerIP        string
	ServerIpv6Net   string
	ServerNumber    int
	Type            interface{}
	OperatingStatus string
}

func (c *ResetClient) List(ctx context.Context) ([]*Reset, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/reset"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ResetList
	//fmt.Println(body)
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(len(body))

	resetlist := make([]*Reset, 0, len(body))

	for i := range body {
		reset := ResetFromSchema(body[i])
		resetlist = append(resetlist, reset)
	}
	return resetlist, resp, nil
}

func (c *ResetClient) GetResetByServernumber(ctx context.Context, servernumber string) (*Reset, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/reset/%s", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Reset
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ResetFromSchema(body), resp, nil
}

func (c *ResetClient) ExecuteReset(ctx context.Context, servernumber string, hwtype string) (*Reset, *Response, error) {
	params := url.Values{}
	params.Set("type", hwtype)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/reset/%s/", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Reset
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return ResetFromSchema(body), resp, nil
}
