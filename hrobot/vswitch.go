package hrobot

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type VSwitch struct {
	ID        int
	Name      string
	Vlan      int
	Cancelled bool
}

type VSwitchSingle struct {
	ID           int
	Name         string
	Vlan         int
	Cancelled    bool
	Server       []interface{}
	Subnet       []interface{}
	CloudNetwork []interface{}
}

type VSwitchClient struct {
	client *Client
}

func (c *VSwitchClient) GetVSwitchList(ctx context.Context) ([]*VSwitch, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/vswitch/"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VSwitchList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	vswitchlist := make([]*VSwitch, 0, len(body))

	for i := range body {
		vswitch := VswitchFromSchema(body[i])
		vswitchlist = append(vswitchlist, vswitch)
	}

	return vswitchlist, resp, nil
}

func (c *VSwitchClient) AddVSwitch(ctx context.Context, opt *AddvSwitchOps) (*VSwitchSingle, *Response, error) {
	params, _ := query.Values(opt)
	fmt.Println(params.Encode())
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/vswitch/"), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VSwitchSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return VSwitchSingleFromSchema(body), resp, nil
}

func (c *VSwitchClient) GetVSwitchById(ctx context.Context, id string) (*VSwitchSingle, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/vswitch/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VSwitchSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	return VSwitchSingleFromSchema(body), resp, nil
}

func (c *VSwitchClient) UpdateVSwitchById(ctx context.Context, id string, opt *AddvSwitchOps) (*VSwitchSingle, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/vswitch/%s", id), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VSwitchSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return VSwitchSingleFromSchema(body), resp, nil
}

// cancellation_date format yyyy-MM-dd
func (c *VSwitchClient) CancelVSwitch(ctx context.Context, id string, cancellation_date string) (*Response, error) {
	params := url.Values{}
	params.Add("cancellation_date", cancellation_date)
	fmt.Println(params.Encode())
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/vswitch/%s", id), params)

	if err != nil {
		return nil, err
	}

	var body schema.VSwitchSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// adds server to vswitch, server can be server number or ip
func (c *VSwitchClient) AddToServer(ctx context.Context, id string, server interface{}) (*Response, error) {
	params := stringToURLValue(server, "server[]")
	//params.Add("cancellation_date", server[0])
	fmt.Println(params.Encode())
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/vswitch/%s/server", id), params)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// removes server by servernumber or ip from vswitch
func (c *VSwitchClient) RemoveServer(ctx context.Context, id string, server interface{}) (*Response, error) {
	params := stringToURLValue(server, "server[]")
	//params.Add("cancellation_date", server[0])
	fmt.Println(params.Encode())
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/vswitch/%s/server", id), params)

	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// converts a string or slice of strings to url value
func stringToURLValue(value interface{}, key string) url.Values {
	params := url.Values{}
	switch value.(type) {
	case string:
		s := reflect.ValueOf(value)
		params.Add(key, s.String())
		return params
	case []string:
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			params.Add(key, s.Index(i).String())
		}
		return params
	default:
		log.Fatalf("%v is not a string or slice of strings \n ", value)
		return nil
	}
}

type AddvSwitchOps struct {
	Name    string `url:"name"`
	Vlan_ID int    `url:"vlan"`
}

type VSwitchServer struct {
}
