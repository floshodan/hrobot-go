package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type StorageBox struct {
	ID           int
	Login        string
	Name         string
	Product      string
	Cancelled    bool
	Locked       bool
	Location     string
	LinkedServer int
	PaidUntil    string
}

type StorageBoxSingle struct {
	ID                   int
	Login                string
	Name                 string
	Product              string
	Cancelled            bool
	Locked               bool
	Location             string
	LinkedServer         int
	PaidUntil            string
	DiskQuota            int
	DiskUsage            int
	DiskUsageData        int
	DiskUsageSnapshots   int
	WebDAV               bool
	Samba                bool
	SSH                  bool
	ExternalReachability bool
	ZFS                  bool
	Server               string
	HostSystem           string
}

type Password struct {
	Password string
}

type StorageBoxClient struct {
	client *Client
}

func (c *StorageBoxClient) GetStorageBoxList(ctx context.Context) ([]*StorageBox, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", "/storagebox/", nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.StorageBoxList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	storageboxlist := make([]*StorageBox, 0, len(body))

	for i := range body {
		storagebox := StorageBoxFromSchema(body[i])
		storageboxlist = append(storageboxlist, storagebox)
	}

	return storageboxlist, resp, nil
}

func (c *StorageBoxClient) GetStorageBoxById(ctx context.Context, id int) (*StorageBoxSingle, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/storagebox/%d", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.StorageBoxSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	return StorageBoxSingleFromSchema(body), resp, nil
}

func (c *StorageBoxClient) UpdateStorageBoxById(ctx context.Context, id int, opt *UpdateStorageBoxOps) (*StorageBoxSingle, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/storagebox/%d", id), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.StorageBoxSingle
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return StorageBoxSingleFromSchema(body), resp, nil
}

func (c *StorageBoxClient) ResetStorageBoxPasswordById(ctx context.Context, id int, password *string) (*Password, *Response, error) {
	data := url.Values{}
	if password != nil {
		data.Set("password", *password)
	}
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/storagebox/%d/password", id), data)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Password
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return PasswordFromSchema(body), resp, nil
}

type UpdateStorageBoxOps struct {
	StorageBoxName       string `url:"storagebox_name"`
	Samba                bool   `json:"samba"`
	WebDAV               bool   `json:"webdav"`
	SSH                  bool   `json:"ssh"`
	ExternalReachability bool   `json:"external_reachability"`
	ZFS                  bool   `json:"zfs"`
}
