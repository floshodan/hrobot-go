package hrobot

import (
	"context"
	"fmt"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type SSHKey struct {
	Name        string
	Fingerprint string
	Type        string
	Size        int
	Data        string
}

type SSHKeyClient struct {
	client *Client
}

// GetByID retrieves a SSH key by its ID. If the SSH key does not exist, nil is returned.
func (c *SSHKeyClient) List(ctx context.Context) ([]*SSHKey, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/key"), nil)
	//	fmt.Println(req)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SSHKeys
	//fmt.Println(body)
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(len(body))

	sshKeys := make([]*SSHKey, 0, len(body))

	for i := range body {
		key := SSHKeyFromSchema(body[i])
		sshKeys = append(sshKeys, key)
	}

	return sshKeys, resp, nil
}

func (c *SSHKeyClient) Create(ctx context.Context, opt *CreateKeyOpts) (*SSHKey, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/key"), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Key
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return SSHKeyFromSchema(body), resp, err
}

func (c *SSHKeyClient) GetByFingerprint(ctx context.Context, fingerprint string) (*SSHKey, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/key/%s", fingerprint), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Key
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SSHKeyFromSchema(body), resp, err
}

func (c *SSHKeyClient) Update(ctx context.Context, fingerprint string, opt *UpdateKeyOpts) (*SSHKey, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/key/%s", fingerprint), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Key
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return SSHKeyFromSchema(body), resp, err

}

func (c *SSHKeyClient) Delete(ctx context.Context, fingerprint string) (*Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/key/%s", fingerprint), nil)

	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req, nil)
	if err != nil {
		return nil, err
	}
	return resp, err

}

type CreateKeyOpts struct {
	Name string `url:"name"`
	Data string `url:"data"`
}

type UpdateKeyOpts struct {
	Name string `url:"name"`
}
