package connpass

import (
	"net/url"
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

const (
	defaultBaseURL = "https://connpass.com/api/v1/event/"
)

type Connpass struct {
	ResultsReturned  *int `json:"results_returned"`
	ResultsAvailable *int `json:"results_available"`
	ResultsStart     *int `json:"results_start"`
	Events           *[]Event `json:"events"`
}

type Event struct {
	EventId          *int `json:"event_id"`
	Title            *string `json:"title"`
	Catch            *string `json:"catch"`
	Description      *string `json:"description"`
	EventUrl         *string `json:"event_url"`
	HashTag          *string `json:"hash_tag"`
	StartedAt        *string `json:"started_at"`
	EndedAt          *string `json:"ended_at"`
	Limit            *int `json:"limit"`
	EventType        *string `json:"event_type"`
	Series                        `json:"series"`
	Address          *string `json:"address"`
	Place            *string `json:"place"`
	//Lat              float32 `json:lat`
	//Lon              float32 `json:lon`
	OwnerId          *int `json:"owner_id"`
	OwnerNickname    *string `json:"owner_nickname"`
	OwnerDisplayName *string `json:"owner_display_name"`
	Accepted         *int `json:"accepted"`
	Waiting          *int `json:"waiting"`
	Updated_at       *string `json:"updated_at"`
}

type Series struct {
	Id    *int `json:"id"`
	Title *string `json:"title"`
	Url   *string `json:"url"`
}

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	common  service // Reuse a single struct instead of allocating one for each service on the heap.

			// service
	EventId *eventIdService
	Keyword *keywordService
}

type service struct {
	client *Client
}

func NewClient() *Client {

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:http.DefaultClient,
		BaseURL:baseURL,
	}

	c.common.client = c
	c.EventId = (*eventIdService)(&c.common)
	c.Keyword = (*keywordService)(&c.common)
	return c
}

func ( c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

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
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}
