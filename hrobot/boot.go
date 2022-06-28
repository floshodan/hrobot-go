package hrobot

import (
	"context"
	"fmt"
	"net/url"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type BootClient struct {
	client *Client
}

type BootList struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Os            interface{}
	Arch          interface{}
	Active        bool
	Password      interface{}
	AuthorizedKey []interface{}
	HostKey       []interface{}
	BootTime      interface{}
	Linux         interface{}
	Vnc           interface{}
	Windows       interface{}
	Plesk         interface{}
	Cpanel        interface{}
}

type BootRescue struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Os            interface{}
	Arch          interface{}
	Active        bool
	Password      interface{}
	AuthorizedKey []interface{}
	HostKey       []interface{}
	BootTime      interface{}
}

type BootLinux struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Dist          []string
	Arch          []int
	Lang          []string
	Active        bool
	Password      string
	AuthorizedKey []interface{}
	HostKey       []interface{}
}

type BootVnc struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Dist          []string
	Arch          []int
	Lang          []string
	Active        bool
	Password      string
}

type BootWindows struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Dist          []string
	Lang          []string
	Active        bool
	Password      string
}

func (c *BootClient) GetBootOptions(ctx context.Context, servernumber string) (*BootList, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.BootList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	return BootListFromSchema(body), resp, nil
}

// query boot options for the Rescue System
func (c *BootClient) GetRescue(ctx context.Context, servernumber string) (*BootRescue, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/rescue", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RescueList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

func (c *BootClient) ActivateRescue(ctx context.Context, servernumber string, opt *RescueOpts) (*BootRescue, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/boot/%s/rescue", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RescueList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

func (c *BootClient) DeactivateRescue(ctx context.Context, servernumber string) (*BootRescue, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/boot/%s/rescue", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RescueList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

// Show data of last rescue activation
func (c *BootClient) GetLastRescue(ctx context.Context, servernumber string) (*BootRescue, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/rescue/last", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.RescueList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

func (c *BootClient) GetLinux(ctx context.Context, servernumber string) (*BootLinux, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/linux", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.LinuxList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return LinuxFromSchema(body), resp, nil
}

func (c *BootClient) ActivateLinux(ctx context.Context, servernumber string, opt *LinuxOpts) (*BootLinux, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/boot/%s/linux", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.LinuxList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return LinuxFromSchema(body), resp, nil
}

func (c *BootClient) DeactivateLinux(ctx context.Context, servernumber string) (*BootLinux, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/boot/%s/linux", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.LinuxList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return LinuxFromSchema(body), resp, nil
}

// Show data of last rescue activation
func (c *BootClient) GetLastLinux(ctx context.Context, servernumber string) (*BootLinux, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/linux/last", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.LinuxList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return LinuxFromSchema(body), resp, nil
}

func (c *BootClient) GetVNC(ctx context.Context, servernumber string) (*BootVnc, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/vnc", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VncList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return VncFromSchema(body), resp, nil
}

func (c *BootClient) ActivateVNC(ctx context.Context, servernumber string, opt *VNCOpts) (*BootVnc, *Response, error) {
	params, _ := query.Values(opt)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/boot/%s/vnc", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VncList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return VncFromSchema(body), resp, nil
}

func (c *BootClient) DeactivateVNC(ctx context.Context, servernumber string) (*BootVnc, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/boot/%s/vnc", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.VncList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return VncFromSchema(body), resp, nil
}

func (c *BootClient) GetWindows(ctx context.Context, servernumber string) (*BootWindows, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s/windows", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.WindowsList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return WindowsFromSchema(body), resp, nil
}

func (c *BootClient) ActivateWindows(ctx context.Context, servernumber string, language string) (*BootWindows, *Response, error) {
	params := url.Values{}
	params.Set("lang", language)

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/boot/%s/windows", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.WindowsList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return WindowsFromSchema(body), resp, nil
}

func (c *BootClient) DeactivateWindows(ctx context.Context, servernumber string) (*BootWindows, *Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/boot/%s/windows", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.WindowsList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return WindowsFromSchema(body), resp, nil
}

type RescueOpts struct {
	OS             string `url:"os"` //required
	Arch           string `url:"arch"`
	Authorized_Key string `url:"authorized_key"`
}

type LinuxOpts struct {
	Dist           string `url:"dist"` //required
	Lang           string `url:"lang"` //required
	Authorized_Key string `url:"authorized_key"`
}

type VNCOpts struct {
	Dist string `url:"dist"` //required
	Lang string `url:"lang"` //required
}

type PleskOpts struct {
	Dist     string `url:"dist"` //required
	Lang     string `url:"lang"` //required
	Hostname string `url:"hostname"`
}

type CpanelOpts struct {
	Dist     string `url:"dist"` //required
	Lang     string `url:"lang"` //required
	Hostname string `url:"hostname"`
}
