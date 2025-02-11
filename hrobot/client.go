package hrobot

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

// Endpoint is th
const Endpoint = "https://robot-ws.your-server.de"

type Client struct {
	endpoint    string
	basicAuth   string
	httpClient  *http.Client
	debugWriter io.Writer
	userAgent   string

	SSHKey     SSHKeyClient
	Server     ServerClient
	Order      OrderClient
	Boot       BootClient
	IP         IPClient
	Subnet     SubnetClient
	Reset      ResetClient
	WakeOnLane WOLClient
	Firewall   FirewallClient
	VSwitch    VSwitchClient
	Failover   FailoverClient
	RDNS       RDNSClient
	StorageBox  StorageBoxClient
}

type ClientOption func(*Client)

func WithBasicAuth(username, password string) ClientOption {
	auth := username + ":" + password
	return func(client *Client) {
		client.basicAuth = base64.StdEncoding.EncodeToString([]byte(auth))
	}
}

// Token must be provided in this format: username:password
func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.basicAuth = base64.StdEncoding.EncodeToString([]byte(token))
	}
}

// WithDebugWriter configures a Client to print debug information to the given
// writer. To, for example, print debug information on stderr, set it to os.Stderr.
func WithDebugWriter(debugWriter io.Writer) ClientOption {
	return func(client *Client) {
		client.debugWriter = debugWriter
	}
}

// WithHTTPClient configures a Client to perform HTTP requests with httpClient.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{
		endpoint:   Endpoint,
		httpClient: &http.Client{},
		userAgent:  "curl",
		//		backoffFunc:  ExponentialBackoff(2, 500*time.Millisecond),
		//		pollInterval: 500 * time.Millisecond,
	}

	for _, option := range options {
		option(client)
	}

	client.SSHKey = SSHKeyClient{client: client}
	client.Server = ServerClient{client: client}
	client.Order = OrderClient{client: client}
	client.Boot = BootClient{client: client}
	client.IP = IPClient{client: client}
	client.Subnet = SubnetClient{client: client}
	client.Reset = ResetClient{client: client}
	client.WakeOnLane = WOLClient{client: client}
	client.Firewall = FirewallClient{client: client}
	client.VSwitch = VSwitchClient{client: client}
	client.Failover = FailoverClient{client: client}
	client.RDNS = RDNSClient{client: client}
	client.StorageBox = StorageBoxClient{client: client}

	return client
}

// NewRequest creates an HTTP request against the API. The returned request
// is assigned with ctx and has all necessary headers set (auth, user agent, etc.).
func (c *Client) NewRequest(ctx context.Context, method, path string, data url.Values) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if c.basicAuth != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.basicAuth))
	}
	if data == nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if method == "POST" || method == "DELETE" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}
	req = req.WithContext(ctx)
	return req, nil
}

// Do performs an HTTP request against the API.
func (c *Client) Do(r *http.Request, v interface{}) (*Response, error) {
	var err error
	resp, err := c.httpClient.Do(r)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	// return just the response due to empty body in DELETE method
	if r.Method == "DELETE" || v == nil {
		return response, nil
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		switch resp.StatusCode {
		case 400:
			err = fmt.Errorf("invalid input error")
			return response, err
		case 401:
			err = fmt.Errorf("401 unauthorized error")
			return response, err
		case 403:
			err = fmt.Errorf("403 Forbidden")
			return response, err
		case 404:
			err = fmt.Errorf("404 Not found")
			return response, err
		case 409:
			err = fmt.Errorf("409 Already in use")
			return response, err
		case 503:
			err = fmt.Errorf("503 - Service Unavailable")
			return response, err
		default:
			err = fmt.Errorf("hcrobot: server responded with status code %d", resp.StatusCode)
			return response, err

		}
		/* TODO Proper ResponseCode Parsing
		//err = errorFromResponse(resp, body)
		if err == nil {
			err = fmt.Errorf("hcrobot: server responded with status code %d", resp.StatusCode)
			return response, err
		}
		//else if isRetryable(err) {
		//				c.backoff(retries)
		//			retries++
		//		continue
		//	}
		return response, err
		*/
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &v); err != nil { // Parse []byte to the go struct pointer
		fmt.Println(io.Copy(os.Stdout, bytes.NewReader(body)))
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, bytes.NewReader(body))
		} else {
			err = json.Unmarshal(body, v)
		}
	}

	return response, err
}

type Response struct {
	*http.Response
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

//func errorFromResponse(s schema.Error) Error {}

// https://github.com/google/go-github/commit/994f6f8405f052a117d2d0b500054341048fbb08
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return strings.ToLower(u.String()), nil
}
