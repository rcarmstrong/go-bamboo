// package bamboo provides a client to communicate with Atlassian Bamboo CI Server API
//
// Usage:
//  import
//
// A Bamboo client exposes various services that control access to different parts of
// the Bamboo API. For example:
//  client := bamboo.NewSimpleClient(nil, "myUsername", "myPassword")
//
//  // Optionally set a different connection URL for the bamboo client.
//  // Defaults to "http://localhost:8085/rest/api/latest/"
//  client.SetURL("https://my.bambooserver.com:8085/")
//
//  planNames, _, err := client.Plans.ListPlanNames()

package bamboo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "http://localhost:8085/rest/api/latest/"
)

// Client manages the communication with the Bamboo API
type Client struct {
	client      *http.Client // HTTP client used to communicate with the API
	BaseURL     *url.URL
	SimpleCreds *SimpleCredentials // User credentials

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Bamboo API
	Info        *InfoService
	Plans       *PlanService
	Deploys     *DeployService
	Branches    *PlanBranchService
	Projects    *ProjectService
	Results     *ResultService
	Comments    *CommentService
	Labels      *LabelService
	Clone       *CloneService
	Server      *ServerService
	Permissions *Permissions
	Encryption  *Encryption
}

type service struct {
	name   string
	client *Client
}

// SetURL for the client to use for the Bamboo API
func (c *Client) SetURL(desiredURL string) error {
	newURL, err := url.Parse(desiredURL)
	if err != nil {
		return err
	}

	if newURL.Scheme == "" {
		return newErrBadURL("URL scheme was blank")
	}

	if !strings.HasSuffix(newURL.Path, "/rest/api/latest/") {
		newURL.Path += "/rest/api/latest/"
	}
	c.BaseURL = newURL
	return nil
}

// NewSimpleClient returns a new Bamboo API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an admin username/password
func NewSimpleClient(httpClient *http.Client, username, password string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, SimpleCreds: &SimpleCredentials{Username: username, Password: password}}
	c.common.client = c
	c.Plans = (*PlanService)(&c.common)
	c.Deploys = (*DeployService)(&c.common)
	c.Branches = (*PlanBranchService)(&c.common)
	c.Projects = (*ProjectService)(&c.common)
	c.Info = (*InfoService)(&c.common)
	c.Results = (*ResultService)(&c.common)
	c.Comments = (*CommentService)(&c.common)
	c.Labels = (*LabelService)(&c.common)
	c.Clone = (*CloneService)(&c.common)
	c.Server = (*ServerService)(&c.common)
	c.Permissions = (*Permissions)(&c.common)
	c.Encryption = (*Encryption)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	creds := c.SimpleCreds
	req.SetBasicAuth(creds.Username, creds.Password)
	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body) //nolint:errcheck
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}
