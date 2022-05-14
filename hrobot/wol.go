package hrobot

import (
	"context"
	"fmt"

	"github.com/floshodan/hrobot-go/hrobot/schema"
)

type WOLClient struct {
	client *Client
}

type WOL struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
}

func (c *WOLClient) GetWOLByServernumber(ctx context.Context, servernumber string) (*WOL, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/wol/%s", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.WOL
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return WOLFromSchema(body), resp, nil
}
