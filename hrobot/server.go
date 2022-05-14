package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type Server struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	ServerName    string
	Product       string
	Dc            string
	Traffic       string
	Status        string
	Cancelled     bool
	PaidUntil     string
	IP            []string
	Subnet        []struct {
		IP   string
		Mask string
	}
	LinkedStoragebox interface{}
}

type SingleServer struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	ServerName    string
	Product       string
	Dc            string
	Traffic       string
	Status        string
	Cancelled     bool
	PaidUntil     string
	IP            []string
	Subnet        []struct {
		IP   string
		Mask string
	}
	Reset            bool
	Rescue           bool
	Vnc              bool
	Windows          bool
	Plesk            bool
	Cpanel           bool
	Wol              bool
	HotSwap          bool
	LinkedStoragebox interface{}
}

type Cancellation struct {
	ServerIP                 string
	ServerIpv6Net            string
	ServerNumber             int
	ServerName               string
	EarliestCancellationDate string
	Cancelled                bool
	ReservationPossible      bool
	Reserved                 bool
	CancellationDate         interface{}
	CancellationReason       []string
}

type ServerClient struct {
	client *Client
}

func (c *ServerClient) List(ctx context.Context) ([]*Server, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/server"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.ServerList
	//fmt.Println(body)
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(len(body))

	servers := make([]*Server, 0, len(body))

	for i := range body {
		server := ServerFromSchema(body[i])
		servers = append(servers, server)
	}

	return servers, resp, nil
}

func (c *ServerClient) GetServerById(ctx context.Context, id string) (*SingleServer, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/server/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleServer
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SingleServerFromSchema(body), resp, nil
}

// updates servername
func (c *ServerClient) UpdateServerName(ctx context.Context, id string, servername string) (*SingleServer, *Response, error) {
	data := url.Values{}
	data.Set("server_name", servername)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/server/%s", id), data)

	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleServer
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SingleServerFromSchema(body), resp, nil
}

// get the cancellation status of a server
func (c *ServerClient) GetCancellation(ctx context.Context, servernumber string) (*Cancellation, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/server/%s/cancellation", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Cancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return CancellationFromSchema(body), resp, nil
}

// cancel a server
func (c *ServerClient) PostCancellation(ctx context.Context, servernumber string, opt *CancellationOps) (*Cancellation, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/server/%s/cancellation", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Cancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return CancellationFromSchema(body), resp, nil
}

// revoke a cancellation
func (c *ServerClient) DeleteCancellation(ctx context.Context, servernumber string) (*Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/server/%s/cancellation", servernumber), nil)

	if err != nil {
		return nil, err
	}

	var body schema.Cancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// cancel orderd server
func (c *ServerClient) PostReversal(ctx context.Context, servernumber string, reversal_reason string) (*Cancellation, *Response, error) {
	params := url.Values{}
	params.Add("revesal_reason", reversal_reason)

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/server/%s/reversal", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Cancellation
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	return CancellationFromSchema(body), resp, nil
}

type CancellationOps struct {
	CancellationDate   string `url:"cancellation_date"`
	CancellationReason string `url:"cancellation_reason"`
	ReverseLocation    bool   `url:"reserve_location"`
}
