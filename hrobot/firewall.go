package hrobot

import (
	"context"
	"fmt"
	"reflect"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type FirewallClient struct {
	client *Client
}

type Firewall struct {
	ServerIP     string
	ServerNumber int
	Status       string
	Filter_IPv6  bool
	WhitelistHos bool
	Port         string
	Rules        struct {
		Input []struct {
			IPVersion string
			Name      string
			DstIP     interface{}
			SrcIP     string
			DstPort   string
			SrcPort   interface{}
			Protocol  interface{}
			TCPFlags  interface{}
			Action    string
		}
		Output []struct {
			IPVersion string
			Name      string
			DstIP     interface{}
			SrcIP     string
			DstPort   string
			SrcPort   interface{}
			Protocol  interface{}
			TCPFlags  interface{}
			Action    string
		}
	}
}

type FirewallRules struct {
	Rules struct {
		Input []struct {
			IPVersion string      `json:"ip_version"`
			Name      string      `json:"name"`
			DstIP     interface{} `json:"dst_ip"`
			SrcIP     string      `json:"src_ip"`
			DstPort   string      `json:"dst_port"`
			SrcPort   interface{} `json:"src_port"`
			Protocol  interface{} `json:"protocol"`
			TCPFlags  interface{} `json:"tcp_flags"`
			Action    string      `json:"action"`
		} `json:"input"`
	} `json:"rules"`
}

type InputRule struct {
	IPVersion string `url:"ip_version"`
	Name      string `url:"name"`
	DstIP     string `url:"dst_ip"`
	SrcIP     string `url:"src_ip"`
	DstPort   string `url:"dst_port"`
	SrcPort   string `url:"src_port"`
	Protocol  string `url:"protocol"`
	TCPFlags  string `url:"tcp_flags"`
	Action    string `url:"action"`
}

type FirewallTemplate struct {
	ID           int
	Name         string
	Filter_IPv6  bool
	WhitelistHos bool
	IsDefault    bool
}

type FirewallTemplateWithRules struct {
	ID           int `json:"id"`
	Name         string
	Filter_IPv6  bool
	WhitelistHos bool `json:"whitelist_hos"`
	IsDefault    bool `json:"is_default"`
	Rules        struct {
		Input []struct {
			IPVersion    string
			Name         string
			DstIP        interface{}
			SrcIP        string
			DstPort      string
			SrcPort      interface{}
			Protocol     interface{}
			TCPFlags     interface{}
			PacketLength interface{}
			Action       string
		}
		Output []struct {
			IPVersion    string
			Name         string
			DstIP        interface{}
			SrcIP        string
			DstPort      string
			SrcPort      interface{}
			Protocol     interface{}
			TCPFlags     interface{}
			PacketLength interface{}
			Action       string
		}
	}
}

func (c *FirewallClient) GetFirewallByServernumber(ctx context.Context, servernumber string) (*Firewall, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/firewall/%s", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Firewall
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FirewallFromSchema(body), resp, nil
}

func (c *FirewallClient) PostFirewallByServernumber(ctx context.Context, servernumber string, opt *FirewallOps) (*Firewall, *Response, error) {
	params, _ := query.Values(opt)

	//add rules to params
	for index, service := range opt.Rules {
		v := reflect.ValueOf(service)
		for i := 0; i < v.NumField(); i++ {
			//if empty string the field must be empty
			if v.Field(i).String() != "" {
				params.Add(fmt.Sprintf("rules[input][%v][%v]", index, v.Type().Field(i).Tag.Get("url")), v.Field(i).String())
			}
		}
	}
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/firewall/%s", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.Firewall
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FirewallFromSchema(body), resp, nil

}

func (c *FirewallClient) DeleteFirewallByServerNumber(ctx context.Context, servernumber string) (*Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/firewall/%s", servernumber), nil)

	if err != nil {
		return nil, err
	}

	var body schema.Firewall
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *FirewallClient) GetFirewallByIP(ctx context.Context, ip string) (*Firewall, *Response, error) {
	return c.GetFirewallByServernumber(ctx, ip)
}

func (c *FirewallClient) PostFirewallByIP(ctx context.Context, ip string, opt *FirewallOps) (*Firewall, *Response, error) {
	return c.PostFirewallByServernumber(ctx, ip, &FirewallOps{})
}

func (c *FirewallClient) DeleteFirewallByIP(ctx context.Context, ip string) (*Response, error) {
	return c.DeleteFirewallByServerNumber(ctx, ip)
}

func (c *FirewallClient) GetFirewallTemplate(ctx context.Context) ([]*FirewallTemplate, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/firewall/template"), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.FirewallTemplateList
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	firewalltemplatelist := make([]*FirewallTemplate, 0, len(body))

	for i := range body {
		firewalltemplate := FirewallTemplateFromSchema(body[i])
		firewalltemplatelist = append(firewalltemplatelist, firewalltemplate)
	}

	return firewalltemplatelist, resp, nil
}

func (c *FirewallClient) PostFirewallTemplate(ctx context.Context, opt *FirewallTemplateOps) (*FirewallTemplateWithRules, *Response, error) {
	params, _ := query.Values(opt)

	//add rules to params
	for index, service := range opt.Rules {
		v := reflect.ValueOf(service)
		for i := 0; i < v.NumField(); i++ {
			//if empty string the field must be empty
			if v.Field(i).String() != "" {
				params.Add(fmt.Sprintf("rules[input][%v][%v]", index, v.Type().Field(i).Tag.Get("url")), v.Field(i).String())
			}
		}
	}
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/firewall/template"), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.FirewallTemplateWithRules
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FirewallTemplateWithRulesFromSchema(body), resp, nil

}

func (c *FirewallClient) GetFirewallTemplateById(ctx context.Context, id string) (*FirewallTemplateWithRules, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/firewall/template/%s", id), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.FirewallTemplateWithRules
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return FirewallTemplateWithRulesFromSchema(body), resp, nil
}

func (c *FirewallClient) UpdateFirewallTemplateById(ctx context.Context, id string, opt *FirewallTemplateOps) (*FirewallTemplateWithRules, *Response, error) {
	params, _ := query.Values(opt)
	//add rules to params
	for index, service := range opt.Rules {
		v := reflect.ValueOf(service)
		for i := 0; i < v.NumField(); i++ {
			//if empty string the field must be empty
			if v.Field(i).String() != "" {
				params.Add(fmt.Sprintf("rules[input][%v][%v]", index, v.Type().Field(i).Tag.Get("url")), v.Field(i).String())
			}
		}
	}
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/firewall/template/%s", id), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.FirewallTemplateWithRules
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}
	return FirewallTemplateWithRulesFromSchema(body), resp, nil
}

func (c *FirewallClient) DeleteFirewallTemplateById(ctx context.Context, id string) (*Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("/firewall/template/%s", id), nil)

	if err != nil {
		return nil, err
	}

	var body schema.Firewall
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type FirewallOps struct {
	Status        string      `url:"status"`        //change the status of the firewall ('active' or 'disabled')
	Filter_IPv6   string      `url:"filter_ipv6"`   // activate or deactivate the ipv6 filter ('true' or 'false', optional)
	Whitelist_hos string      `url:"whitelist_hos"` // change the flag of hetzner services whitelisting (true or false)
	Rules         []InputRule `url:"-"`
	Template_id   string      `url:"template_id,omitempty"` // not possible in combination of whitelist_hos and rules
}

type FirewallTemplateOps struct {
	Name          string      `url:"name"`          //template name
	Filter_IPv6   string      `url:"filter_ipv6"`   // activate or deactivate the ipv6 filter ('true' or 'false', optional)
	Whitelist_hos bool        `url:"whitelist_hos"` // change the flag of hetzner services whitelisting (true or false)
	Is_Default    bool        `url:"is_default"`
	Rules         []InputRule `url:"-"`
}

type RulesJSONtoURLENCode func()
